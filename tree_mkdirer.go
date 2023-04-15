//go:build !wasm

package gtree

import (
	"context"
	"os"
	"strings"
	"sync"
)

func newMkdirer(fileExtensions []string) mkdirer {
	return &defaultMkdirer{
		fileConsiderer: newFileConsiderer(fileExtensions),
	}
}

type defaultMkdirer struct {
	fileConsiderer *fileConsiderer
}

const workerMkdirNum = 10

func (dm *defaultMkdirer) mkdir(ctx context.Context, roots <-chan *Node) <-chan error {
	errc := make(chan error, 1)

	go func() {
		defer close(errc)

		wg := &sync.WaitGroup{}
		for i := 0; i < workerMkdirNum; i++ {
			wg.Add(1)
			go dm.worker(ctx, wg, roots, errc)
		}
		wg.Wait()
	}()

	return errc
}

func (dm *defaultMkdirer) worker(ctx context.Context, wg *sync.WaitGroup, roots <-chan *Node, errc chan<- error) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case root, ok := <-roots:
			if !ok {
				return
			}
			if dm.isExistRoot(root) {
				errc <- ErrExistPath
				return
			}
			if err := dm.makeDirectoriesAndFiles(root); err != nil {
				errc <- err
				return
			}
		}
	}
}

func (*defaultMkdirer) isExistRoot(root *Node) bool {
	if _, err := os.Stat(root.path()); !os.IsNotExist(err) {
		return true
	}
	return false
}

func (dm *defaultMkdirer) makeDirectoriesAndFiles(current *Node) error {
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

var _ mkdirer = (*defaultMkdirer)(nil)
