//go:build !wasm

package gtree

import (
	"io"

	"github.com/fatih/color"
)

type treeSimple struct {
	grower   growerSimple
	spreader spreaderSimple
	mkdirer  mkdirerSimple
}

func newTreeSimple(conf *config) iTree {
	growerFactory := func(lastNodeFormat, intermedialNodeFormat branchFormat, dryrun bool, encode encode) growerSimple {
		if encode != encodeDefault {
			return newNopGrowerSimple()
		}
		return newGrowerSimple(lastNodeFormat, intermedialNodeFormat, dryrun)
	}

	spreaderFactory := func(encode encode, dryrun bool, fileExtensions []string) spreaderSimple {
		if dryrun {
			return newColorizeSpreaderSimple(fileExtensions)
		}
		return newSpreaderSimple(encode)
	}

	mkdirerFactory := func(fileExtensions []string) mkdirerSimple {
		return newMkdirerSimple(fileExtensions)
	}

	return &treeSimple{
		grower: growerFactory(
			conf.lastNodeFormat,
			conf.intermedialNodeFormat,
			conf.dryrun,
			conf.encode,
		),
		spreader: spreaderFactory(
			conf.encode,
			conf.dryrun,
			conf.fileExtensions,
		),
		mkdirer: mkdirerFactory(
			conf.fileExtensions,
		),
	}
}

func (t *treeSimple) output(w io.Writer, r io.Reader, conf *config) error {
	rg := newRootGeneratorSimple(r, conf.space)
	roots, err := rg.generate()
	if err != nil {
		return err
	}

	if err := t.grow(roots); err != nil {
		return err
	}
	return t.spread(w, roots)
}

func (t *treeSimple) outputProgrammably(w io.Writer, root *Node, conf *config) error {
	if err := t.grow([]*Node{root}); err != nil {
		return err
	}
	return t.spread(w, []*Node{root})
}

func (t *treeSimple) makedir(r io.Reader, conf *config) error {
	rg := newRootGeneratorSimple(r, conf.space)
	roots, err := rg.generate()
	if err != nil {
		return err
	}

	if err := t.grow(roots); err != nil {
		return err
	}
	return t.mkdir(roots)
}

func (t *treeSimple) makedirProgrammably(root *Node, conf *config) error {
	t.enableValidation()
	// when detect invalid node name, return error. process end.
	if err := t.grow([]*Node{root}); err != nil {
		return err
	}
	if conf.dryrun {
		// when detected no invalid node name, output tree.
		return t.spread(color.Output, []*Node{root})
	}
	// when detected no invalid node name, no output tree.
	return t.mkdir([]*Node{root})
}

// 関心事は各ノードの枝の形成
type growerSimple interface {
	grow([]*Node) error
	enableValidation()
}

// 関心事はtreeの出力
type spreaderSimple interface {
	spread(io.Writer, []*Node) error
}

// 関心事はファイルの生成
// interfaceを使う必要はないが、growerSimple/spreaderSimpleと合わせたいため
type mkdirerSimple interface {
	mkdir([]*Node) error
}

func (t *treeSimple) grow(roots []*Node) error {
	return t.grower.grow(roots)
}

func (t *treeSimple) spread(w io.Writer, roots []*Node) error {
	return t.spreader.spread(w, roots)
}

func (t *treeSimple) mkdir(roots []*Node) error {
	return t.mkdirer.mkdir(roots)
}
