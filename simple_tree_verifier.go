package gtree

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func newVerifierSimple() verifierSimple {
	return &defaultVerifierSimple{}
}

type defaultVerifierSimple struct{}

// cat testdata/sample9.md | sudo go run cmd/gtree/*.go verify
func (dv *defaultVerifierSimple) verify(roots []*Node) error {
	dirsMarkdown := map[string]struct{}{}
	for i := range roots {
		if err := dv.recursive(roots[i], dirsMarkdown); err != nil {
			return err
		}
	}

	dirsFilesystem := map[string]struct{}{}
	existDirs := []string{}
	root := "example"
	fileSystem := os.DirFS(root)
	err := fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		dir := filepath.Join(root, path)
		// fmt.Println(dir)
		if _, ok := dirsMarkdown[dir]; !ok {
			// Markdownに無いパスがディレクトリに有る => strictモードでエラー
			existDirs = append(existDirs, dir)
		}

		dirsFilesystem[dir] = struct{}{}
		return nil
	})
	if err != nil {
		return err
	}

	// Markdownに有るパスがディレクトリに無い時 => 通常/strictモード共通でエラー
	noExistDirs := []string{}
	for dir := range dirsMarkdown {
		if _, ok := dirsFilesystem[dir]; !ok {
			noExistDirs = append(noExistDirs, dir)
		}
	}

	fmt.Println("Markdownに無いパスがディレクトリに有る:")
	for _, dir := range existDirs {
		fmt.Println(dir)
	}

	fmt.Println("\nMarkdownに有るパスがディレクトリに無い:")
	for _, dir := range noExistDirs {
		fmt.Println(dir)
	}

	// Output:
	// Markdownに無いパスがディレクトリに有る:
	// example/noexist
	// example/noexist/xxx
	//
	// Markdownに有るパスがディレクトリに無い:
	// example/like_cli/kkk

	return nil
}

func (dv *defaultVerifierSimple) recursive(node *Node, dirs map[string]struct{}) error {
	// fmt.Println(node.path())
	dirs[node.path()] = struct{}{}

	for i := range node.children {
		if err := dv.recursive(node.children[i], dirs); err != nil {
			return err
		}
	}
	return nil
}
