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

	fmt.Println("---")

	return strings.TrimSpace(output)
}

// ここできちんとツリーの状態が作られれば、あとはどうにかなりそうな気がするけど、、ここが
// TODO: 限界がきた。深さ優先探索の考え方で解く
func genTree(scanner *bufio.Scanner) *node {
	var rootNode *node
	var parentNode *node
	var prevNodeHierarchy int
	for {
		if !scanner.Scan() {
			break
		}

		row := scanner.Text()
		currentNode := newNode(row)

		// rootの場合
		if parentNode == nil {
			parentNode = currentNode
			prevNodeHierarchy = 1

			rootNode = currentNode
		} else {
			if currentNode.hierarchy == parentNode.hierarchy+1 {
				currentNode.parent = parentNode
				parentNode.children = append(parentNode.children, currentNode)
			}

			if !scanner.Scan() {
				break
			}
			row = scanner.Text()
			nextNode := newNode(row)

			// 前行のノードと親ノードが異なる
			if currentNode.hierarchy != prevNodeHierarchy {
				prevNodeHierarchy = currentNode.hierarchy

				// 現行のノード階層が次行のノード階層と異なる
				if currentNode.hierarchy != nextNode.hierarchy {
					parentNode = currentNode

					// 次行ノードの親が現行ノード
					if currentNode.hierarchy == (nextNode.hierarchy - 1) {
						nextNode.parent = currentNode
						currentNode.children = append(currentNode.children, nextNode)
					}

				} else {
					nextNode.parent = parentNode
					parentNode.children = append(parentNode.children, nextNode)
				}
			}
		}
	}

	return rootNode
}

// 描画するための枝を確定するロジック
// TODO:　あとは、ここで子のノードの個数とか同階層の上下にノードがあるか、とか見ていけば出来そうな気はする。
func computeTree(currentNode *node) {
	// rootでない
	if currentNode.hierarchy != 1 {
		fmt.Println("- debug -", len(currentNode.parent.children))
		for i := range currentNode.parent.children {
			fmt.Println(currentNode.parent.children[i].name)
		}

		// 親ノードの直接の子で最後の子
		isParentUnderEndRow := currentNode.name == currentNode.parent.children[len(currentNode.parent.children)-1].name

		if isParentUnderEndRow {
			currentNode.branch = convertEndTabTo(currentNode.hierarchy - 1)
		} else {
			currentNode.branch = convertIntermediateTabTo(currentNode.hierarchy - 1)
		}
	}

	for i := range currentNode.children {
		computeTree(currentNode.children[i])
	}
}

const convertedEndTab = "└" + "─" + "─"

func convertEndTabTo(tabCnt int) string {
	converted := ""
	if tabCnt == 0 {
		return converted
	}

	for i := 0; i < tabCnt-1; i++ {
		converted += "    "
	}
	converted += convertedEndTab
	return converted
}

const convertedIntermediateTab = "├" + "─" + "─"

func convertIntermediateTabTo(tabCnt int) string {
	converted := ""
	if tabCnt == 0 {
		return converted
	}

	for i := 0; i < tabCnt-1; i++ {
		converted += "    "
	}
	converted += convertedIntermediateTab
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
