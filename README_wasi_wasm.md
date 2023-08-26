## 1. Install WASM runtime
- https://wasmtime.dev/

## 2. Download *`gtree.wasm`*
- https://github.com/ddddddO/gtree/releases

## 3. Use in console!
### *Help*
```console
$ wasmtime gtree.wasm help
NAME:
   gtree - This CLI uses Markdown to generate directory trees and directories itself, and also verifies directories.
           The symbols that can be used in Markdown are '-', '+', '*', and '#'.

USAGE:
   gtree [global options] command [command options] [arguments...]

VERSION:
   1.9.6 / revision xxx

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

### *Template*
```console
$ wasmtime gtree.wasm template
- gtree
        - cmd
                - gtree
                        - main.go
        - testdata
                - sample1.md
                - sample2.md
        - Makefile
        - tree.go
```

### *Output*
```console
$ wasmtime gtree.wasm template | wasmtime gtree.wasm output
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

### *Mkdir*
```console
$ wasmtime gtree.wasm template | wasmtime gtree.wasm mkdir --dry-run -e .go -e .md -e Makefile
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
$ tree gtree
gtree [error opening dir]

0 directories, 0 files
$ wasmtime gtree.wasm template | wasmtime --dir=$PWD gtree.wasm mkdir -e .go -e .md -e Makefile
$ tree gtree
gtree
├── Makefile
├── cmd
│   └── gtree
│       └── main.go
├── testdata
│   ├── sample1.md
│   └── sample2.md
└── tree.go

3 directories, 5 files
```

### *Verify*
```console
$ tree example
example
├── find_pipe_programmable-gtree
│   └── main.go
├── go-list_pipe_programmable-gtree
│   └── main.go
├── like_cli
│   ├── adapter
│   │   ├── executor.go
│   │   └── indentation.go
│   └── main.go
├── noexist
│   └── xxx
└── programmable
    └── main.go

6 directories, 7 files
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
$ cat testdata/sample9.md | wasmtime --dir=$PWD gtree.wasm verify --strict
Extra paths exist:
        example/noexist
        example/noexist/xxx
Required paths does not exist:
        example/like_cli/kkk
```