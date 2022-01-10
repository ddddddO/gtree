package gtree

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
)

type treeer interface {
	addRoot(root *Node)
	grow() error
	expand(w io.Writer) error
	generate() error
}

type tree struct {
	roots                 []*Node
	formatLastNode        branchFormat
	formatIntermedialNode branchFormat
	dryrunMode            bool
}

type branchFormat struct {
	directly, indirectly string
}

func newTree(encode encode, formatLastNode, formatIntermedialNode branchFormat, dryrun bool) treeer {
	switch encode {
	case encodeJSON:
		return &jsonTree{
			&tree{},
		}
	case encodeYAML:
		return &yamlTree{
			&tree{},
		}
	case encodeTOML:
		return &tomlTree{
			&tree{},
		}
	default:
		return &tree{
			formatLastNode:        formatLastNode,
			formatIntermedialNode: formatIntermedialNode,
			dryrunMode:            dryrun,
		}
	}
}

// Execute outputs a tree to w with r as Markdown format input.
func Execute(w io.Writer, r io.Reader, optFns ...OptFn) error {
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
	return tree.expand(w)
}

// Generate generates directories.
func Generate(r io.Reader, optFns ...OptFn) error {
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
	return tree.generate()
}

// Sprout：芽が出る
// 全入力をrootを頂点としたツリー上のデータに変換する。
func sprout(scanner *bufio.Scanner, conf *config) (treeer, error) {
	var (
		stack            *stack
		counter          = newCounter()
		generateNodeFunc = decideGenerateFunc(conf.space)
		tree             = newTree(conf.encode, conf.formatLastNode, conf.formatIntermedialNode, conf.dryrun)
	)

	for scanner.Scan() {
		row := scanner.Text()
		currentNode := generateNodeFunc(row, counter.next())

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

		// 深さ優先探索的な？考え方
		stackSize := stack.size()
		for i := 0; i < stackSize; i++ {
			tmpNode := stack.pop()

			if currentNode.isDirectlyUnderParent(tmpNode) {
				tmpNode.addChild(currentNode)
				currentNode.setParent(tmpNode)
				stack.push(tmpNode).push(currentNode)
				break
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return tree, nil
}

func (t *tree) addRoot(root *Node) {
	t.roots = append(t.roots, root)
}

func (t *tree) grow() error {
	for _, root := range t.roots {
		if err := t.determineBranch(root); err != nil {
			return err
		}
	}
	return nil
}

func (t *tree) determineBranch(current *Node) error {
	if err := t.assembleBranch(current); err != nil {
		return err
	}

	for _, child := range current.Children {
		if err := t.determineBranch(child); err != nil {
			return err
		}
	}
	return nil
}

func (t *tree) assembleBranch(current *Node) error {
	if current.isRoot() {
		return nil
	}

	t.assembleBranchDirectly(current)

	// rootまで親を遡って枝を構成する
	tmpParent := current.parent
	for {
		// rootまで遡った
		if tmpParent.isRoot() {
			t.assembleBranchFinally(current, tmpParent)

			if t.dryrunMode {
				if err := current.validatePath(); err != nil {
					return err
				}
			}

			break
		}

		t.assembleBranchIndirectly(current, tmpParent)

		tmpParent = tmpParent.parent
	}
	return nil
}

func (t *tree) assembleBranchDirectly(current *Node) {
	current.branch.path = current.Name

	if current.isLastOfHierarchy() {
		current.branch.value += t.formatLastNode.directly
	} else {
		current.branch.value += t.formatIntermedialNode.directly
	}
}

func (t *tree) assembleBranchIndirectly(current, parent *Node) {
	current.branch.path = filepath.Join(parent.Name, current.branch.path)

	if parent.isLastOfHierarchy() {
		current.branch.value = t.formatLastNode.indirectly + current.branch.value
	} else {
		current.branch.value = t.formatIntermedialNode.indirectly + current.branch.value
	}
}

func (t *tree) assembleBranchFinally(current, root *Node) {
	current.branch.path = filepath.Join(root.Name, current.branch.path)
}

func (t *tree) expand(w io.Writer) error {
	branches := ""
	for _, root := range t.roots {
		branches += (*tree)(nil).expandBranch(root, "")
	}

	return (*tree)(nil).write(w, branches)
}

func (*tree) expandBranch(current *Node, out string) string {
	out += current.getBranch()
	for _, child := range current.Children {
		out = (*tree)(nil).expandBranch(child, out)
	}
	return out
}

func (*tree) write(w io.Writer, in string) error {
	buf := bufio.NewWriter(w)
	if _, err := buf.WriteString(in); err != nil {
		return err
	}
	return buf.Flush()
}

func (t *tree) generate() error {
	for _, root := range t.roots {
		if err := (*tree)(nil).generateDirectory(root); err != nil {
			return err
		}
	}
	return nil
}

const permission = 0777

func (*tree) generateDirectory(current *Node) error {
	// only root node exists
	if current.isRoot() && !current.hasChild() {
		err := os.MkdirAll(current.branch.path, permission)
		if err != nil {
			return err
		}
		return nil
	}

	for _, child := range current.Children {
		if !child.hasChild() {
			err := os.MkdirAll(child.branch.path, permission)
			if err != nil {
				return err
			}
			continue
		}

		err := (*tree)(nil).generateDirectory(child)
		if err != nil {
			return err
		}
	}
	return nil
}
