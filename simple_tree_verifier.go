//go:build !wasm

package gtree

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func newVerifierSimple(dir string, strict bool) verifierSimple {
	targetDir := "."
	if len(dir) != 0 {
		targetDir = dir
	}

	return &defaultVerifierSimple{
		strict:    strict,
		targetDir: targetDir,
	}
}

type defaultVerifierSimple struct {
	strict    bool
	targetDir string
}

// cat testdata/sample9.md | sudo go run cmd/gtree/*.go verify --strict --target-dir /home/ochi/github.com/ddddddO/gtree
// cat testdata/sample10.md | sudo go run cmd/gtree/*.go verify --strict --target-dir /home/ochi/github.com/ddddddO/gtree
func (dv *defaultVerifierSimple) verify(roots []*Node) error {
	for i := range roots {
		extra, noExists, err := dv.verifyRoot(roots[i])
		if err != nil {
			return err
		}
		if err := dv.handleErr(extra, noExists); err != nil {
			return err
		}
	}

	return nil
}

func (dv *defaultVerifierSimple) verifyRoot(root *Node) ([]string, []string, error) {
	dirsMarkdown := map[string]struct{}{}
	if err := dv.recursive(root, dirsMarkdown); err != nil {
		return nil, nil, err
	}

	dirsFilesystem := map[string]struct{}{}
	extraDirs := []string{}
	rootPath := root.path()
	fileSystem := os.DirFS(filepath.Join(dv.targetDir, rootPath))
	err := fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		dir := filepath.Join(dv.targetDir, rootPath, path)

		if err != nil {
			if errors.Is(err, fs.ErrNotExist) {
				// markdown上のrootが検査対象パスに無いとエラー
				return verifyError{noExists: []string{dir}}
			}
			return err
		}

		if _, ok := dirsMarkdown[dir]; !ok {
			// Markdownに無いパスがディレクトリに有る => strictモードでエラー
			extraDirs = append(extraDirs, dir)
		}

		dirsFilesystem[dir] = struct{}{}
		return nil
	})
	if err != nil {
		return nil, nil, err
	}

	// Markdownに有るパスがディレクトリに無い時 => 通常/strictモード共通でエラー
	noExistDirs := []string{}
	for dir := range dirsMarkdown {
		if _, ok := dirsFilesystem[dir]; !ok {
			noExistDirs = append(noExistDirs, dir)
		}
	}

	return extraDirs, noExistDirs, nil
}

func (dv *defaultVerifierSimple) recursive(node *Node, dirs map[string]struct{}) error {
	dirs[filepath.Join(dv.targetDir, node.path())] = struct{}{}

	for i := range node.children {
		if err := dv.recursive(node.children[i], dirs); err != nil {
			return err
		}
	}
	return nil
}

func (dv *defaultVerifierSimple) handleErr(extra, noExists []string) error {
	if (dv.strict && len(extra) != 0) || len(noExists) != 0 {
		return verifyError{
			strict:   dv.strict,
			extra:    extra,
			noExists: noExists,
		}
	}
	return nil
}

type verifyError struct {
	strict   bool
	extra    []string
	noExists []string
}

func (v verifyError) Error() string {
	tabPrefix := func(arr []string) string {
		tmp := ""
		for i := range arr {
			tmp += fmt.Sprintf("\t%s\n", arr[i])
		}
		return tmp
	}

	msg := ""
	if v.strict && len(v.extra) != 0 {
		msg += fmt.Sprintf("Extra paths exist:\n%s", tabPrefix(v.extra))
	}
	if len(v.noExists) != 0 {
		msg += fmt.Sprintf("Required paths does not exist:\n%s", tabPrefix(v.noExists))
	}
	return strings.TrimSuffix(msg, "\n")
}
