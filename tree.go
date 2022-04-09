package gtree

import (
	"io"
)

// Output outputs a tree to w with r as Markdown format input.
func Output(w io.Writer, r io.Reader, optFns ...OptFn) error {
	conf, err := newConfig(optFns...)
	if err != nil {
		return err
	}
	rs, err := generateRoots(r, conf.space)
	if err != nil {
		return err
	}

	tree := initializeTree(conf, rs)
	if err := tree.grow(); err != nil {
		return err
	}
	return tree.spread(w)
}

// Mkdir makes directories.
func Mkdir(r io.Reader, optFns ...OptFn) error {
	conf, err := newConfig(optFns...)
	if err != nil {
		return err
	}
	rs, err := generateRoots(r, conf.space)
	if err != nil {
		return err
	}

	tree := initializeTree(conf, rs)
	if err := tree.grow(); err != nil {
		return err
	}
	return tree.mkdir()
}

func initializeTree(conf *config, rs []*Node) *tree {
	g := newGrower(conf.encode, conf.lastNodeFormat, conf.intermedialNodeFormat, conf.dryrun)
	s := newSpreader(conf.encode)
	m := newMkdirer(conf.fileExtensions)
	return newTree(rs, g, s, m)
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

// TODO: メソッド名見直す
func (t *tree) enableValidation() {
	t.grower.setDryRun(true)
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
