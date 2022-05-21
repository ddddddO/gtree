package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ddddddO/gtree"
)

// cd github.com/ddddddO/gtree
// go list -deps ./... | go run sample/go-list_pipe_programmable-gtree/main.go
func main() {
	var (
		root = gtree.NewRoot("[All Dependencies]")
		node *gtree.Node
	)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		splited := strings.Split(line, "/")

		for i, s := range splited {
			if i == 0 {
				node = root.Add(s)
				continue
			}
			node = node.Add(s)
		}
	}

	if err := gtree.OutputProgrammably(os.Stdout, root); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// Output:
	// [All Dependencies]
	// ├── internal
	// │   ├── goarch
	// │   ├── unsafeheader
	// │   ├── abi
	// │   ├── cpu
	// │   ├── bytealg
	// │   ├── goexperiment
	// │   ├── goos
	// │   ├── reflectlite
	// │   ├── race
	// │   ├── itoa
	// │   ├── fmtsort
	// │   ├── oserror
	// │   ├── syscall
	// │   │   ├── unix
	// │   │   └── execenv
	// │   ├── poll
	// │   └── testlog
	// ├── unsafe
	// ├── runtime
	// │   └── internal
	// │       ├── atomic
	// │       ├── math
	// │       ├── sys
	// │       └── syscall
	// ├── errors
	// ├── sync
	// │   └── atomic
	// ├── io
	// │   ├── fs
	// │   └── ioutil
	// ├── unicode
	// │   ├── utf8
	// │   └── utf16
	// ├── bytes
	// ├── strings
	// ├── bufio
	// ├── encoding
	// │   ├── binary
	// │   ├── base64
	// │   └── json
	// ├── math
	// │   └── bits
	// ├── strconv
	// ├── reflect
	// ├── sort
	// ├── syscall
	// ├── time
	// ├── path
	// │   └── filepath
	// ├── os
	// ├── fmt
	// ├── golang.org
	// │   └── x
	// │       └── sys
	// │           ├── internal
	// │           │   └── unsafeheader
	// │           └── unix
	// ├── github.com
	// │   ├── mattn
	// │   │   ├── go-isatty
	// │   │   └── go-colorable
	// │   ├── fatih
	// │   │   └── color
	// │   ├── pelletier
	// │   │   └── go-toml
	// │   │       └── v2
	// │   │           └── internal
	// │   │               ├── danger
	// │   │               ├── ast
	// │   │               └── tracker
	// │   ├── ddddddO
	// │   │   └── gtree
	// │   │       ├── cmd
	// │   │       │   └── gtree
	// │   │       └── sample
	// │   │           ├── find_pipe_programmable-gtree
	// │   │           ├── go-list_pipe_programmable-gtree
	// │   │           ├── like_cli
	// │   │           │   └── adapter
	// │   │           └── programmable
	// │   ├── russross
	// │   │   └── blackfriday
	// │   │       └── v2
	// │   ├── cpuguy83
	// │   │   └── go-md2man
	// │   │       └── v2
	// │   │           └── md2man
	// │   └── urfave
	// │       └── cli
	// │           └── v2
	// ├── regexp
	// │   └── syntax
	// ├── gopkg.in
	// │   └── yaml.v2
	// ├── context
	// ├── flag
	// ├── html
	// ├── text
	// │   ├── tabwriter
	// │   └── template
	// │       └── parse
	// ├── net
	// │   └── url
	// └── log
}
