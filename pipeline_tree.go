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
}

func newTreePipeline(conf *config) iTree {
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

	return &treePipeline{
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

func (t *treePipeline) output(w io.Writer, r io.Reader, conf *config) error {
	ctx, cancel := context.WithCancel(conf.ctx)
	defer cancel()

	splitStream, errcsl := split(ctx, r)
	rootStream, errcr := newRootGeneratorPipeline(conf.space).generate(ctx, splitStream)
	growStream, errcg := t.grow(ctx, rootStream)
	errcs := t.spread(ctx, w, growStream)
	return t.handlePipelineErr(ctx, errcsl, errcr, errcg, errcs)
}

func (t *treePipeline) outputProgrammably(w io.Writer, root *Node, conf *config) error {
	ctx, cancel := context.WithCancel(conf.ctx)
	defer cancel()

	rootStream := make(chan *Node)
	go func() {
		defer close(rootStream)
		rootStream <- root
	}()
	growStream, errcg := t.grow(ctx, rootStream)
	errcs := t.spread(ctx, w, growStream)
	return t.handlePipelineErr(ctx, errcg, errcs)
}

func (t *treePipeline) makedir(r io.Reader, conf *config) error {
	ctx, cancel := context.WithCancel(conf.ctx)
	defer cancel()

	splitStream, errcsl := split(ctx, r)
	rootStream, errcr := newRootGeneratorPipeline(conf.space).generate(ctx, splitStream)
	growStream, errcg := t.grow(ctx, rootStream)
	errcm := t.mkdir(ctx, growStream)
	return t.handlePipelineErr(ctx, errcsl, errcr, errcg, errcm)
}

func (t *treePipeline) makedirProgrammably(root *Node, conf *config) error {
	ctx, cancel := context.WithCancel(conf.ctx)
	defer cancel()

	rootStream := make(chan *Node)
	go func() {
		defer close(rootStream)
		rootStream <- root
	}()
	t.enableValidation()
	// when detect invalid node name, return error. process end.
	growStream, errcg := t.grow(ctx, rootStream)
	if conf.dryrun {
		// when detected no invalid node name, output tree.
		errcs := t.spread(ctx, color.Output, growStream)
		return t.handlePipelineErr(ctx, errcg, errcs)
	}
	// when detected no invalid node name, no output tree.
	errcm := t.mkdir(ctx, growStream)
	return t.handlePipelineErr(ctx, errcg, errcm)
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

func (t *treePipeline) grow(ctx context.Context, roots <-chan *Node) (<-chan *Node, <-chan error) {
	return t.grower.grow(ctx, roots)
}

func (t *treePipeline) spread(ctx context.Context, w io.Writer, roots <-chan *Node) <-chan error {
	return t.spreader.spread(ctx, w, roots)
}

func (t *treePipeline) mkdir(ctx context.Context, roots <-chan *Node) <-chan error {
	return t.mkdirer.mkdir(ctx, roots)
}

// パイプラインの全ステージで最初のエラーを返却
func (*treePipeline) handlePipelineErr(ctx context.Context, echs ...<-chan error) error {
	eg, ectx := errgroup.WithContext(ctx)
	for i := range echs {
		i := i
		eg.Go(func() error {
			for {
				select {
				case err := <-echs[i]:
					if err != nil {
						return err
					}
					return nil
				case <-ectx.Done():
					return ectx.Err()
				}
			}
		})
	}
	return eg.Wait()
}
