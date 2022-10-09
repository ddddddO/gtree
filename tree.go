package gtree

import (
	"io"
)

func newTree(conf *config, roots []*Node) *tree {
	g := newGrower(conf.lastNodeFormat, conf.intermedialNodeFormat, conf.dryrun)
	if conf.encode != encodeDefault {
		g = newNopGrower()
	}

	s := newSpreader(conf.encode)
	if conf.dryrun {
		s = newColorizeSpreader(conf.fileExtensions)
	}

	m := newMkdirer(conf.fileExtensions)

	return &tree{
		roots:    roots,
		grower:   g,
		spreader: s,
		mkdirer:  m,
	}
}

type tree struct {
	roots    []*Node
	grower   grower
	spreader spreader
	mkdirer  mkdirer
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
