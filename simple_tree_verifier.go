package gtree

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func newVerifierSimple(dir string, strict bool) verifierSimple {
	targetDir := "."
	if len(targetDir) != 0 {
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
		exists, noExists, err := dv.verifyRoot(roots[i])
		if err != nil {
			return err
		}
		if err := dv.handleErr(exists, noExists); err != nil {
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
	existDirs := []string{}
	rootPath := root.path()
	fileSystem := os.DirFS(filepath.Join(dv.targetDir, rootPath))
	err := fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		dir := filepath.Join(dv.targetDir, rootPath, path)
		if _, ok := dirsMarkdown[dir]; !ok {
			// Markdownに無いパスがディレクトリに有る => strictモードでエラー
			existDirs = append(existDirs, dir)
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

	return existDirs, noExistDirs, nil
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

func (dv *defaultVerifierSimple) handleErr(exists, noExists []string) error {
	if (dv.strict && len(exists) != 0) || len(noExists) != 0 {
		return VerifyError{
			strict:   dv.strict,
			exists:   exists,
			noExists: noExists,
		}
	}
	return nil
}

type VerifyError struct {
	strict   bool
	exists   []string
	noExists []string
}

func (v VerifyError) Error() string {
	tabPrefix := func(arr []string) string {
		tmp := ""
		for i := range arr {
			tmp += fmt.Sprintf("\t%s\n", arr[i])
		}
		return tmp
	}

	msg := ""
	if v.strict && len(v.exists) != 0 {
		msg += fmt.Sprintf("Extra paths exist:\n%s", tabPrefix(v.exists))
	}
	if len(v.noExists) != 0 {
		msg += fmt.Sprintf("Required paths does not exist:\n%s", tabPrefix(v.noExists))
	}
	return msg
}
