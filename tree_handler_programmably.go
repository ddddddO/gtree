// Package gtree provides tree-structured output.
package gtree

import (
	"context"
	"errors"
	"io"

	"github.com/fatih/color"
)

var (
	idxCounter = newCounter()
)

var (
	// ErrNilNode is returned if the argument *gtree.Node of OutputProgrammably / MkdirProgrammably function is nill.
	ErrNilNode = errors.New("nil node")
	// ErrNotRoot is returned if the argument *gtree.Node of OutputProgrammably / MkdirProgrammably function is not root of the tree.
	ErrNotRoot = errors.New("not root node")
)

// OutputProgrammably outputs tree to w.
// This function requires node generated by NewRoot function.
func OutputProgrammably(w io.Writer, root *Node, options ...Option) error {
	if root == nil {
		return ErrNilNode
	}
	if !root.isRoot() {
		return ErrNotRoot
	}

	conf, err := newConfig(options)
	if err != nil {
		return err
	}

	idxCounter.reset()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	rootStream := make(chan *Node)
	go func() {
		defer close(rootStream)
		rootStream <- root
	}()

	tree := newTree(conf)
	growingStream, errcg := tree.grow(ctx, rootStream)
	errcs := tree.spread(ctx, w, growingStream)

	return handlePipelineErr(errcg, errcs)
}

var (
	// ErrExistPath is returned if the argument *gtree.Node of MkdirProgrammably function is path already exists.
	ErrExistPath = errors.New("path already exists")
)

// MkdirProgrammably makes directories.
// This function requires node generated by NewRoot function.
func MkdirProgrammably(root *Node, options ...Option) error {
	if root == nil {
		return ErrNilNode
	}
	if !root.isRoot() {
		return ErrNotRoot
	}

	conf, err := newConfig(options)
	if err != nil {
		return err
	}

	idxCounter.reset()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	rootStream := make(chan *Node)
	go func() {
		defer close(rootStream)
		rootStream <- root
	}()

	tree := newTree(conf)
	tree.enableValidation()
	// when detect invalid node name, return error. process end.
	growingStream, errcg := tree.grow(ctx, rootStream)
	if conf.dryrun {
		// when detected no invalid node name, output tree.
		errcs := tree.spread(ctx, color.Output, growingStream)
		return handlePipelineErr(errcg, errcs)
	}
	// when detected no invalid node name, no output tree.
	errcm := tree.mkdir(ctx, growingStream)
	return handlePipelineErr(errcg, errcm)
}

func (t *tree) enableValidation() {
	t.grower.enableValidation()
}

// NewRoot creates a starting node for building tree.
func NewRoot(text string) *Node {
	return newNode(text, rootHierarchyNum, idxCounter.next())
}

// Add adds a node and returns an instance of it.
// If a node with the same text already exists in the same hierarchy of the tree, that node will be returned.
func (parent *Node) Add(text string) *Node {
	if child := parent.findChildByText(text); child != nil {
		return child
	}

	current := newNode(text, parent.hierarchy+1, idxCounter.next())
	current.setParent(parent)
	parent.addChild(current)
	return current
}

func (parent *Node) findChildByText(text string) *Node {
	for _, child := range parent.children {
		if text == child.name {
			return child
		}
	}
	return nil
}
