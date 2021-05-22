package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
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

type stack struct {
	nodes []*node
}

func newStack() *stack {
	return &stack{}
}

func (s *stack) push(n *node) {
	s.nodes = append(s.nodes, n)
}

func (s *stack) pop() *node {
	lastIndex := len(s.nodes) - 1
	tmp := s.nodes[lastIndex]
	s.nodes = s.nodes[:lastIndex]
	return tmp
}

func (s *stack) size() int {
	return len(s.nodes)
}

func (s *stack) lastStackedHierarchy() int {
	lastIndex := len(s.nodes) - 1
	return s.nodes[lastIndex].hierarchy
}

func main() {
	var f string
	flag.StringVar(&f, "f", "", "markdown file path")
	flag.Parse()

	var input io.Reader
	if f == "" || f == "-" {
		input = os.Stdin
	} else {
		filePath, err := filepath.Abs(f)
		if err != nil {
			fmt.Errorf("%+v", err)
			os.Exit(1)
		}
		input, err = os.Open(filePath)
		if err != nil {
			fmt.Errorf("%+v", err)
			os.Exit(1)
		}
		defer input.(*os.File).Close()
	}

	fmt.Println(gen(input))
}

func gen(input io.Reader) string {
	scanner := bufio.NewScanner(input)

	// ここで、全入力をrootを頂点としたツリー上のデータに変換する。
	tree := genTree(scanner)
	computeTree(tree)
	output := expandTree(tree, "")

	return strings.TrimSpace(output)
}

// 深さ優先探索的な考え方
func genTree(scanner *bufio.Scanner) *node {
	var rootNode *node
	tmpStack := newStack()

	// rootを取得
	if scanner.Scan() {
		row := scanner.Text()
		rootNode = newNode(row)
		tmpStack.push(rootNode)
	}

	for scanner.Scan() {
		row := scanner.Text()
		currentNode := newNode(row)

		// 親+1の階層
		if currentNode.hierarchy == tmpStack.lastStackedHierarchy()+1 {
			lastStackedNode := tmpStack.pop()
			currentNode.parent = lastStackedNode
			lastStackedNode.children = append(lastStackedNode.children, currentNode)
			tmpStack.push(lastStackedNode)
			tmpStack.push(currentNode)
			continue
		}

		// 最後のスタックと同階層
		if currentNode.hierarchy == tmpStack.lastStackedHierarchy() {
			_ = tmpStack.pop()

			lastStackedNode := tmpStack.pop()
			currentNode.parent = lastStackedNode
			lastStackedNode.children = append(lastStackedNode.children, currentNode)
			tmpStack.push(lastStackedNode)
			tmpStack.push(currentNode)
			continue
		}

		// 最後のスタックよりrootに近い
		stackSize := tmpStack.size()
		for i := 0; i < stackSize; i++ {
			_ = tmpStack.pop()

			if currentNode.hierarchy == tmpStack.lastStackedHierarchy()+1 {
				lastStackedNode := tmpStack.pop()
				currentNode.parent = lastStackedNode
				lastStackedNode.children = append(lastStackedNode.children, currentNode)
				tmpStack.push(lastStackedNode)
				tmpStack.push(currentNode)
				break
			}
		}
	}

	return rootNode
}

// 描画するための枝を確定するロジック
func computeTree(currentNode *node) {
	// rootでない
	if currentNode.hierarchy != 1 {
		// 親ノードの直接の子で最後の子
		isParentUnderEndRow := currentNode.name == currentNode.parent.children[len(currentNode.parent.children)-1].name

		if isParentUnderEndRow {
			currentNode.branch = convertEndTabTo(currentNode)
		} else {
			currentNode.branch = convertIntermediateTabTo(currentNode)
		}
	}

	for i := range currentNode.children {
		computeTree(currentNode.children[i])
	}
}

const convertedEndTab = "└" + "─" + "─"

const fourSpaces = "    "

func convertEndTabTo(currentNode *node) string {
	converted := ""
	tabCnt := currentNode.hierarchy - 1
	if tabCnt == 0 {
		return converted
	}

	converted = dp(currentNode, convertedEndTab, converted, 0)
	if converted != "" {
		return converted
	}

	for i := 0; i < tabCnt-1; i++ {
		converted += fourSpaces
	}
	converted += convertedEndTab
	return converted
}

const convertedIntermediateTab = "├" + "─" + "─"

func convertIntermediateTabTo(currentNode *node) string {
	converted := ""
	tabCnt := currentNode.hierarchy - 1
	if tabCnt == 0 {
		return converted
	}

	// memo: case 5のuとkを確認する
	// 親がいてかつ、子がいる場合(u)
	if currentNode.parent != nil && currentNode.children != nil {
		for j := 0; j < (currentNode.hierarchy - 2); j++ {
			converted += tmp
		}
		converted += convertedIntermediateTab
		return converted
	}

	converted = dp(currentNode, convertedIntermediateTab, converted, 0)
	if converted != "" {
		return converted
	}

	for i := 0; i < tabCnt-1; i++ {
		converted += fourSpaces
	}
	converted += convertedIntermediateTab
	return converted
}

const tmp = "│   "

// FIXME: rootまで遡らないと多分ダメ
//        多分、ノードからrootまで親を一つずつ遡って、一つずつノード側から枝を構成する記号を組み立てて行かないとダメっぽいし、そうした方が保守できそうな形になりそう。
func dp(currentNode *node, template string, converted string, circuitCnt int /*何回目のdpか。初回は0*/) string {
	if currentNode.parent == nil {
		return converted
	}

	// 親の親がいてかつ、親と同階層の下にノードがいる場合(k)
	if currentNode.parent.parent != nil {
		children := currentNode.parent.parent.children // 親の親の子たち
		for i := range children {
			if currentNode.parent.name == children[i].name && len(children) > i+1 {
				for j := 0; j < (currentNode.hierarchy - 2); j++ {
					converted += tmp
				}
				for k := 0; k < circuitCnt; k++ {
					converted += fourSpaces
				}
				converted += template
				return converted
			}
		}

		circuitCnt++
		return dp(currentNode.parent, template, converted, circuitCnt)
	}
	return ""
}

// 枝を展開する
func expandTree(node *node, output string) string {
	output += node.buildBranch()
	for i := range node.children {
		output = expandTree(node.children[i], output)
	}
	return output
}
