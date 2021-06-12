package main

import (
	"bufio"
)

func generateTree(scanner *bufio.Scanner, isTwoSpaces, isFourSpaces bool) *node {
	var rootNode *node
	tmpStack := newStack()

	// rootを取得
	if scanner.Scan() {
		row := scanner.Text()
		rootNode = newNode(row, isTwoSpaces, isFourSpaces)
		tmpStack.push(rootNode)
	}

	for scanner.Scan() {
		row := scanner.Text()
		currentNode := newNode(row, isTwoSpaces, isFourSpaces)

		// 深さ優先探索的な考え方
		stackSize := tmpStack.size()
		for i := 0; i < stackSize; i++ {
			tmpNode := tmpStack.pop()

			// 現在のノードが親の直接の子
			if currentNode.hierarchy == tmpNode.hierarchy+1 {
				computeNode(tmpStack, tmpNode, currentNode)
				break
			}
		}
	}

	return rootNode
}

func computeNode(stack *stack, parentNode, childNode *node) {
	childNode.parent = parentNode
	parentNode.children = append(parentNode.children, childNode)
	stack.push(parentNode)
	stack.push(childNode)
}

// 描画するための枝を確定するロジック
func determineTreeBranch(currentNode *node) {
	// root
	if currentNode.hierarchy == 1 {
		for i := range currentNode.children {
			determineTreeBranch(currentNode.children[i])
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
		determineTreeBranch(currentNode.children[i])
	}
}

// 枝を展開する
func expandTree(currentNode *node, output string) string {
	output += currentNode.buildBranch()
	for i := range currentNode.children {
		output = expandTree(currentNode.children[i], output)
	}
	return output
}