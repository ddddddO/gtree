package gtree

import (
	"context"
	"io"

	"golang.org/x/sync/errgroup"
)

type tree struct {
	grower   grower
	spreader spreader
	mkdirer  mkdirer
}

// 関心事は各ノードの枝の形成
type grower interface {
	grow([]*Node) (<-chan *Node, <-chan error)
	enableValidation()
}

// 関心事はtreeの出力
type spreader interface {
	spread(io.Writer, <-chan *Node) <-chan error
}

// 関心事はファイルの生成
// interfaceを使う必要はないが、grower/spreaderと合わせたいため
type mkdirer interface {
	mkdir(<-chan *Node) <-chan error
}

func newTree(conf *config) *tree {
	growerFactory := func(lastNodeFormat, intermedialNodeFormat branchFormat, dryrun bool, encode encode) grower {
		if encode != encodeDefault {
			return newNopGrower()
		}
		return newGrower(lastNodeFormat, intermedialNodeFormat, dryrun)
	}

	spreaderFactory := func(encode encode, dryrun bool, fileExtensions []string) spreader {
		if dryrun {
			return newColorizeSpreader(fileExtensions)
		}
		return newSpreader(encode)
	}

	mkdirerFactory := func(fileExtensions []string) mkdirer {
		return newMkdirer(fileExtensions)
	}

	return &tree{
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

func (t *tree) grow(roots []*Node) (<-chan *Node, <-chan error) {
	return t.grower.grow(roots)
}

func (t *tree) spread(w io.Writer, roots <-chan *Node) <-chan error {
	return t.spreader.spread(w, roots)
}

func (t *tree) mkdir(roots <-chan *Node) <-chan error {
	return t.mkdirer.mkdir(roots)
}

func handleErr(echs ...<-chan error) error {
	eg, _ := errgroup.WithContext(context.TODO())
	for i := range echs {
		i := i
		eg.Go(func() error {
			for e := range echs[i] {
				if e != nil {
					return e
				}
			}
			return nil
		})
	}
	return eg.Wait()
}
