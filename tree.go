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
		roots         []*Node
		stack         *stack
		nodeGenerator = newNodeGenerator(conf)
	)

	for scanner.Scan() {
		row := scanner.Text()
		currentNode := nodeGenerator.generate(row)

		if err := currentNode.validate(); err != nil {
			return nil, err
		}

		if currentNode.isRoot() {
			stack = newStack()
			roots = append(roots, currentNode)
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

	return &tree{
		roots:                 roots,
		formatLastNode:        conf.formatLastNode,
		formatIntermedialNode: conf.formatIntermedialNode,
	}, nil
}

func (t *tree) grow() *tree {
	for _, root := range t.roots {
		t.determineBranches(root)
	}
	return t
}

func (t *tree) determineBranches(currentNode *Node) {
	if currentNode.isRoot() {
		for _, child := range currentNode.children {
			t.determineBranches(child)
		}
		return
	}

	t.assembleBranchDirectly(currentNode)

	// rootまで親を遡って枝を構成する
	tmpParent := currentNode.parent
	for {
		// rootまで遡った
		if tmpParent.isRoot() {
			break
		}

		t.assembleBranchIndirectly(currentNode, tmpParent)

		tmpParent = tmpParent.parent
	}

	for _, child := range currentNode.children {
		t.determineBranches(child)
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
		branches += (*tree)(nil).expandBranches(root, "")
	}

	buf := bufio.NewWriter(w)
	if _, err := buf.WriteString(branches); err != nil {
		return err
	}
	return buf.Flush()
}

func (*tree) expandBranches(currentNode *Node, output string) string {
	output += currentNode.getBranch()
	for _, child := range currentNode.children {
		output = (*tree)(nil).expandBranches(child, output)
	}
	return output
}
