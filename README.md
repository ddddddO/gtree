[![GitHub Pages](https://img.shields.io/badge/-GitHub_Pages-00A98F.svg?logo=github&style=flat)](https://ddddddo.github.io/gtree/)<br>
[![GitHub release](https://img.shields.io/github/release/ddddddO/gtree.svg?label=Release&color=darkcyan)](https://github.com/ddddddO/gtree/releases) [![Go Reference](https://pkg.go.dev/badge/github.com/ddddddO/gtree)](https://pkg.go.dev/github.com/ddddddO/gtree)<br>
[![License](https://img.shields.io/badge/License-BSD_2--Clause-orange.svg?color=darkcyan)](https://github.com/ddddddO/gtree/blob/master/LICENSE) [![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go#uncategorized)<br>
[![codecov](https://codecov.io/gh/ddddddO/gtree/branch/master/graph/badge.svg?token=JLGSLF33RH)](https://codecov.io/gh/ddddddO/gtree) [![Go Report Card](https://goreportcard.com/badge/github.com/ddddddO/gtree)](https://goreportcard.com/report/github.com/ddddddO/gtree) [![ci](https://github.com/ddddddO/gtree/actions/workflows/ci.yaml/badge.svg)](https://github.com/ddddddO/gtree/actions/workflows/ci.yaml)

Generate directory trees🌳 and the directories itself📁 using Markdown or Programmatically. Provide CLI, Golang library and Web (using WebAssembly).

```console
$ gtree template --description | gtree output
# Description
├── Output tree from markdown or programmatically.
│   ├── Output formats are
│   │   ├── tree
│   │   ├── yaml
│   │   ├── toml
│   │   └── json
│   └── Default tree.
├── Make directories from markdown or programmatically.
│   ├── It is possible to dry run.
│   └── You can use `-e` flag to make specified extensions as file.
├── Output a markdown template
│   └── that can be used with either `output` subcommand or `mkdir` subcommand.
└── Provide the following
    ├── CLI
    ├── Go library
    └── Web
```


# Web

https://ddddddo.github.io/gtree/

This page is that converts from Markdown to tree!<br>
This page calls a function that outputs tree. This function is a Go package compiled as WebAssembly.<br>
The symbols that can be used in Markdown are `*`, `-`, `+`, and `#`.<br>
You can change the branches like in the image below.<br>
Also, once loaded, you can enjoy offline!<br>

![](web_example.gif)

[source code](cmd/gtree-wasm/)


# CLI

## Installation

### Go (requires 1.18 or later)

```console
$ go install github.com/ddddddO/gtree/cmd/gtree@latest
```

### Homebrew

```console
$ brew install ddddddO/tap/gtree
```

### Scoop

```console
$ scoop bucket add ddddddO https://github.com/ddddddO/scoop-bucket.git
$ scoop install ddddddO/gtree
```

### deb
```console
$ export GTREE_VERSION=X.X.X
$ curl -o gtree.deb -L https://github.com/ddddddO/gtree/releases/download/v$GTREE_VERSION/gtree_$GTREE_VERSION-1_amd64.deb
$ dpkg -i gtree.deb
```

### rpm
```console
$ export GTREE_VERSION=X.X.X
$ yum install https://github.com/ddddddO/gtree/releases/download/v$GTREE_VERSION/gtree_$GTREE_VERSION-1_amd64.rpm
```

### apk
```console
$ export GTREE_VERSION=X.X.X
$ curl -o gtree.apk -L https://github.com/ddddddO/gtree/releases/download/v$GTREE_VERSION/gtree_$GTREE_VERSION-1_amd64.apk
$ apk add --allow-untrusted gtree.apk
```


### [Docker image](https://github.com/ddddddO/gtree/pkgs/container/gtree)

```console
$ docker pull ghcr.io/ddddddo/gtree:latest
$ docker run ghcr.io/ddddddo/gtree:latest template | docker run -i ghcr.io/ddddddo/gtree:latest output
gtree
├── cmd
│   └── gtree
│       └── main.go
├── testdata
│   ├── sample1.md
│   └── sample2.md
├── Makefile
└── tree.go
```

### etc

**download binary from [here](https://github.com/ddddddO/gtree/releases).**

## Usage

```console
$ gtree --help
NAME:
   gtree - This CLI outputs tree or makes directories from markdown.

USAGE:
   gtree [global options] command [command options] [arguments...]

COMMANDS:
   output, o, out     Output tree from markdown. Let's try 'gtree template | gtree output'. Output format is tree or yaml or toml or json. Default tree.
   mkdir, m           Make directories(and files) from markdown. It is possible to dry run. Let's try 'gtree template | gtree mkdir -e .go -e .md -e Makefile'.
   template, t, tmpl  Output markdown template.
   version, v         Output gtree version.
   help, h            Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help (default: false)
```

- The symbols that can be used in Markdown are `*`, `-`, `+`, and `#`.

### *Output* subcommand
```console
$ gtree output --help
NAME:
   gtree output - Output tree from markdown. Let's try 'gtree template | gtree output'. Output format is tree or yaml or toml or json. Default tree.

USAGE:
   gtree output [command options] [arguments...]

OPTIONS:
   --file value, -f value  Markdown file path. (default: stdin)
   --two-spaces, --ts      Markdown is Two Spaces indentation. (default: tab spaces)
   --four-spaces, --fs     Markdown is Four Spaces indentation. (default: tab spaces)
   --json, -j              Output JSON format. (default: stdout)
   --yaml, -y              Output YAML format. (default: stdout)
   --toml, -t              Output TOML format. (default: stdout)
   --watch, -w             Watching markdown file. (default: false)
   --help, -h              show help (default: false)
```

```console
$ gtree template
- gtree
        - cmd
                - gtree
                        - main.go
        - testdata
                - sample1.md
                - sample2.md
        - Makefile
        - tree.go
$ gtree template | gtree output
gtree
├── cmd
│   └── gtree
│       └── main.go
├── testdata
│   ├── sample1.md
│   └── sample2.md
├── Makefile
└── tree.go
```

When Markdown is indented as a tab.

```
├── gtree output -f testdata/sample1.md
├── cat testdata/sample1.md | gtree output -f -
└── cat testdata/sample1.md | gtree output
```

For 2 or 4 spaces instead of tabs, `-ts` or `-fs` is required.


- Usage other than representing a directory.

```console
$ cat testdata/sample2.md | gtree output
k8s_resources
├── (Tier3)
│   └── (Tier2)
│       └── (Tier1)
│           └── (Tier0)
├── Deployment
│   └── ReplicaSet
│       └── Pod
│           └── container(s)
├── CronJob
│   └── Job
│       └── Pod
│           └── container(s)
├── (empty)
│   └── DaemonSet
│       └── Pod
│           └── container(s)
└── (empty)
    └── StatefulSet
        └── Pod
            └── container(s)
```

---
- Two spaces indent

```console
$ cat testdata/sample4.md | gtree output -ts
a
├── i
│   ├── u
│   │   ├── k
│   │   └── kk
│   └── t
├── e
│   └── o
└── g
```

- Four spaces indent

```console
$ cat testdata/sample5.md | gtree output -fs
a
├── i
│   ├── u
│   │   ├── k
│   │   └── kk
│   └── t
├── e
│   └── o
└── g
```

- Multiple roots

```console
$ cat testdata/sample6.md | gtree output
a
├── i
│   ├── u
│   │   ├── k
│   │   └── kk
│   └── t
├── e
│   └── o
└── g
a
├── i
│   ├── u
│   │   ├── k
│   │   └── kk
│   └── t
├── e
│   └── o
└── g
```

- Output JSON

```console
$ cat testdata/sample5.md | gtree output -fs -j | jq
{
  "value": "a",
  "children": [
    {
      "value": "i",
      "children": [
        {
          "value": "u",
          "children": [
            {
              "value": "k",
              "children": null
            },
            {
              "value": "kk",
              "children": null
            }
          ]
        },
        {
          "value": "t",
          "children": null
        }
      ]
    },
    {
      "value": "e",
      "children": [
        {
          "value": "o",
          "children": null
        }
      ]
    },
    {
      "value": "g",
      "children": null
    }
  ]
}
```

- Output YAML

```console
$ cat testdata/sample5.md | gtree output -fs -y
value: a
children:
- value: i
  children:
  - value: u
    children:
    - value: k
      children: []
    - value: kk
      children: []
  - value: t
    children: []
- value: e
  children:
  - value: o
    children: []
- value: g
  children: []
```

- Output TOML

```console
$ cat testdata/sample5.md | gtree output -fs -t
value = 'a'
[[children]]
value = 'i'
[[children.children]]
value = 'u'
[[children.children.children]]
value = 'k'
children = []
[[children.children.children]]
value = 'kk'
children = []

[[children.children]]
value = 't'
children = []

[[children]]
value = 'e'
[[children.children]]
value = 'o'
children = []

[[children]]
value = 'g'
children = []

```


---
### *Mkdir* subcommand

```console
$ gtree mkdir --help
NAME:
   gtree mkdir - Make directories from markdown. It is possible to dry run. Let's try 'gtree template | gtree mkdir -e .go -e .md -e Makefile'.

USAGE:
   gtree mkdir [command options] [arguments...]

OPTIONS:
   --file value, -f value                    Markdown file path. (default: stdin)
   --two-spaces, --ts                        Markdown is Two Spaces indentation. (default: tab spaces)
   --four-spaces, --fs                       Markdown is Four Spaces indentation. (default: tab spaces)
   --dry-run, -d, --dr                       Dry run. Detects node that is invalid for directory generation. The order of the output and made directories does not always match. (default: false)
   --extension value, -e value, --ext value  Specified extension will be created as file.
   --help, -h                                show help (default: false)
```

```console
$ gtree template
- gtree
        - cmd
                - gtree
                        - main.go
        - testdata
                - sample1.md
                - sample2.md
        - Makefile
        - tree.go
$ gtree template | gtree mkdir
$ tree gtree/
gtree/
├── cmd
│   └── gtree
│       └── main.go
├── Makefile
├── testdata
│   ├── sample1.md
│   └── sample2.md
└── tree.go

8 directories, 0 files
```

#### *make directories and files*
```console
$ gtree template
- gtree
        - cmd
                - gtree
                        - main.go
        - testdata
                - sample1.md
                - sample2.md
        - Makefile
        - tree.go
$ gtree template | gtree mkdir -e .go -e .md -e Makefile
$ tree gtree/
gtree/
├── cmd
│   └── gtree
│       └── main.go
├── Makefile
├── testdata
│   ├── sample1.md
│   └── sample2.md
└── tree.go

3 directories, 5 files
```

#### *dry run*
Does not create a file and directory.

```console
$ gtree template | gtree mkdir --dry-run -e .go -e .md -e Makefile
gtree
├── cmd
│   └── gtree
│       └── main.go
├── testdata
│   ├── sample1.md
│   └── sample2.md
├── Makefile
└── tree.go

4 directories, 5 files
```

![](cli_mkdir_dryrun.png)


Any invalid file or directory name will result in an error.

```console
$ gtree mkdir --dry-run --ts <<EOS
- root
  - aa
  - bb
    - b/b
EOS
invalid node name: b/b
```

```console
$ gtree mkdir --dry-run --ts <<EOS
- /root
  - aa
  - bb
    - bb
EOS
invalid path: /root/aa
```


# Package(1) / like CLI

## Installation

Go version requires 1.18 or later.

```console
$ go get github.com/ddddddO/gtree
```

## Usage

- The symbols that can be used in Markdown are `*`, `-`, `+`, and `#`.

### *Output* func

```go
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/ddddddO/gtree"
)

func main() {
	// input Markdown is tab indented
	r1 := bytes.NewBufferString(strings.TrimSpace(`
- root
	- dddd
		- kkkkkkk
			- lllll
				- ffff
				- LLL
					- WWWWW
						- ZZZZZ
				- ppppp
					- KKK
						- 1111111
							- AAAAAAA
	- eee`))
	if err := gtree.Output(os.Stdout, r1); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// Output:
	// root
	// ├── dddd
	// │   └── kkkkkkk
	// │       └── lllll
	// │           ├── ffff
	// │           ├── LLL
	// │           │   └── WWWWW
	// │           │       └── ZZZZZ
	// │           └── ppppp
	// │               └── KKK
	// │                   └── 1111111
	// │                       └── AAAAAAA
	// └── eee

	// input Markdown is two spaces indented
	r2 := bytes.NewBufferString(strings.TrimSpace(`
- a
  - i
    - u
      - k
      - kk
    - t
  - e
    - o
  - g`))
	// When indentation is four spaces, use WithIndentFourSpaces func instead of WithIndentTwoSpaces func.
	// and, you can customize branch format.
	if err := gtree.Output(os.Stdout, r2,
		gtree.WithIndentTwoSpaces(),
		gtree.WithBranchFormatIntermedialNode("+->", ":   "),
		gtree.WithBranchFormatLastNode("+->", "    "),
	); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// Output:
	// a
	// +-> i
	// :   +-> u
	// :   :   +-> k
	// :   :   +-> kk
	// :   +-> t
	// +-> e
	// :   +-> o
	// +-> g
}

```

- You can also output JSON/YAML/TOML.
  - `gtree.WithEncodeJSON()`
  - `gtree.WithEncodeTOML()`
  - `gtree.WithEncodeYAML()`

---

### *Mkdir* func

- `gtree.Mkdir` func makes directories.
	- You can use `gtree.WithFileExtensions` func to make specified extensions as file.



# Package(2) / generate a tree programmatically

## Installation

Go version requires 1.18 or later.

```console
$ go get github.com/ddddddO/gtree
```

## Usage

### *OutputProgrammably* func

```go
package main

import (
	"fmt"
	"os"

	"github.com/ddddddO/gtree"
)

func main() {
	var root *gtree.Node = gtree.NewRoot("root")
	root.Add("child 1").Add("child 2").Add("child 3")
	var child4 *gtree.Node = root.Add("child 1").Add("child 2").Add("child 4")
	child4.Add("child 5")
	child4.Add("child 6").Add("child 7")
	root.Add("child 8")
	// you can customize branch format.
	if err := gtree.OutputProgrammably(os.Stdout, root,
		gtree.WithBranchFormatIntermedialNode("+--", ":   "),
		gtree.WithBranchFormatLastNode("+--", "    "),
	); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// Output:
	// root
	// +-- child 1
	// :   +-- child 2
	// :       +-- child 3
	// :       +-- child 4
	// :           +-- child 5
	// :           +-- child 6
	// :               +-- child 7
	// +-- child 8

	primate := preparePrimate()
	// default branch format.
	if err := gtree.OutputProgrammably(os.Stdout, primate); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// Output:
	// Primate
	// ├── Strepsirrhini
	// │   ├── Lemuriformes
	// │   │   ├── Lemuroidea
	// │   │   │   ├── Cheirogaleidae
	// │   │   │   ├── Indriidae
	// │   │   │   ├── Lemuridae
	// │   │   │   └── Lepilemuridae
	// │   │   └── Daubentonioidea
	// │   │       └── Daubentoniidae
	// │   └── Lorisiformes
	// │       ├── Galagidae
	// │       └── Lorisidae
	// └── Haplorrhini
	//     ├── Tarsiiformes
	//     │   └── Tarsiidae
	//     └── Simiiformes
	//         ├── Platyrrhini
	//         │   ├── Ceboidea
	//         │   │   ├── Atelidae
	//         │   │   └── Cebidae
	//         │   └── Pithecioidea
	//         │       └── Pitheciidae
	//         └── Catarrhini
	//             ├── Cercopithecoidea
	//             │   └── Cercopithecidae
	//             └── Hominoidea
	//                 ├── Hylobatidae
	//                 └── Hominidae
}

func preparePrimate() *gtree.Node {
	primate := gtree.NewRoot("Primate")
	strepsirrhini := primate.Add("Strepsirrhini")
	haplorrhini := primate.Add("Haplorrhini")
	lemuriformes := strepsirrhini.Add("Lemuriformes")
	lorisiformes := strepsirrhini.Add("Lorisiformes")

	lemuroidea := lemuriformes.Add("Lemuroidea")
	lemuroidea.Add("Cheirogaleidae")
	lemuroidea.Add("Indriidae")
	lemuroidea.Add("Lemuridae")
	lemuroidea.Add("Lepilemuridae")

	lemuriformes.Add("Daubentonioidea").Add("Daubentoniidae")

	lorisiformes.Add("Galagidae")
	lorisiformes.Add("Lorisidae")

	haplorrhini.Add("Tarsiiformes").Add("Tarsiidae")
	simiiformes := haplorrhini.Add("Simiiformes")

	platyrrhini := haplorrhini.Add("Platyrrhini")
	ceboidea := platyrrhini.Add("Ceboidea")
	ceboidea.Add("Atelidae")
	ceboidea.Add("Cebidae")
	platyrrhini.Add("Pithecioidea").Add("Pitheciidae")

	catarrhini := simiiformes.Add("Catarrhini")
	catarrhini.Add("Cercopithecoidea").Add("Cercopithecidae")
	hominoidea := catarrhini.Add("Hominoidea")
	hominoidea.Add("Hylobatidae")
	hominoidea.Add("Hominidae")

	return primate
}

```

- You can also output JSON 👉 [ref](https://github.com/ddddddO/gtree/blob/master/example/programmable/main.go#L61)

- You can also output YAML 👉 [ref](https://github.com/ddddddO/gtree/blob/master/example/programmable/main.go#L198)

- You can also output TOML 👉 [ref](https://github.com/ddddddO/gtree/blob/master/example/programmable/main.go#L262)

---

- The program below converts the result of `find` into a tree.
```go
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
// $ find . -type d -name .git -prune -o -type f -print
// ./config.go
// ./node_generator_test.go
// ./example/like_cli/adapter/indentation.go
// ./example/like_cli/adapter/executor.go
// ./example/like_cli/main.go
// ./example/find_pipe_programmable-gtree/main.go
// ...
// $ find . -type d -name .git -prune -o -type f -print | go run example/find_pipe_programmable-gtree/main.go
// << See "Output:" below. >>
func main() {
	var (
		root *gtree.Node
		node *gtree.Node
	)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text() // e.g.) "./example/find_pipe_programmable-gtree/main.go"
		splited := strings.Split(line, "/") // e.g.) [. example find_pipe_programmable-gtree main.go]

		for i, s := range splited {
			if root == nil {
				root = gtree.NewRoot(s) // s := "."
				node = root
				continue
			}
			if i == 0 {
				continue
			}

			tmp := node.Add(s)
			node = tmp
		}
		node = root
	}

	if err := gtree.OutputProgrammably(os.Stdout, root); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// Output:
	// .
	// ├── config.go
	// ├── node_generator_test.go
	// ├── example
	// │   ├── like_cli
	// │   │   ├── adapter
	// │   │   │   ├── indentation.go
	// │   │   │   └── executor.go
	// │   │   └── main.go
	// │   ├── find_pipe_programmable-gtree
	// │   │   └── main.go
	// │   ├── go-list_pipe_programmable-gtree
	// │   │   └── main.go
	// │   └── programmable
	// │       └── main.go
	// ├── file_considerer.go
	// ├── node.go
	// ├── node_generator.go
	// ├── .gitignore
	// ...
}

```

- Convert `go list -deps ./...` to tree 👉 [link](https://github.com/ddddddO/gtree/blob/master/example/go-list_pipe_programmable-gtree/main.go)
  - inspired by [nikolaydubina/go-recipes](https://github.com/nikolaydubina/go-recipes#readme) !

### *MkdirProgrammably* func

```go
package main

import (
	"fmt"

	"github.com/ddddddO/gtree"
)

func main() {
	primate := preparePrimate()
	if err := gtree.MkdirProgrammably(primate); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// Output(using Linux 'tree' command):
	// 22:20:43 > tree Primate/
	// Primate/
	// ├── Haplorrhini
	// │   ├── Simiiformes
	// │   │   ├── Catarrhini
	// │   │   │   ├── Cercopithecoidea
	// │   │   │   │   └── Cercopithecidae
	// │   │   │   └── Hominoidea
	// │   │   │       ├── Hominidae
	// │   │   │       └── Hylobatidae
	// │   │   └── Platyrrhini
	// │   │       ├── Ceboidea
	// │   │       │   ├── Atelidae
	// │   │       │   └── Cebidae
	// │   │       └── Pithecioidea
	// │   │           └── Pitheciidae
	// │   └── Tarsiiformes
	// │       └── Tarsiidae
	// └── Strepsirrhini
	// 	├── Lemuriformes
	// 	│   ├── Daubentonioidea
	// 	│   │   └── Daubentoniidae
	// 	│   └── Lemuroidea
	// 	│       ├── Cheirogaleidae
	// 	│       ├── Indriidae
	// 	│       ├── Lemuridae
	// 	│       └── Lepilemuridae
	// 	└── Lorisiformes
	// 		├── Galagidae
	// 		└── Lorisidae
	//
	// 28 directories, 0 files
}
```

[details](https://github.com/ddddddO/gtree/blob/master/example/programmable/main.go#L354)

---

- Make directories and files with specific extensions.

```go
package main

import (
	"fmt"

	"github.com/ddddddO/gtree"
)

func main() {
	gtreeDir := gtree.NewRoot("gtree")
	gtreeDir.Add("cmd").Add("gtree").Add("main.go")
	gtreeDir.Add("Makefile")
	testdataDir := gtreeDir.Add("testdata")
	testdataDir.Add("sample1.md")
	testdataDir.Add("sample2.md")
	gtreeDir.Add("tree.go")

	// make directories and files with specific extensions.
	if err := gtree.MkdirProgrammably(
		gtreeDir,
		gtree.WithFileExtensions([]string{".go", ".md", "Makefile"}),
	); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// Output(using Linux 'tree' command):
	// 09:44:50 > tree gtree/
	// gtree/
	// ├── cmd
	// │   └── gtree
	// │       └── main.go
	// ├── Makefile
	// ├── testdata
	// │   ├── sample1.md
	// │   └── sample2.md
	// └── tree.go
	//
	// 3 directories, 5 files
}
```



# Process


## e.g. [*gtree/tree_handler.go*](https://github.com/ddddddO/gtree/blob/master/tree_handler.go)

<image src="./process.svg" width=100%>


# Performance


> **Warning**<br>
> Depends on the environment.

- Comparison before and after software architecture was changed.
- In the case of few Roots, previous architecture is faster in execution😅
- However, for multiple Roots, execution speed tends to be faster💪✨

<image src="./performance.svg" width=100%>

<details><summary>Benchmark log</summary>

## Before pipelining
```console
$ go test -benchmem -bench Benchmark -benchtime 100x tree_handler_benchmark_test.go
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-7200U CPU @ 2.50GHz
BenchmarkOutput_singleRoot-4                 100             24948 ns/op           14400 B/op        185 allocs/op
BenchmarkOutput_tenRoots-4                   100            283504 ns/op           65320 B/op       1710 allocs/op
BenchmarkOutput_fiftyRoots-4                 100            995910 ns/op          305658 B/op       8474 allocs/op
BenchmarkOutput_hundredRoots-4               100           1649849 ns/op          631196 B/op      16927 allocs/op
BenchmarkOutput_fiveHundredsRoots-4          100           7069195 ns/op         3198760 B/op      84534 allocs/op
BenchmarkOutput_thousandRoots-4              100          15263601 ns/op         6587536 B/op     169039 allocs/op
BenchmarkOutput_3000Roots-4                  100          47383020 ns/op        19505832 B/op     507046 allocs/op
BenchmarkOutput_6000Roots-4                  100          96684490 ns/op        39203833 B/op    1014051 allocs/op
BenchmarkOutput_10000Roots-4                 100         164449890 ns/op        67676933 B/op    1690056 allocs/op
BenchmarkOutput_20000Roots-4                 100         339919674 ns/op        135144612 B/op   3380061 allocs/op
PASS
ok      command-line-arguments  68.124s
```

## After pipelining
```console
$ go test -benchmem -bench Benchmark -benchtime 100x tree_handler_benchmark_test.go
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-7200U CPU @ 2.50GHz
BenchmarkOutput_singleRoot-4                 100            130346 ns/op           24021 B/op        304 allocs/op
BenchmarkOutput_tenRoots-4                   100            324765 ns/op          120458 B/op       2301 allocs/op
BenchmarkOutput_fiftyRoots-4                 100            822572 ns/op          565455 B/op      11182 allocs/op
BenchmarkOutput_hundredRoots-4               100           1489798 ns/op         1147089 B/op      22288 allocs/op
BenchmarkOutput_fiveHundredsRoots-4          100           6546967 ns/op         5763994 B/op     111099 allocs/op
BenchmarkOutput_thousandRoots-4              100          13072156 ns/op        11705324 B/op     222112 allocs/op
BenchmarkOutput_3000Roots-4                  100          37430405 ns/op        34838155 B/op     666142 allocs/op
BenchmarkOutput_6000Roots-4                  100          75890968 ns/op        69852265 B/op    1332183 allocs/op
BenchmarkOutput_10000Roots-4                 100         125694694 ns/op        118704856 B/op   2220232 allocs/op
BenchmarkOutput_20000Roots-4                 100         242875644 ns/op        237242818 B/op   4440315 allocs/op
PASS
ok      command-line-arguments  50.962s
```

</details>

# Documents


- [Markdown形式の入力からtreeを出力するCLI](https://zenn.dev/ddddddo/articles/ad97623a004496)
- [Goでtreeを表現する](https://zenn.dev/ddddddo/articles/8cd85c68763f2e)
- [Markdown形式の入力からファイル/ディレクトリを生成するCLI/Goパッケージ](https://zenn.dev/ddddddo/articles/460d12e8c07763)
- [感想](https://scrapbox.io/ddddddo/useful_tools)


# Star History


[![Star History Chart](https://api.star-history.com/svg?repos=ddddddO/gtree&type=Date)](https://star-history.com/#ddddddO/gtree&Date)

