package gtree

import (
	"io"
)

func initializeTree(conf *config, roots []*Node) *tree {
	g := newGrower(conf.lastNodeFormat, conf.intermedialNodeFormat, conf.dryrun)
	if conf.encode != encodeDefault {
		g = newNoopGrower()
	}

	s := newSpreader(conf.encode)
	if conf.dryrun {
		s = newColorizeSpreader(conf.fileExtensions)
	}

	m := newMkdirer(conf.fileExtensions)

	return newTree(roots, g, s, m)
}

type tree struct {
	roots    []*Node
	grower   grower
	spreader spreader
	mkdirer  mkdirer
}

func newTree(
	roots []*Node,
	grower grower,
	spreader spreader,
	mkdirer mkdirer,
) *tree {
	return &tree{
		roots:    roots,
		grower:   grower,
		spreader: spreader,
		mkdirer:  mkdirer,
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
