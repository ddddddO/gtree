package main

import (
	"bufio"
)

type tree struct {
	root *node
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

			// 現在のノードが親の直接の子
			if currentNode.hierarchy == tmpNode.hierarchy+1 {
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
	// root
	if currentNode.hierarchy == 1 {
		for i := range currentNode.children {
			(*tree)(nil).determineBranches(currentNode.children[i])
		}
		return
	}

	parentNode := currentNode.parent
	lastChildIndex := len(parentNode.children) - 1
	// 階層の最後のノード
	if currentNode.index == parentNode.children[lastChildIndex].index {
		currentNode.branch += "└──"
	} else { // 階層の途中のノード
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
		if tmpNode.index == tmpParent.children[lastChildIndex].index {
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
	return (*tree)(nil).expandBranches(t.root, "")
}

// 枝を展開する
func (*tree) expandBranches(currentNode *node, output string) string {
	output += currentNode.buildBranch()
	for i := range currentNode.children {
		output = (*tree)(nil).expandBranches(currentNode.children[i], output)
	}
	return output
}
