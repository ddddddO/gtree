package gtree

import (
	"bufio"
	"io"

	"github.com/pkg/errors"
)

// optFunc is functional options pattern
type optFn func(*config) error

// IndentTwoSpaces returns function for two spaces indent input.
func IndentTwoSpaces() optFn {
	return func(c *config) error {
		c.IsTwoSpaces = true
		return nil
	}
}

// IndentFourSpaces returns function for four spaces indent input.
func IndentFourSpaces() optFn {
	return func(c *config) error {
		c.IsFourSpaces = true
		return nil
	}
}

var errInvalidOption = errors.New("invalid option")

type config struct {
	IsTwoSpaces  bool
	IsFourSpaces bool
}

func newConfig(optFns ...optFn) (*config, error) {
	c := &config{}
	for _, opt := range optFns {
		if err := opt(c); err != nil {
			return nil, err
		}
	}

	if c.IsTwoSpaces && c.IsFourSpaces {
		return nil, errInvalidOption
	}
	return c, nil
}

type lastNodeFormat struct {
	directly, indirectly string
}

type intermedialNodeFormat struct {
	directly, indirectly string
}

type tree struct {
	roots                 []*Node
	lastNodeFormat        lastNodeFormat
	intermedialNodeFormat intermedialNodeFormat
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
	tree.grow()
	return tree.expand(w)
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

	// TODO: ユーザーが枝のフォーマットを決められるようにする
	return &tree{
		roots: roots,
		lastNodeFormat: lastNodeFormat{
			directly:   "└──",
			indirectly: "    ",
		},
		intermedialNodeFormat: intermedialNodeFormat{
			directly:   "├──",
			indirectly: "│   ",
		},
	}, nil
}

func (t *tree) grow() {
	for _, root := range t.roots {
		t.determineBranches(root)
	}
}

func (t *tree) determineBranches(currentNode *Node) {
	if currentNode.isRoot() {
		for _, child := range currentNode.children {
			t.determineBranches(child)
		}
		return
	}

	if currentNode.isLastNodeOfHierarchy() {
		currentNode.branch += t.lastNodeFormat.directly
	} else {
		currentNode.branch += t.intermedialNodeFormat.directly
	}

	// rootまで親を遡って枝を構成する
	tmpParent := currentNode.parent
	for {
		// rootまで遡った
		if tmpParent.isRoot() {
			break
		}

		if tmpParent.isLastNodeOfHierarchy() {
			currentNode.branch = t.lastNodeFormat.indirectly + currentNode.branch
		} else {
			currentNode.branch = t.intermedialNodeFormat.indirectly + currentNode.branch
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
	output += currentNode.buildBranch()
	for _, child := range currentNode.children {
		output = expandBranches(child, output)
	}
	return output
}
