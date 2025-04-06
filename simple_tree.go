//go:build !tinywasm

package gtree

import (
	"io"
	"iter"

	"github.com/fatih/color"
)

type treeSimple struct {
	grower       growerSimple
	spreader     spreaderSimple
	mkdirer      mkdirerSimple
	verifier     verifierSimple
	growSpreader growSpreaderSimple
	walker       walkerSimple
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

	mkdirerFactory := func(targetDir string, fileExtensions []string) mkdirerSimple {
		return newMkdirerSimple(targetDir, fileExtensions)
	}

	verifierFactory := func(targetDir string, strict bool) verifierSimple {
		return newVerifierSimple(targetDir, strict)
	}

	growSpreaderFactory := func(lastNodeFormat, intermedialNodeFormat branchFormat) growSpreaderSimple {
		return newGrowSpreaderSimple(lastNodeFormat, intermedialNodeFormat)
	}

	walkerFactory := func() walkerSimple {
		return newWalkerSimple()
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
			cfg.targetDir,
			cfg.fileExtensions,
		),
		verifier: verifierFactory(
			cfg.targetDir,
			cfg.strictVerify,
		),
		growSpreader: growSpreaderFactory(
			cfg.lastNodeFormat,
			cfg.intermedialNodeFormat,
		),
		walker: walkerFactory(),
	}
}

func (t *treeSimple) output(w io.Writer, r io.Reader, cfg *config) error {
	// ベンチマークを取るための
	if cfg.noUseIterOfSimpleOutput {
		roots, err := newRootGeneratorSimple(r).generate()
		if err != nil {
			return err
		}
		if err := t.grower.grow(roots); err != nil {
			return err
		}
		return t.spreader.spread(w, roots)
	}

	for err := range t.spreader.spreadIter(w, t.grower.growIter(newRootGeneratorSimple(r).generateIter())) {
		if err != nil {
			return err
		}
	}
	return nil
}

func (t *treeSimple) outputProgrammably(w io.Writer, root *Node, cfg *config) error {
	if cfg.encode != encodeDefault {
		if err := t.grower.grow([]*Node{root}); err != nil {
			return err
		}
		return t.spreader.spread(w, []*Node{root})
	}

	// JSONなどのフォーマットに変えずに出力する場合は、枝の形成と出力を分けない
	return t.growSpreader.growAndSpread(w, []*Node{root})
}

func (t *treeSimple) mkdir(r io.Reader, cfg *config) error {
	roots, err := newRootGeneratorSimple(r).generate()
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
	roots, err := newRootGeneratorSimple(r).generate()
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

func (t *treeSimple) walk(r io.Reader, callback func(*WalkerNode) error, cfg *config) error {
	roots, err := newRootGeneratorSimple(r).generate()
	if err != nil {
		return err
	}

	if err := t.grower.grow(roots); err != nil {
		return err
	}
	return t.walker.walk(roots, callback)
}

func (t *treeSimple) walkProgrammably(root *Node, callback func(*WalkerNode) error, cfg *config) error {
	if err := t.grower.grow([]*Node{root}); err != nil {
		return err
	}
	return t.walker.walk([]*Node{root}, callback)
}

func (t *treeSimple) walkIterProgrammably(root *Node, cfg *config) iter.Seq2[*WalkerNode, error] {
	return func(yield func(*WalkerNode, error) bool) {
		if err := t.grower.grow([]*Node{root}); err != nil {
			yield(nil, err)
			return
		}

		walkIter := t.walker.walkIter(root)
		next, stop := iter.Pull2(walkIter)
		defer stop()

		for {
			wn, err, ok := next()
			if !ok {
				return
			}
			if err != nil {
				yield(nil, err)
				return
			}
			if !yield(wn, nil) {
				return
			}
		}
	}
}

// 関心事は各ノードの枝の形成
type growerSimple interface {
	grow([]*Node) error
	growIter(iter.Seq2[*Node, error]) iter.Seq2[*Node, error]
	enableValidation()
}

// 関心事はtreeの出力
type spreaderSimple interface {
	spread(io.Writer, []*Node) error
	spreadIter(io.Writer, iter.Seq2[*Node, error]) iter.Seq[error]
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

// TODO: このコミット辺りのリファクタリング
// 関心事は枝の形成と出力
// 枝の組み立てと出力を同じloop内でしないと、例えばこれまで通り 1.grower -> 2.spreader の処理順だと、1が遅すぎる場合利用者からするといつ出力されるのか？となるし、処理回数的には無駄なため
type growSpreaderSimple interface {
	growAndSpread(io.Writer, []*Node) error
}

type walkerSimple interface {
	walk([]*Node, func(*WalkerNode) error) error
	walkIter(*Node) iter.Seq2[*WalkerNode, error]
}
