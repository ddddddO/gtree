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
	tree := sprout(bufio.NewScanner(input), conf.IsTwoSpaces, conf.IsFourSpaces) // 全入力をrootを頂点としたツリー上のデータに変換する。
	tree.grow()
	return tree.expand()
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
	determineBranches(t.root)
}

func determineBranches(currentNode *node) {
	isRoot := currentNode.hierarchy == rootHierarchyNum
	if isRoot {
		for i := range currentNode.children {
			determineBranches(currentNode.children[i])
		}
		return
	}

	parentNode := currentNode.parent
	if isLastNodeOfHierarchy(currentNode, parentNode) {
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
		if isLastNodeOfHierarchy(tmpNode, tmpParent) {
			currentNode.branch = "    " + currentNode.branch
		} else {
			currentNode.branch = "│   " + currentNode.branch
		}
		tmpNode = tmpParent
	}

	for i := range currentNode.children {
		determineBranches(currentNode.children[i])
	}
}

func isLastNodeOfHierarchy(tmpNode, parentNode *node) bool {
	lastChildIndex := len(parentNode.children) - 1
	return tmpNode.index == parentNode.children[lastChildIndex].index
}

func (t *tree) expand() string {
	branches := expandBranches(t.root, "")
	return strings.TrimSpace(branches)
}

func expandBranches(currentNode *node, output string) string {
	output += currentNode.buildBranch()
	for i := range currentNode.children {
		output = expandBranches(currentNode.children[i], output)
	}
	return output
}
