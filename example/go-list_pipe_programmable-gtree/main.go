package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ddddddO/gtree"
)

// Example:
// $ cd github.com/ddddddO/gtree
// $ go list -deps ./cmd/gtree/
// internal/goarch
// unsafe
// internal/unsafeheader
// internal/abi
// internal/cpu
// ...
// $ go list -deps ./cmd/gtree/ | go run example/go-list_pipe_programmable-gtree/main.go
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

	if err := gtree.OutputFromRoot(os.Stdout, root); err != nil {
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
	// │   ├── coverage
	// │   │   └── rtcov
	// │   ├── goexperiment
	// │   ├── goos
	// │   ├── reflectlite
	// │   ├── itoa
	// │   ├── race
	// │   ├── fmtsort
	// │   ├── oserror
	// │   ├── syscall
	// │   │   ├── unix
	// │   │   └── execenv
	// │   ├── poll
	// │   ├── safefilepath
	// │   └── testlog
	// ├── unsafe
	// ├── runtime
	// │   └── internal
	// │       ├── atomic
	// │       ├── math
	// │       ├── sys
	// │       └── syscall
	// ├── errors
	// ├── math
	// │   └── bits
	// ├── unicode
	// │   ├── utf8
	// │   └── utf16
	// ├── strconv
	// ├── sync
	// │   └── atomic
	// ├── reflect
	// ├── sort
	// ├── io
	// │   ├── fs
	// │   └── ioutil
	// ├── syscall
	// ├── time
	// ├── path
	// │   └── filepath
	// ├── os
	// ├── fmt
	// ├── bytes
	// ├── strings
	// ├── bufio
	// ├── container
	// │   └── list
	// ├── context
	// ├── encoding
	// │   ├── binary
	// │   ├── base64
	// │   └── json
	// ├── github.com
	// │   ├── ddddddO
	// │   │   └── gtree
	// │   │       ├── markdown
	// │   │       └── cmd
	// │   │           └── gtree
	// │   ├── mattn
	// │   │   ├── go-isatty
	// │   │   └── go-colorable
	// │   ├── fatih
	// │   │   └── color
	// │   ├── pelletier
	// │   │   └── go-toml
	// │   │       └── v2
	// │   │           ├── internal
	// │   │           │   ├── characters
	// │   │           │   ├── danger
	// │   │           │   └── tracker
	// │   │           └── unstable
	// │   ├── russross
	// │   │   └── blackfriday
	// │   │       └── v2
	// │   ├── cpuguy83
	// │   │   └── go-md2man
	// │   │       └── v2
	// │   │           └── md2man
	// │   ├── xrash
	// │   │   └── smetrics
	// │   └── urfave
	// │       └── cli
	// │           └── v2
	// ├── golang.org
	// │   └── x
	// │       ├── sys
	// │       │   └── unix
	// │       └── sync
	// │           └── errgroup
	// ├── regexp
	// │   └── syntax
	// ├── gopkg.in
	// │   └── yaml.v3
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
