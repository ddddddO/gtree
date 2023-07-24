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
	verifier verifierSimple
}

var _ tree = (*treeSimple)(nil)

func newTreeSimple(cfg *config) tree {
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

	verifierFactory := func(targetDir string, strict bool) verifierSimple {
		return newVerifierSimple(targetDir, strict)
	}

	return &treeSimple{
		grower: growerFactory(
			cfg.lastNodeFormat,
			cfg.intermedialNodeFormat,
			cfg.dryrun,
			cfg.encode,
		),
		spreader: spreaderFactory(
			cfg.encode,
			cfg.dryrun,
			cfg.fileExtensions,
		),
		mkdirer: mkdirerFactory(
			cfg.fileExtensions,
		),
		verifier: verifierFactory(
			cfg.targetDir,
			cfg.strictVerify,
		),
	}
}

func (t *treeSimple) output(w io.Writer, r io.Reader, cfg *config) error {
	rg := newRootGeneratorSimple(r, cfg.space)
	roots, err := rg.generate()
	if err != nil {
		return err
	}

	if err := t.grower.grow(roots); err != nil {
		return err
	}
	return t.spreader.spread(w, roots)
}

func (t *treeSimple) outputProgrammably(w io.Writer, root *Node, cfg *config) error {
	if err := t.grower.grow([]*Node{root}); err != nil {
		return err
	}
	return t.spreader.spread(w, []*Node{root})
}

func (t *treeSimple) mkdir(r io.Reader, cfg *config) error {
	rg := newRootGeneratorSimple(r, cfg.space)
	roots, err := rg.generate()
	if err != nil {
		return err
	}

	if err := t.grower.grow(roots); err != nil {
		return err
	}
	return t.mkdirer.mkdir(roots)
}

func (t *treeSimple) mkdirProgrammably(root *Node, cfg *config) error {
	t.grower.enableValidation()
	// when detect invalid node name, return error. process end.
	if err := t.grower.grow([]*Node{root}); err != nil {
		return err
	}
	if cfg.dryrun {
		// when detected no invalid node name, output tree.
		return t.spreader.spread(color.Output, []*Node{root})
	}
	// when detected no invalid node name, no output tree.
	return t.mkdirer.mkdir([]*Node{root})
}

func (t *treeSimple) verify(r io.Reader, cfg *config) error {
	rg := newRootGeneratorSimple(r, cfg.space)
	roots, err := rg.generate()
	if err != nil {
		return err
	}

	t.grower.enableValidation()
	if err := t.grower.grow(roots); err != nil {
		return err
	}
	return t.verifier.verify(roots)
}

func (t *treeSimple) verifyProgrammably(root *Node, cfg *config) error {
	t.grower.enableValidation()
	// when detect invalid node name, return error. process end.
	if err := t.grower.grow([]*Node{root}); err != nil {
		return err
	}
	return t.verifier.verify([]*Node{root})
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

// 関心事はディレクトリの検証
// interfaceを使う必要はないが、growerSimple/spreaderSimpleと合わせたいため
type verifierSimple interface {
	verify([]*Node) error
}
