package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type node struct {
	name      string
	hierarchy int
	branch    string
	parent    *node
	children  []*node
}

func newNode(row string) *node {
	myselfNode := &node{}
	name := ""
	hierarchy := 1

	for _, r := range row {
		// https://ja.wikipedia.org/wiki/ASCII
		switch r {
		case 45: // -
			continue
		case 32: // space
			continue
		case 9: // tab
			hierarchy++
			continue
		default: // directry or file name char
			name += string(r)
		}
	}
	myselfNode.name = name
	myselfNode.hierarchy = hierarchy
	return myselfNode
}

func (n *node) buildBranch() string {
	if n.hierarchy == 1 {
		return n.name + fmt.Sprintln()
	}
	return n.branch + " " + n.name + fmt.Sprintln()
}

func main() {
	// s := `aa`
	// r := strings.NewReader(s)
	r := os.Stdin
	fmt.Println(gen(r))
}

func gen(input io.Reader) string {
	var scanner *bufio.Scanner

	switch input.(type) {
	case *strings.Reader:
		scanner = bufio.NewScanner(input)
	case *os.File:
		scanner = bufio.NewScanner(input)
	default:
		panic("unsupported type")
	}

	// ここで、全入力をrootを頂点としたツリー上のデータに変換する。
	tree := genTree(scanner)
	computeTree(tree)
	output := expandTree(tree, "")

	return strings.TrimSpace(output)
}

func genTree(scanner *bufio.Scanner) *node {
	var rootNode *node
	var parentNode *node
	var prevNodeHierarchy int
	for scanner.Scan() {
		row := scanner.Text()
		currentNode := newNode(row)

		// rootの場合
		if parentNode == nil {
			parentNode = currentNode
			prevNodeHierarchy = 1

			rootNode = currentNode
		} else {
			currentNode.parent = parentNode
			parentNode.children = append(parentNode.children, currentNode)

			// 前行のノードと親ノードが異なる場合
			if currentNode.hierarchy != prevNodeHierarchy {
				parentNode = currentNode
				prevNodeHierarchy = currentNode.hierarchy
			}
		}
	}

	return rootNode
}

// 描画するための枝を確定するロジック
// TODO:　あとは、ここで子のノードの個数とか同階層の上下にノードがあるか、とか見ていけば出来そうな気はする。
func computeTree(node *node) {
	// root
	if node.hierarchy == 1 {
		node.branch = convertTab(node.hierarchy - 1)
	} else {
		node.branch = convertTab(node.hierarchy - 1)
	}

	for i := range node.children {
		computeTree(node.children[i])
	}
}

const convertedTab = "└" + "─" + "─"

func convertTab(tabCnt int) string {
	converted := ""
	if tabCnt == 0 {
		return converted
	}

	for i := 0; i < tabCnt-1; i++ {
		converted += "    "
	}
	converted += convertedTab
	return converted
}

var output = ""

// 枝を展開する
func expandTree(node *node, output string) string {
	output += node.buildBranch()

	for i := range node.children {
		output = expandTree(node.children[i], output)
	}
	return output
}
