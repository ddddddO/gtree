//go:build !wasm

package gtree

import (
	"os"
	"strings"
)

func newMkdirerSimple(fileExtensions []string) mkdirerSimple {
	return &defaultMkdirerSimple{
		fileConsiderer: newFileConsiderer(fileExtensions),
	}
}

type defaultMkdirerSimple struct {
	fileConsiderer *fileConsiderer
}

func (dm *defaultMkdirerSimple) mkdir(roots []*Node) error {
	if dm.isExistRoot(roots) {
		return ErrExistPath
	}

	for _, root := range roots {
		if err := dm.makeDirectoriesAndFiles(root); err != nil {
			return err
		}
	}
	return nil
}

func (*defaultMkdirerSimple) isExistRoot(roots []*Node) bool {
	for _, root := range roots {
		if _, err := os.Stat(root.path()); !os.IsNotExist(err) {
			return true
		}
	}
	return false
}

func (dm *defaultMkdirerSimple) makeDirectoriesAndFiles(current *Node) error {
	if dm.fileConsiderer.isFile(current) {
		dir := strings.TrimSuffix(current.path(), current.name)
		if err := dm.mkdirAll(dir); err != nil {
			return err
		}
		return dm.mkfile(current.path())
	}

	if !current.hasChild() {
		return dm.mkdirAll(current.path())
	}

	for _, child := range current.children {
		if err := dm.makeDirectoriesAndFiles(child); err != nil {
			return err
		}
	}
	return nil
}

func (*defaultMkdirerSimple) mkdirAll(dir string) error {
	return os.MkdirAll(dir, permission)
}

func (*defaultMkdirerSimple) mkfile(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	return f.Close()
}

var _ mkdirerSimple = (*defaultMkdirerSimple)(nil)
