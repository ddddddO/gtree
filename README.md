<img src="heart-balloon.svg" width="77"><br>
[![GitHub Pages](https://img.shields.io/badge/-GitHub_Pages-00A98F.svg?logo=github&style=flat)](https://ddddddo.github.io/gtree/)<br>
[![GitHub release](https://img.shields.io/github/release/ddddddO/gtree.svg?label=Release&color=darkcyan)](https://github.com/ddddddO/gtree/releases) [![Go Reference](https://pkg.go.dev/badge/github.com/ddddddO/gtree)](https://pkg.go.dev/github.com/ddddddO/gtree)<br>
[![License](https://img.shields.io/badge/License-BSD_2--Clause-orange.svg?color=darkcyan)](https://github.com/ddddddO/gtree/blob/master/LICENSE) [![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go#uncategorized)<br>
[![codecov](https://codecov.io/gh/ddddddO/gtree/branch/master/graph/badge.svg?token=JLGSLF33RH)](https://codecov.io/gh/ddddddO/gtree) [![Go Report Card](https://goreportcard.com/badge/github.com/ddddddO/gtree)](https://goreportcard.com/report/github.com/ddddddO/gtree) [![ci](https://github.com/ddddddO/gtree/actions/workflows/ci.yaml/badge.svg)](https://github.com/ddddddO/gtree/actions/workflows/ci.yaml)


<img src="demo.gif"><br>

Using either Markdown or Programmatically to generate directory trees🌳 and directories🗂, and to verify directories🔍.
Provide CLI, Golang library and Web.

- **[Web](https://github.com/ddddddO/gtree#web)**
- **[CLI](https://github.com/ddddddO/gtree#cli)**
- **[Library - Like CLI](https://github.com/ddddddO/gtree#package1--like-cli)**
- **[Library - Programmable](https://github.com/ddddddO/gtree#package2--generate-a-tree-programmatically)**


# Web

https://ddddddo.github.io/gtree/

This page is that converts from Markdown to tree!<br>
This page calls a function that outputs tree. This function is a Go package compiled as WebAssembly.<br>
The symbols that can be used in Markdown are `*`, `-`, `+`, and `#`.<br>
You can change the branches like in the image below.<br>
Also, once loaded, you can enjoy offline!<br>

![](web_example.gif)

You can open it in your browser with
```console
$ gtree web
```

[source code](cmd/gtree-wasm/)


# CLI

## Installation

<pre>
<strong>Go (requires 1.18 or later)</strong>
$ go install github.com/ddddddO/gtree/cmd/gtree@latest

<strong>Homebrew</strong>
$ brew install ddddddO/tap/gtree

<strong>Scoop</strong>
$ scoop bucket add ddddddO https://github.com/ddddddO/scoop-bucket.git
$ scoop install ddddddO/gtree

<strong>deb</strong>
$ export GTREE_VERSION=X.X.X
$ curl -o gtree.deb -L https://github.com/ddddddO/gtree/releases/download/v$GTREE_VERSION/gtree_$GTREE_VERSION-1_amd64.deb
$ dpkg -i gtree.deb

<strong>rpm</strong>
$ export GTREE_VERSION=X.X.X
$ yum install https://github.com/ddddddO/gtree/releases/download/v$GTREE_VERSION/gtree_$GTREE_VERSION-1_amd64.rpm

<strong>apk</strong>
$ export GTREE_VERSION=X.X.X
$ curl -o gtree.apk -L https://github.com/ddddddO/gtree/releases/download/v$GTREE_VERSION/gtree_$GTREE_VERSION-1_amd64.apk
$ apk add --allow-untrusted gtree.apk

<s><strong><a href="https://aur.archlinux.org/packages/gtree">AUR</a></strong></s>
$ wip...

<strong><a href="https://github.com/NixOS/nixpkgs/blob/master/pkgs/tools/text/gtree/default.nix">Nix</a></strong>
$ nix-env -i gtree
or
$ nix-shell -p gtree

<strong><a href="https://github.com/macports/macports-ports/blob/master/sysutils/gtree/Portfile">MacPorts</a></strong>
$ port install gtree

<strong><a href="https://github.com/aquaproj/aqua-registry/blob/main/pkgs/ddddddO/gtree/pkg.yaml">aqua</a></strong>

<strong><a href="https://github.com/ddddddO/gtree/pkgs/container/gtree">Docker</a></strong>
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
</pre>

### etc

**download binary from [here](https://github.com/ddddddO/gtree/releases).**

## Usage

```console
$ gtree --help
NAME:
   gtree - This CLI uses Markdown to generate directory trees and directories itself, and also verifies directories.
           The symbols that can be used in Markdown are '-', '+', '*', and '#'.

USAGE:
   gtree [global options] command [command options] [arguments...]

VERSION:
   1.9.3 / revision 035aa9a

COMMANDS:
   output, o, out     Outputs tree from markdown.
                      Let's try 'gtree template | gtree output'.
   mkdir, m           Makes directories and files from markdown. It is possible to dry run.
                      Let's try 'gtree template | gtree mkdir -e .go -e .md -e Makefile'.
   verify, vf         Verifies tree structure represented in markdown by comparing it with existing directories.
                      Let's try 'gtree template | gtree verify'.
   template, t, tmpl  Outputs markdown template. Use it to try out gtree CLI.
   web, w, www        Opens "Tree Maker" in your browser and shows the URL in terminal.
   gocode, gc, code   Outputs a sample Go program calling "gtree" package.
   version, v         Prints the version.
   help, h            Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

### *Output* subcommand
```console
$ gtree output --help
NAME:
   gtree output - Outputs tree from markdown.
                  Let's try 'gtree template | gtree output'.

USAGE:
   gtree output [command options] [arguments...]

OPTIONS:
   --file value, -f value               specify the path to markdown file. (default: stdin)
   --two-spaces, --ts                   set this option when the markdown indent is 2 spaces. (default: tab spaces)
   --four-spaces, --fs                  set this option when the markdown indent is 4 spaces. (default: tab spaces)
   --massive, -m                        set this option when there are very many blocks of markdown. (default: false)
   --massive-timeout value, --mt value  set this option if you want to set a timeout. (default: 0s)
   --json, -j                           set this option when outputting JSON. (default: tree)
   --yaml, -y                           set this option when outputting YAML. (default: tree)
   --toml, -t                           set this option when outputting TOML. (default: tree)
   --watch, -w                          follow changes in markdown file. (default: false)
   --help, -h                           show help
```

#### Try it!

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


#### Usage other than representing a directory.

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
└── (empty)
    ├── DaemonSet
    │   └── Pod
    │       └── container(s)
    └── StatefulSet
        └── Pod
            └── container(s)
```

#### Two spaces indent

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

#### Four spaces indent

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

#### Multiple roots

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

#### Output JSON

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

#### Output YAML

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

#### Output TOML

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


### *Mkdir* subcommand

```console
$ gtree mkdir --help
NAME:
   gtree mkdir - Makes directories and files from markdown. It is possible to dry run.
                 Let's try 'gtree template | gtree mkdir -e .go -e .md -e Makefile'.

USAGE:
   gtree mkdir [command options] [arguments...]

OPTIONS:
   --file value, -f value  specify the path to markdown file. (default: stdin)
   --two-spaces, --ts      set this option when the markdown indent is 2 spaces. (default: tab spaces)
   --four-spaces, --fs     set this option when the markdown indent is 4 spaces. (default: tab spaces)
   --dry-run, -d, --dr     dry run. detects node that is invalid for directory generation.
      the order of the output and made directories does not always match. (default: false)
   --extension value, -e value, --ext value [ --extension value, -e value, --ext value ]  set this option if you want to create file instead of directory.
      for example, if you want to generate files with ".go" extension: "-e .go"
   --help, -h  show help
```

#### Try it!

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
invalid node name: /root
```

### *Verify* subcommand
```console
$ gtree verify --help
NAME:
   gtree verify - Verifies tree structure represented in markdown by comparing it with existing directories.
                  Let's try 'gtree template | gtree verify'.

USAGE:
   gtree verify [command options] [arguments...]

OPTIONS:
   --file value, -f value  specify the path to markdown file. (default: stdin)
   --two-spaces, --ts      set this option when the markdown indent is 2 spaces. (default: tab spaces)
   --four-spaces, --fs     set this option when the markdown indent is 4 spaces. (default: tab spaces)
   --target-dir value      set this option if you want to specify the directory you want to verify. (default: current directory)
   --strict                set this option if you want strict directory match validation. (default: non strict)
   --help, -h              show help
```

#### Try it!

```console
~/github.com/ddddddO/gtree
$ cat testdata/sample9.md
- example
        - find_pipe_programmable-gtree
                - main.go
        - go-list_pipe_programmable-gtree
                - main.go
        - like_cli
                - adapter
                        - executor.go
                        - indentation.go
                - main.go
                - kkk
        - programmable
                - main.go
~/github.com/ddddddO/gtree
$ cat testdata/sample9.md | gtree verify --strict
Extra paths exist:
        example/noexist
        example/noexist/xxx
Required paths does not exist:
        example/like_cli/kkk
```


# Package(1) / like CLI

## Installation

Go version requires 1.18 or later.

```console
$ go get github.com/ddddddO/gtree
```

## Usage

The symbols that can be used in Markdown are `*`, `-`, `+`, and `#`.

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

#### You can also output JSON/YAML/TOML.
- `gtree.WithEncodeJSON()`
- `gtree.WithEncodeTOML()`
- `gtree.WithEncodeYAML()`


### *Mkdir* func

#### `gtree.Mkdir` func makes directories.

You can use `gtree.WithFileExtensions` func to make specified extensions as file.


### *Verify* func

#### `gtree.Verify` func verifies directories.

You can use `gtree.WithTargetDir` func / `gtree.WithStrictVerify` func.

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

#### The program below converts the result of `find` into a tree.

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
		line := scanner.Text()              // e.g.) "./example/find_pipe_programmable-gtree/main.go"
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

The above Go program can be output with the command below.

```console
$ gtree gocode
```

#### Convert `go list -deps ./...` to tree 👉 [sample program](https://github.com/ddddddO/gtree/blob/master/example/go-list_pipe_programmable-gtree/main.go)

The above Go program can be output with the command below.

```console
$ gtree gocode --godeps-to-tree
```

inspired by [nikolaydubina/go-recipes](https://github.com/nikolaydubina/go-recipes#readme) !


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


#### Make directories and files with specific extensions.

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

### *VerifyProgrammably* func

You can use `gtree.WithTargetDir` func / `gtree.WithStrictVerify` func.

# Process

> **Note**<br>
> This process is for the Massive Roots mode.

## e.g. [*gtree/pipeline_tree.go*](https://github.com/ddddddO/gtree/blob/master/pipeline_tree.go)

<image src="./process.svg" width=100%>


# Performance

> **Warning**<br>
> Depends on the environment.

- Comparison simple implementation and pipeline implementation.
- In the case of few Roots, simple implementation is faster in execution!
	- Use this one by default.
- However, for multiple Roots, pipeline implementation execution speed tends to be faster💪✨
	- In the CLI, it is available by specifying `--massive`.
	- In the Go program, it is available by specifying `WithMassive` func.

<image src="./performance.svg" width=100%>

<details><summary>Benchmark log</summary>

## Simple implementation
```console
11:19:22 > go test -benchmem -bench Benchmark -benchtime 100x benchmark_simple_test.go
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-7200U CPU @ 2.50GHz
BenchmarkOutput_singleRoot-4                 100             35375 ns/op           13856 B/op        171 allocs/op
BenchmarkOutput_tenRoots-4                   100            200540 ns/op           72920 B/op       1597 allocs/op
BenchmarkOutput_fiftyRoots-4                 100            730156 ns/op          569851 B/op       7919 allocs/op
BenchmarkOutput_hundredRoots-4               100           1706493 ns/op         1714260 B/op      15820 allocs/op
BenchmarkOutput_fiveHundredsRoots-4          100          16412090 ns/op        32245140 B/op      79022 allocs/op
BenchmarkOutput_thousandRoots-4              100          55142492 ns/op        120929674 B/op    158025 allocs/op
BenchmarkOutput_3000Roots-4                  100         489121246 ns/op        1035617527 B/op   474029 allocs/op
BenchmarkOutput_6000Roots-4                  100        1613641261 ns/op        4087694372 B/op   948033 allocs/op
BenchmarkOutput_10000Roots-4                 100        3913090646 ns/op        11293191221 B/op         1580035 allocs/op
PASS
ok      command-line-arguments  614.944s
```

## Pipeline implementation
```console
11:29:43 > go test -benchmem -bench Benchmark -benchtime 100x benchmark_pipeline_test.go
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-7200U CPU @ 2.50GHz
BenchmarkOutput_pipeline_singleRoot-4                100            188706 ns/op           24236 B/op        300 allocs/op
BenchmarkOutput_pipeline_tenRoots-4                  100            367758 ns/op          115970 B/op       2186 allocs/op
BenchmarkOutput_pipeline_fiftyRoots-4                100            947879 ns/op          542188 B/op      10592 allocs/op
BenchmarkOutput_pipeline_hundredRoots-4              100           1711537 ns/op         1099636 B/op      21094 allocs/op
BenchmarkOutput_pipeline_fiveHundredsRoots-4         100           6892261 ns/op         5524905 B/op     105107 allocs/op
BenchmarkOutput_pipeline_thousandRoots-4             100          13100335 ns/op        11225942 B/op     210115 allocs/op
BenchmarkOutput_pipeline_3000Roots-4                 100          40694497 ns/op        33399766 B/op     630142 allocs/op
BenchmarkOutput_pipeline_6000Roots-4                 100          85807944 ns/op        66974524 B/op    1260171 allocs/op
BenchmarkOutput_pipeline_10000Roots-4                100         151486713 ns/op        113908462 B/op   2100208 allocs/op
PASS
ok      command-line-arguments  30.670s
```

</details>

# Diagram
<image src="./diagram.svg" width=100%>

# Documents

- [Markdown形式の入力からtreeを出力するCLI](https://zenn.dev/ddddddo/articles/ad97623a004496)
- [Goでtreeを表現する](https://zenn.dev/ddddddo/articles/8cd85c68763f2e)
- [Markdown形式の入力からファイル/ディレクトリを生成するCLI/Goパッケージ](https://zenn.dev/ddddddo/articles/460d12e8c07763)
- [感想](https://scrapbox.io/ddddddo/useful_tools)


# Star History

[![Star History Chart](https://api.star-history.com/svg?repos=ddddddO/gtree&type=Date)](https://star-history.com/#ddddddO/gtree&Date)
