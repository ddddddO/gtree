//go:build !wasm

package gtree

import (
	"context"
	"io"

	"github.com/fatih/color"
	"golang.org/x/sync/errgroup"
)

type treePipeline struct {
	grower   growerPipeline
	spreader spreaderPipeline
	mkdirer  mkdirerPipeline
	verifier verifierPipeline
}

var _ tree = (*treePipeline)(nil)

func newTreePipeline(cfg *config) tree {
	growerFactory := func(lastNodeFormat, intermedialNodeFormat branchFormat, dryrun bool, encode encode) growerPipeline {
		if encode != encodeDefault {
			return newNopGrowerPipeline()
		}
		return newGrowerPipeline(lastNodeFormat, intermedialNodeFormat, dryrun)
	}

	spreaderFactory := func(encode encode, dryrun bool, fileExtensions []string) spreaderPipeline {
		if dryrun {
			return newColorizeSpreaderPipeline(fileExtensions)
		}
		return newSpreaderPipeline(encode)
	}

	mkdirerFactory := func(fileExtensions []string) mkdirerPipeline {
		return newMkdirerPipeline(fileExtensions)
	}

	verifierFactory := func(targetDir string, strict bool) verifierPipeline {
		return newVerifierPipeline(targetDir, strict)
	}

	return &treePipeline{
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

func (t *treePipeline) output(w io.Writer, r io.Reader, cfg *config) error {
	ctx, cancel := context.WithCancel(cfg.ctx)
	defer cancel()

	splitStream, errcsl := split(ctx, r)
	rootStream, errcr := newRootGeneratorPipeline(cfg.space).generate(ctx, splitStream)
	growStream, errcg := t.grower.grow(ctx, rootStream)
	errcs := t.spreader.spread(ctx, w, growStream)
	return t.handlePipelineErr(ctx, errcsl, errcr, errcg, errcs)
}

func (t *treePipeline) outputProgrammably(w io.Writer, root *Node, cfg *config) error {
	ctx, cancel := context.WithCancel(cfg.ctx)
	defer cancel()

	rootStream := make(chan *Node)
	go func() {
		defer close(rootStream)
		rootStream <- root
	}()
	growStream, errcg := t.grower.grow(ctx, rootStream)
	errcs := t.spreader.spread(ctx, w, growStream)
	return t.handlePipelineErr(ctx, errcg, errcs)
}

func (t *treePipeline) mkdir(r io.Reader, cfg *config) error {
	ctx, cancel := context.WithCancel(cfg.ctx)
	defer cancel()

	splitStream, errcsl := split(ctx, r)
	rootStream, errcr := newRootGeneratorPipeline(cfg.space).generate(ctx, splitStream)
	growStream, errcg := t.grower.grow(ctx, rootStream)
	errcm := t.mkdirer.mkdir(ctx, growStream)
	return t.handlePipelineErr(ctx, errcsl, errcr, errcg, errcm)
}

func (t *treePipeline) mkdirProgrammably(root *Node, cfg *config) error {
	ctx, cancel := context.WithCancel(cfg.ctx)
	defer cancel()

	rootStream := make(chan *Node)
	go func() {
		defer close(rootStream)
		rootStream <- root
	}()
	t.grower.enableValidation()
	// when detect invalid node name, return error. process end.
	growStream, errcg := t.grower.grow(ctx, rootStream)
	if cfg.dryrun {
		// when detected no invalid node name, output tree.
		errcs := t.spreader.spread(ctx, color.Output, growStream)
		return t.handlePipelineErr(ctx, errcg, errcs)
	}
	// when detected no invalid node name, no output tree.
	errcm := t.mkdirer.mkdir(ctx, growStream)
	return t.handlePipelineErr(ctx, errcg, errcm)
}

func (t *treePipeline) verify(r io.Reader, cfg *config) error {
	ctx, cancel := context.WithCancel(cfg.ctx)
	defer cancel()

	t.grower.enableValidation()
	splitStream, errcsl := split(ctx, r)
	rootStream, errcr := newRootGeneratorPipeline(cfg.space).generate(ctx, splitStream)
	growStream, errcg := t.grower.grow(ctx, rootStream)
	errcv := t.verifier.verify(ctx, growStream)
	return t.handlePipelineErr(ctx, errcsl, errcr, errcg, errcv)
}

func (t *treePipeline) verifyProgrammably(root *Node, cfg *config) error {
	ctx, cancel := context.WithCancel(cfg.ctx)
	defer cancel()

	rootStream := make(chan *Node)
	go func() {
		defer close(rootStream)
		rootStream <- root
	}()
	t.grower.enableValidation()
	// when detect invalid node name, return error. process end.
	growStream, errcg := t.grower.grow(ctx, rootStream)
	// when detected no invalid node name, no output tree.
	errcv := t.verifier.verify(ctx, growStream)
	return t.handlePipelineErr(ctx, errcg, errcv)
}

// 関心事は各ノードの枝の形成
type growerPipeline interface {
	grow(context.Context, <-chan *Node) (<-chan *Node, <-chan error)
	enableValidation()
}

// 関心事はtreeの出力
type spreaderPipeline interface {
	spread(context.Context, io.Writer, <-chan *Node) <-chan error
}

// 関心事はファイルの生成
// interfaceを使う必要はないが、growerPipeline/spreaderPipelineと合わせたいため
type mkdirerPipeline interface {
	mkdir(context.Context, <-chan *Node) <-chan error
}

// 関心事はディレクトリの検証
// interfaceを使う必要はないが、growerPipeline/spreaderPipelineと合わせたいため
type verifierPipeline interface {
	verify(context.Context, <-chan *Node) <-chan error
}

// パイプラインの全ステージで最初のエラーを返却
func (*treePipeline) handlePipelineErr(ctx context.Context, echs ...<-chan error) error {
	eg, ectx := errgroup.WithContext(ctx)
	for i := range echs {
		i := i
		eg.Go(func() error {
			select {
			case err, ok := <-echs[i]:
				if !ok {
					return nil
				}
				if err != nil {
					return err
				}
			case <-ectx.Done():
				return ectx.Err()
			}
			return nil
		})
	}
	return eg.Wait()
}
