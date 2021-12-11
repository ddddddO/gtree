package main

import (
	"bufio"
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

	if err := gtree.ExecuteProgrammably(os.Stdout, root); err != nil {
		panic(err)
	}
	// Output:
	// [All Dependencies]
	// ├── unsafe
	// ├── internal
	// │   ├── unsafeheader
	// │   ├── abi
	// │   ├── cpu
	// │   ├── bytealg
	// │   ├── goexperiment
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
	// ├── runtime
	// │   └── internal
	// │       ├── atomic
	// │       ├── sys
	// │       └── math
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
	// ├── github.com
	// │   ├── pelletier
	// │   │   └── go-toml
	// │   │       └── v2
	// │   │           └── internal
	// │   │               ├── danger
	// │   │               ├── ast
	// │   │               └── tracker
	// │   ├── pkg
	// │   │   └── errors
	// │   └── ddddddO
	// │       └── gtree
	// │           ├── cmd
	// │           │   └── gtree
	// │           └── sample
	// │               ├── find_pipe_programmable-gtree
	// │               ├── go-list_pipe_programmable-gtree
	// │               ├── like_cli
	// │               │   └── adapter
	// │               └── programmable
	// ├── regexp
	// │   └── syntax
	// ├── gopkg.in
	// │   └── yaml.v2
	// └── flag
}
