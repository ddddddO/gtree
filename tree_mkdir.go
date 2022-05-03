package gtree

import (
	"os"
	"strings"
)

// 関心事はファイルを生成すること
// interfaceを使う必要はないが、他と合わせるため
type mkdirer interface {
	mkdir([]*Node) error
}

func newMkdirer(fileExtensions []string) mkdirer {
	return &defaultMkdirer{
		fileExtensions: fileExtensions,
	}
}

type defaultMkdirer struct {
	fileExtensions []string
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
	if !current.hasChild() {
		if dm.needsMkfile(current.name) {
			dir := strings.TrimSuffix(current.path(), current.name)
			if err := dm.mkdirAll(dir); err != nil {
				return err
			}
			return dm.mkfile(current.path())
		}
		return dm.mkdirAll(current.path())
	}

	for _, child := range current.children {
		if err := dm.makeDirectoriesAndFiles(child); err != nil {
			return err
		}
	}
	return nil
}

func (dm *defaultMkdirer) needsMkfile(name string) bool {
	for _, e := range dm.fileExtensions {
		if strings.HasSuffix(name, e) {
			return true
		}
	}
	return false
}

const permission = 0777

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
