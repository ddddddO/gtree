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

var (
	isTwoSpaces  bool
	isFourSpaces bool
)

func newNode(row string) *node {
	myselfNode := &node{}
	name := ""
	hierarchy := 1

	spaceCnt := 0
	for _, r := range row {
		// https://ja.wikipedia.org/wiki/ASCII
		switch r {
		case 45: // -
			if isTwoSpaces && spaceCnt%2 == 0 {
				tmp := spaceCnt / 2
				hierarchy += tmp
			}
			if isFourSpaces && spaceCnt%4 == 0 {
				tmp := spaceCnt / 4
				hierarchy += tmp
			}
		case 32: // space
			spaceCnt++
		case 9: // tab
			hierarchy++
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

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("failed to gentree...\nplease review the file format.\nhint: %s\n", err)
			os.Exit(1)
		}
	}()

	var f string
	flag.StringVar(&f, "f", "", "markdown file path")
	flag.BoolVar(&isTwoSpaces, "ts", false, "for indent two spaces")
	flag.BoolVar(&isFourSpaces, "fs", false, "for indent four spaces")
	flag.Parse()

	if isTwoSpaces && isFourSpaces {
		fmt.Errorf("%s", `choose either "ts" or "fs".`)
		os.Exit(1)
	}

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
	tree := generateTree(scanner)
	determineTreeBranch(tree)
	output := expandTree(tree, "")

	return strings.TrimSpace(output)
}

func generateTree(scanner *bufio.Scanner) *node {
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

		// 深さ優先探索的な考え方
		stackSize := tmpStack.size()
		for i := 0; i < stackSize; i++ {
			tmpNode := tmpStack.pop()

			if currentNode.hierarchy == tmpNode.hierarchy+1 {
				parentNode := tmpNode
				computeNode(tmpStack, currentNode, parentNode)
				break
			}
		}
	}

	return rootNode
}

func computeNode(stack *stack, currentNode, parentNode *node) {
	currentNode.parent = parentNode
	parentNode.children = append(parentNode.children, currentNode)
	stack.push(parentNode)
	stack.push(currentNode)
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
	if currentNode.name == parentNode.children[lastChildIndex].name {
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
		if tmpNode.name == tmpParent.children[lastChildIndex].name {
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
func expandTree(node *node, output string) string {
	output += node.buildBranch()
	for i := range node.children {
		output = expandTree(node.children[i], output)
	}
	return output
}
