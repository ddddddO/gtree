//go:build !tinywasm

package gtree

import (
	"errors"
	"io"
)

var (
	idxCounter = newCounter()
)

// OutputProgrammably outputs tree to w.
// This function requires node generated by NewRoot function.
func OutputProgrammably(w io.Writer, root *Node, options ...Option) error {
	if err := validateTreeRoot(root); err != nil {
		return err
	}

	idxCounter.reset()

	cfg := newConfig(options)
	return initializeTree(cfg).outputProgrammably(w, root, cfg)
}

// MkdirProgrammably makes directories.
// This function requires node generated by NewRoot function.
func MkdirProgrammably(root *Node, options ...Option) error {
	if err := validateTreeRoot(root); err != nil {
		return err
	}

	idxCounter.reset()

	cfg := newConfig(options)
	return initializeTree(cfg).mkdirProgrammably(root, cfg)
}

// VerifyProgrammably verifies directory.
// This function requires node generated by NewRoot function.
func VerifyProgrammably(root *Node, options ...Option) error {
	if err := validateTreeRoot(root); err != nil {
		return err
	}

	idxCounter.reset()

	cfg := newConfig(options)
	return initializeTree(cfg).verifyProgrammably(root, cfg)
}

// TODO: add doc
func WalkProgrammably(root *Node, cb func(*WalkerNode) error, options ...Option) error {
	if err := validateTreeRoot(root); err != nil {
		return err
	}

	idxCounter.reset()

	cfg := newConfig(options)
	return initializeTree(cfg).walkProgrammably(root, cb, cfg)
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

var (
	// ErrNilNode is returned if the argument *gtree.Node of OutputProgrammably / MkdirProgrammably / VerifyProgrammably function is nill.
	ErrNilNode = errors.New("nil node")
	// ErrNotRoot is returned if the argument *gtree.Node of OutputProgrammably / MkdirProgrammably / VerifyProgrammably function is not root of the tree.
	ErrNotRoot = errors.New("not root node")
)

func validateTreeRoot(root *Node) error {
	if root == nil {
		return ErrNilNode
	}
	if !root.isRoot() {
		return ErrNotRoot
	}
	return nil
}
