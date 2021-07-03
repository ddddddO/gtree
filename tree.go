package gentree

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
	root *node
}

func Execute(input io.Reader, conf Config) string {
	scanner := bufio.NewScanner(input)
	// ここで、全入力をrootを頂点としたツリー上のデータに変換する。
	tree := sprout(scanner, conf.IsTwoSpaces, conf.IsFourSpaces)
	tree.grow()
	output := tree.expand()

	return output
}

// Sprout：芽が出る
func sprout(scanner *bufio.Scanner, isTwoSpaces, isFourSpaces bool) *tree {
	var rootNode *node
	tmpStack := newStack()

	// rootを取得
	if scanner.Scan() {
		row := scanner.Text()
		rootNode = newNode(row, isTwoSpaces, isFourSpaces)
		tmpStack.push(rootNode)
	}

	// rootの子たちを取得
	for scanner.Scan() {
		row := scanner.Text()
		currentNode := newNode(row, isTwoSpaces, isFourSpaces)

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

	return &tree{
		root: rootNode,
	}
}

func (t *tree) grow() {
	(*tree)(nil).determineBranches(t.root)
}

// 描画するための枝を確定するロジック
func (*tree) determineBranches(currentNode *node) {
	isRoot := currentNode.hierarchy == rootHierarchyNum
	if isRoot {
		for i := range currentNode.children {
			(*tree)(nil).determineBranches(currentNode.children[i])
		}
		return
	}

	parentNode := currentNode.parent
	lastChildIndex := len(parentNode.children) - 1
	isLastNodeOfHierarchy := currentNode.index == parentNode.children[lastChildIndex].index
	if isLastNodeOfHierarchy {
		currentNode.branch += "└──"
	} else {
		currentNode.branch += "├──"
	}

	// rootまで親を遡って枝を構成する
	tmpNode := parentNode
	for {
		// rootまで遡った
		if tmpNode.parent == nil {
			break
		}

		tmpParent := tmpNode.parent
		lastChildIndex := len(tmpParent.children) - 1
		isLastNodeOfHierarchy := tmpNode.index == tmpParent.children[lastChildIndex].index
		if isLastNodeOfHierarchy {
			currentNode.branch = "    " + currentNode.branch
		} else {
			currentNode.branch = "│   " + currentNode.branch
		}
		tmpNode = tmpParent
	}

	for i := range currentNode.children {
		(*tree)(nil).determineBranches(currentNode.children[i])
	}
}

func (t *tree) expand() string {
	branches := (*tree)(nil).expandBranches(t.root, "")
	return strings.TrimSpace(branches)
}

// 枝を展開する
func (*tree) expandBranches(currentNode *node, output string) string {
	output += currentNode.buildBranch()
	for i := range currentNode.children {
		output = (*tree)(nil).expandBranches(currentNode.children[i], output)
	}
	return output
}
