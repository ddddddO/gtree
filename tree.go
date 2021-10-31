package gtree

import (
	"bufio"
	"io"
)

type tree struct {
	roots                 []*Node
	formatLastNode        branchFormat
	formatIntermedialNode branchFormat
}

type branchFormat struct {
	directly, indirectly string
}

// Execute outputs a tree to w with r as Markdown format input.
func Execute(w io.Writer, r io.Reader, optFns ...optFn) error {
	conf, err := newConfig(optFns...)
	if err != nil {
		return err
	}
	seed := bufio.NewScanner(r)

	tree, err := sprout(seed, conf)
	if err != nil {
		return err
	}
	return tree.grow().expand(w)
}

// Sprout：芽が出る
// 全入力をrootを頂点としたツリー上のデータに変換する。
func sprout(scanner *bufio.Scanner, conf *config) (*tree, error) {
	var (
		stack        *stack
		generateNode = decideGenerateFunc(conf)
		tree         = &tree{
			formatLastNode:        conf.formatLastNode,
			formatIntermedialNode: conf.formatIntermedialNode,
		}
	)

	for scanner.Scan() {
		row := scanner.Text()
		currentNode := generateNode(row)

		if err := currentNode.validate(); err != nil {
			return nil, err
		}

		if currentNode.isRoot() {
			tree.addRoot(currentNode)
			stack = newStack()
			stack.push(currentNode)
			continue
		}

		if stack == nil {
			return nil, errNilStack
		}

		// 深さ優先探索的な？考え方
		stackSize := stack.size()
		for i := 0; i < stackSize; i++ {
			tmpNode := stack.pop()

			if currentNode.isDirectlyUnderParent(tmpNode) {
				tmpNode.addChild(currentNode)
				currentNode.setParent(tmpNode)
				stack.push(tmpNode).push(currentNode)
				break
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return tree, nil
}

func (t *tree) addRoot(root *Node) {
	t.roots = append(t.roots, root)
}

func (t *tree) grow() *tree {
	for _, root := range t.roots {
		t.determineBranch(root)
	}
	return t
}

func (t *tree) determineBranch(current *Node) {
	t.assembleBranch(current)

	for _, child := range current.children {
		t.determineBranch(child)
	}
}

func (t *tree) assembleBranch(current *Node) {
	if current.isRoot() {
		return
	}

	t.assembleBranchDirectly(current)

	// rootまで親を遡って枝を構成する
	tmpParent := current.parent
	for {
		// rootまで遡った
		if tmpParent.isRoot() {
			break
		}

		t.assembleBranchIndirectly(current, tmpParent)

		tmpParent = tmpParent.parent
	}
}

func (t *tree) assembleBranchDirectly(current *Node) {
	if current.isLastOfHierarchy() {
		current.branch += t.formatLastNode.directly
	} else {
		current.branch += t.formatIntermedialNode.directly
	}
}

func (t *tree) assembleBranchIndirectly(current, parent *Node) {
	if parent.isLastOfHierarchy() {
		current.branch = t.formatLastNode.indirectly + current.branch
	} else {
		current.branch = t.formatIntermedialNode.indirectly + current.branch
	}
}

func (t *tree) expand(w io.Writer) error {
	branches := ""
	for _, root := range t.roots {
		branches += (*tree)(nil).expandBranch(root, "")
	}

	return (*tree)(nil).write(w, branches)
}

func (*tree) expandBranch(current *Node, output string) string {
	output += current.getBranch()
	for _, child := range current.children {
		output = (*tree)(nil).expandBranch(child, output)
	}
	return output
}

func (*tree) write(w io.Writer, branches string) error {
	buf := bufio.NewWriter(w)
	if _, err := buf.WriteString(branches); err != nil {
		return err
	}
	return buf.Flush()
}