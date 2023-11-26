//go:build !tinywasm

package gtree

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
