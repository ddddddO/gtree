//go:build !tinywasm

package gtree

import "iter"

// WalkerNode is used in user-defined function that can be executed with Walk/WalkProgrammably function.
type WalkerNode struct {
	origin *Node
}

// Name returns name of node in completed tree structure.
func (wn *WalkerNode) Name() string {
	return wn.origin.name
}

// Branch returns branch of node in completed tree structure.
func (wn *WalkerNode) Branch() string {
	return wn.origin.branch()
}

// Row returns row of node in completed tree structure.
func (wn *WalkerNode) Row() string {
	if !wn.origin.isRoot() {
		return wn.origin.branch() + " " + wn.origin.name
	}
	return wn.origin.name
}

// Level returns level of node in completed tree structure.
func (wn *WalkerNode) Level() uint {
	return wn.origin.hierarchy
}

// Path returns path of node in completed tree structure.
// Path is the path from the root node to this node.
// The separator is / in any OS execution environment.
func (wn *WalkerNode) Path() string {
	return wn.origin.path()
}

// HasChild returns whether the node in completed tree structure has child nodes.
func (wn *WalkerNode) HasChild() bool {
	return wn.origin.hasChild()
}

func newWalkerSimple() walkerSimple {
	return &defaultWalkerSimple{}
}

type defaultWalkerSimple struct{}

func (dw *defaultWalkerSimple) walk(roots []*Node, callback func(*WalkerNode) error) error {
	for _, root := range roots {
		if err := dw.walkNode(root, callback); err != nil {
			return err
		}
	}
	return nil
}

func (dw *defaultWalkerSimple) walkNode(current *Node, callback func(*WalkerNode) error) error {
	if err := callback(&WalkerNode{origin: current}); err != nil {
		return err
	}

	for _, child := range current.children {
		if err := dw.walkNode(child, callback); err != nil {
			return err
		}
	}
	return nil
}

// TODO: refactor. これは、複数root用ではなく、命名と他の実装と統一感がない。ただ、他もiter実装したいかというとそこまでのモチベーションはないので、このままでもいいかも
func (dw *defaultWalkerSimple) walkIter(root *Node) iter.Seq2[*WalkerNode, error] {
	return func(yiled func(*WalkerNode, error) bool) {
		dw.walkNodeForIter(root, yiled)
	}
}

func (dw *defaultWalkerSimple) walkNodeForIter(current *Node, yiled func(*WalkerNode, error) bool) {
	wn := &WalkerNode{origin: current}
	yiled(wn, nil)

	for _, child := range current.children {
		dw.walkNodeForIter(child, yiled)
	}
}
