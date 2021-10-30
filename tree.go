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
		tmpStack      *stack
		nodeGenerator = newNodeGenerator(conf)
	)

	for scanner.Scan() {
		row := scanner.Text()
		currentNode := nodeGenerator.generate(row)

		if err := currentNode.validate(); err != nil {
			return nil, err
		}

		if currentNode.isRoot() {
			tmpStack = newStack()
			roots = append(roots, currentNode)
			tmpStack.push(currentNode)
			continue
		}

		if tmpStack == nil {
			return nil, errNilStack
		}

		// 深さ優先探索的な？考え方
		stackSize := tmpStack.size()
		for i := 0; i < stackSize; i++ {
			tmpNode := tmpStack.pop()

			if currentNode.isDirectlyUnderParent(tmpNode) {
				tmpNode.addChild(currentNode)
				currentNode.setParent(tmpNode)
				tmpStack.push(tmpNode).push(currentNode)
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

	if currentNode.isLastOfHierarchy() {
		currentNode.branch += t.formatLastNode.directly
	} else {
		currentNode.branch += t.formatIntermedialNode.directly
	}

	// rootまで親を遡って枝を構成する
	tmpParent := currentNode.parent
	for {
		// rootまで遡った
		if tmpParent.isRoot() {
			break
		}

		if tmpParent.isLastOfHierarchy() {
			currentNode.branch = t.formatLastNode.indirectly + currentNode.branch
		} else {
			currentNode.branch = t.formatIntermedialNode.indirectly + currentNode.branch
		}
		tmpParent = tmpParent.parent
	}

	for _, child := range currentNode.children {
		t.determineBranches(child)
	}
}

func (t *tree) expand(w io.Writer) error {
	branches := ""
	for _, root := range t.roots {
		branches += expandBranches(root, "")
	}

	buf := bufio.NewWriter(w)
	if _, err := buf.WriteString(branches); err != nil {
		return err
	}
	return buf.Flush()
}

func expandBranches(currentNode *Node, output string) string {
	output += currentNode.getBranch()
	for _, child := range currentNode.children {
		output = expandBranches(child, output)
	}
	return output
}
