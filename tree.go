package gtree

import (
	"bufio"
	"io"
)

// Output outputs a tree to w with r as Markdown format input.
func Output(w io.Writer, r io.Reader, optFns ...OptFn) error {
	conf, err := newConfig(optFns...)
	if err != nil {
		return err
	}
	seed := bufio.NewScanner(r)

	tree, err := sprout(seed, conf)
	if err != nil {
		return err
	}
	if err := tree.grow(); err != nil {
		return err
	}
	return tree.spread(w)
}

// Mkdir makes directories.
func Mkdir(r io.Reader, optFns ...OptFn) error {
	conf, err := newConfig(optFns...)
	if err != nil {
		return err
	}
	seed := bufio.NewScanner(r)

	tree, err := sprout(seed, conf)
	if err != nil {
		return err
	}
	if err := tree.grow(); err != nil {
		return err
	}
	return tree.mkdir()
}

func sprout(scanner *bufio.Scanner, conf *config) (*tree, error) {
	var (
		stack            *stack
		counter          = newCounter()
		generateNodeFunc = decideGenerateFunc(conf.space)
		tree             = newTree(conf.encode, conf.lastNodeFormat, conf.intermedialNodeFormat, conf.dryrun, conf.fileExtensions)
	)

	for scanner.Scan() {
		currentNode := generateNodeFunc(scanner.Text(), counter.next())
		if err := currentNode.validate(); err != nil {
			return nil, err
		}

		if currentNode.isRoot() {
			counter.reset()
			tree.addRoot(currentNode)
			stack = newStack()
			stack.push(currentNode)
			continue
		}

		if stack == nil {
			return nil, errNilStack
		}

		stack.dfs(currentNode)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return tree, nil
}

type tree struct {
	roots    []*Node
	grower   grower
	spreader spreader
	mkdirer  mkdirer
}

func newTree(
	encode encode,
	lastNodeFormat, intermedialNodeFormat branchFormat,
	dryrun bool,
	fileExtensions []string,
) *tree {
	return &tree{
		grower:   newGrower(encode, lastNodeFormat, intermedialNodeFormat, dryrun),
		spreader: newSpreader(encode),
		mkdirer:  newMkdirer(fileExtensions),
	}
}

type branchFormat struct {
	directly, indirectly string
}

func (t *tree) addRoot(root *Node) {
	t.roots = append(t.roots, root)
}

func (t *tree) grow() error {
	return t.grower.grow(t.roots)
}

// TODO: メソッド名見直す
func (t *tree) enableValidation() {
	t.grower.setDryRun(true)
}

func (t *tree) spread(w io.Writer) error {
	return t.spreader.spread(w, t.roots)
}

func (t *tree) mkdir() error {
	return t.mkdirer.mkdir(t.roots)
}
