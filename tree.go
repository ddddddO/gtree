package gtree

import (
	"bufio"
	"io"
	"strings"
)

type Config struct {
	IsTwoSpaces  bool
	IsFourSpaces bool
}

type tree struct {
	roots []*node
}

func Execute(input io.Reader, conf Config) (string, error) {
	seed := bufio.NewScanner(input)
	nodeGenerator := newNodeGenerator(conf)

	tree, err := sprout(seed, nodeGenerator)
	if err != nil {
		return "", err
	}

	tree.grow()
	return tree.expand(), nil
}

// Sprout：芽が出る
// 全入力をrootを頂点としたツリー上のデータに変換する。
func sprout(scanner *bufio.Scanner, nodeGenerator nodeGenerator) (*tree, error) {
	var (
		roots    []*node
		tmpStack *stack
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
			return nil, ErrNilStack
		}

		// 深さ優先探索的な？考え方
		stackSize := tmpStack.size()
		for i := 0; i < stackSize; i++ {
			tmpNode := tmpStack.pop()
			isCurrentNodeDirectlyUnderParent := currentNode.hierarchy == tmpNode.hierarchy+1

			if isCurrentNodeDirectlyUnderParent {
				parent := tmpNode
				child := currentNode

				parent.children = append(parent.children, child)
				child.parent = parent
				tmpStack.push(parent)
				tmpStack.push(child)
				break
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return &tree{
		roots: roots,
	}, nil
}

func (t *tree) grow() {
	for _, root := range t.roots {
		determineBranches(root)
	}
}

func determineBranches(currentNode *node) {
	if currentNode.isRoot() {
		for _, child := range currentNode.children {
			determineBranches(child)
		}
		return
	}

	if currentNode.isLastNodeOfHierarchy() {
		currentNode.branch += "└──"
	} else {
		currentNode.branch += "├──"
	}

	// rootまで親を遡って枝を構成する
	tmpParent := currentNode.parent
	for {
		// rootまで遡った
		if tmpParent.isRoot() {
			break
		}

		if tmpParent.isLastNodeOfHierarchy() {
			currentNode.branch = "    " + currentNode.branch
		} else {
			currentNode.branch = "│   " + currentNode.branch
		}
		tmpParent = tmpParent.parent
	}

	for _, child := range currentNode.children {
		determineBranches(child)
	}
}

func (t *tree) expand() string {
	branches := ""
	for _, root := range t.roots {
		branches += expandBranches(root, "")
	}
	return strings.TrimSpace(branches)
}

func expandBranches(currentNode *node, output string) string {
	output += currentNode.buildBranch()
	for _, child := range currentNode.children {
		output = expandBranches(child, output)
	}
	return output
}
