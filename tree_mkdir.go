package gtree

import (
	"os"
	"strings"
)

// 関心事はファイルの生成
// interfaceを使う必要はないが、grower/spreaderと合わせたいため
type mkdirer interface {
	mkdir([]*Node) error
}

func newMkdirer(fileExtensions []string) mkdirer {
	return &defaultMkdirer{
		fileConsiderer: newFileConsiderer(fileExtensions),
	}
}

type defaultMkdirer struct {
	fileConsiderer *fileConsiderer
}

func (dm *defaultMkdirer) mkdir(roots []*Node) error {
	for _, root := range roots {
		if err := dm.makeDirectoriesAndFiles(root); err != nil {
			return err
		}
	}
	return nil
}

func (dm *defaultMkdirer) makeDirectoriesAndFiles(current *Node) error {
	if dm.fileConsiderer.nodeIsFile(current) {
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

const permission = 0o755

func (*defaultMkdirer) mkdirAll(dir string) error {
	return os.MkdirAll(dir, permission)
}

func (*defaultMkdirer) mkfile(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	return f.Close()
}
