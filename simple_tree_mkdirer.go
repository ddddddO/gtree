//go:build !tinywasm

package gtree

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

var (
	// ErrExistPath is returned if the argument *gtree.Node of MkdirProgrammably function is path already exists.
	ErrExistPath = errors.New("path already exists")
)

func newMkdirerSimple(dir string, fileExtensions []string) mkdirerSimple {
	targetDir := "."
	if len(dir) != 0 {
		targetDir = dir
	}

	return &defaultMkdirerSimple{
		targetDir:      targetDir,
		fileConsiderer: newFileConsiderer(fileExtensions),
	}
}

type defaultMkdirerSimple struct {
	targetDir      string
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

func (dm *defaultMkdirerSimple) isExistRoot(roots []*Node) bool {
	for _, root := range roots {
		if _, err := os.Stat(filepath.Join(dm.targetDir, root.path())); !os.IsNotExist(err) {
			return true
		}
	}
	return false
}

func (dm *defaultMkdirerSimple) makeDirectoriesAndFiles(current *Node) error {
	if dm.fileConsiderer.isFile(current) {
		dir := strings.TrimSuffix(current.path(), current.name)
		if err := dm.mkdirAll(filepath.Join(dm.targetDir, dir)); err != nil {
			return err
		}
		return dm.mkfile(filepath.Join(dm.targetDir, current.path()))
	}

	if !current.hasChild() {
		return dm.mkdirAll(filepath.Join(dm.targetDir, current.path()))
	}

	for _, child := range current.children {
		if err := dm.makeDirectoriesAndFiles(child); err != nil {
			return err
		}
	}
	return nil
}

const permission = 0o755

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
