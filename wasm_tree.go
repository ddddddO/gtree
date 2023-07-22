//go:build wasm

package gtree

import (
	"io"
)

type tree struct {
	roots []*Node

	grower   grower
	spreader spreader
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

func newTree(cfg *config, roots []*Node) *tree {
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

	return &tree{
		roots: roots,
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
	}
}

func (t *tree) grow() error {
	return t.grower.grow(t.roots)
}

func (t *tree) spread(w io.Writer) error {
	return t.spreader.spread(w, t.roots)
}
