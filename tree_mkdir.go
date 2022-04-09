package gtree

import (
	"os"
	"strings"
)

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
		if dm.needsMkfile(current) {
			dir := strings.TrimSuffix(current.getPath(), current.Name)
			if err := dm.mkdirAll(dir); err != nil {
				return err
			}
			if err := dm.mkfile(current.getPath()); err != nil {
				return err
			}
			return nil
		}

		if err := dm.mkdirAll(current.getPath()); err != nil {
			return err
		}
		return nil
	}

	for _, child := range current.Children {
		if err := dm.makeDirectoriesAndFiles(child); err != nil {
			return err
		}
	}
	return nil
}

func (dm *defaultMkdirer) needsMkfile(current *Node) bool {
	for _, e := range dm.fileExtensions {
		if strings.HasSuffix(current.Name, e) {
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
