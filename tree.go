package gtree

import (
	"io"
)

type tree struct {
	roots []*Node

	grower   grower
	spreader spreader
	mkdirer  mkdirer
}

// 関心事は各ノードの枝の形成
type grower interface {
	grow([]*Node) error
	enableValidation()
}

// 関心事はtreeの出力
type spreader interface {
	spread(io.Writer, []*Node) error
}

// 関心事はファイルの生成
// interfaceを使う必要はないが、grower/spreaderと合わせたいため
type mkdirer interface {
	mkdir([]*Node) error
}

func newTree(conf *config, roots []*Node) *tree {
	growerFactory := func(lastNodeFormat, intermedialNodeFormat *branchFormat, dryrun bool, encode encode) grower {
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
		roots: roots,
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

func (t *tree) grow() error {
	return t.grower.grow(t.roots)
}

func (t *tree) spread(w io.Writer) error {
	return t.spreader.spread(w, t.roots)
}

func (t *tree) mkdir() error {
	return t.mkdirer.mkdir(t.roots)
}
