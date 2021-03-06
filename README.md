# gtree

[![GitHub release](https://img.shields.io/github/release/ddddddO/gtree.svg)](https://github.com/ddddddO/gtree/releases) [![Go Reference](https://pkg.go.dev/badge/github.com/ddddddO/gtree)](https://pkg.go.dev/github.com/ddddddO/gtree) [![ci](https://github.com/ddddddO/gtree/actions/workflows/ci.yaml/badge.svg)](https://github.com/ddddddO/gtree/actions/workflows/ci.yaml) [![codecov](https://codecov.io/gh/ddddddO/gtree/branch/master/graph/badge.svg?token=JLGSLF33RH)](https://codecov.io/gh/ddddddO/gtree) [![Go Report Card](https://goreportcard.com/badge/github.com/ddddddO/gtree)](https://goreportcard.com/report/github.com/ddddddO/gtree) [![License](https://img.shields.io/badge/License-BSD_2--Clause-orange.svg)](https://github.com/ddddddO/gtree/blob/master/LICENSE) [![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go#uncategorized)

Output treeπ³ or Make directories(files)π from Markdown or Programmatically. Provide CLI and Go Package.

```
# Description
βββ Output tree from markdown or programmatically.
β   βββ Output format is tree or yaml or toml or json.
β   βββ Default tree.
βββ Make directories from markdown or programmatically.
β   βββ It is possible to dry run.
β   βββ You can use `-e` flag to make specified extensions as file.
βββ Output a markdown template that can be used with either `output` subcommand or `mkdir` subcommand.
βββ Provide CLI and Go Package.
```

(outputted by `cat testdata/sample0.md | gtree output --fs`)

## Package(1) / like CLI
[read me!](https://github.com/ddddddO/gtree/blob/master/README_Package_1.md#package1--like-cli)


## Package(2) / generate a tree programmatically
[read me!](https://github.com/ddddddO/gtree/blob/master/README_Package_2.md#package2--generate-a-tree-programmatically)

## CLI
### Installation

Go version requires 1.18 or later.

```console
$ go install github.com/ddddddO/gtree/cmd/gtree@latest
```

or using Homebrew.
```console
$ brew install ddddddO/tap/gtree
```

or [docker image](https://github.com/ddddddO/gtree/pkgs/container/gtree).
```console
$ docker pull ghcr.io/ddddddo/gtree:latest
$ docker run ghcr.io/ddddddo/gtree:latest template | docker run -i ghcr.io/ddddddo/gtree:latest output
gtree
βββ cmd
β   βββ gtree
β       βββ main.go
βββ testdata
β   βββ sample1.md
β   βββ sample2.md
βββ makefile
βββ tree.go
```

**or, download binary from [here](https://github.com/ddddddO/gtree/releases).**

### Usage

```console
23:47:15 > gtree --help
NAME:
   gtree - This CLI outputs tree or makes directories from markdown.

USAGE:
   gtree [global options] command [command options] [arguments...]

COMMANDS:
   output, o, out     Output tree from markdown. Let's try 'gtree template | gtree output'. Output format is tree or yaml or toml or json. Default tree.
   mkdir, m           Make directories(and files) from markdown. It is possible to dry run. Let's try 'gtree template | gtree mkdir -e .go -e .md -e makefile'.
   template, t, tmpl  Output markdown template.
   version, v         Output gtree version.
   help, h            Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help (default: false)
```

#### *Output* subcommand
```console
19:43:41 > gtree output --help
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
19:45:03 > gtree template
- gtree
        - cmd
                - gtree
                        - main.go
        - testdata
                - sample1.md
                - sample2.md
        - makefile
        - tree.go
19:47:08 > gtree template | gtree output
gtree
βββ cmd
β   βββ gtree
β       βββ main.go
βββ testdata
β   βββ sample1.md
β   βββ sample2.md
βββ makefile
βββ tree.go
```

When Markdown is indented as a tab.

```
βββ gtree output -f testdata/sample1.md
βββ cat testdata/sample1.md | gtree output -f -
βββ cat testdata/sample1.md | gtree output
```

For 2 or 4 spaces instead of tabs, `-ts` or `-fs` is required.


<details>
<summary>More details</summary>

- Usage other than representing a directory.

```console
16:31:42 > cat testdata/sample2.md | gtree output
k8s_resources
βββ (Tier3)
β   βββ (Tier2)
β       βββ (Tier1)
β           βββ (Tier0)
βββ Deployment
β   βββ ReplicaSet
β       βββ Pod
β           βββ container(s)
βββ CronJob
β   βββ Job
β       βββ Pod
β           βββ container(s)
βββ (empty)
β   βββ DaemonSet
β       βββ Pod
β           βββ container(s)
βββ (empty)
    βββ StatefulSet
        βββ Pod
            βββ container(s)
```

---
- Two spaces indent

```console
01:15:25 > cat testdata/sample4.md | gtree output -ts
a
βββ i
β   βββ u
β   β   βββ k
β   β   βββ kk
β   βββ t
βββ e
β   βββ o
βββ g
```

- Four spaces indent

```console
01:16:46 > cat testdata/sample5.md | gtree output -fs
a
βββ i
β   βββ u
β   β   βββ k
β   β   βββ kk
β   βββ t
βββ e
β   βββ o
βββ g
```

- Multiple roots

```console
13:06:26 > cat testdata/sample6.md | gtree output
a
βββ i
β   βββ u
β   β   βββ k
β   β   βββ kk
β   βββ t
βββ e
β   βββ o
βββ g
a
βββ i
β   βββ u
β   β   βββ k
β   β   βββ kk
β   βββ t
βββ e
β   βββ o
βββ g
```

- Output JSON

```console
22:40:31 > cat testdata/sample5.md | gtree output -fs -j | jq
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
13:15:59 > cat testdata/sample5.md | gtree output -fs -y
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
13:16:32 > cat testdata/sample5.md | gtree output -fs -t
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

</details>

---
#### *Mkdir* subcommand

```console
22:19:20 > gtree mkdir --help
NAME:
   gtree mkdir - Make directories from markdown. It is possible to dry run. Let's try 'gtree template | gtree mkdir -e .go -e .md -e makefile'.

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
22:25:18 > gtree template
- gtree
        - cmd
                - gtree
                        - main.go
        - testdata
                - sample1.md
                - sample2.md
        - makefile
        - tree.go
22:26:06 > gtree template | gtree mkdir
22:26:14 > tree gtree/
gtree/
βββ cmd
β   βββ gtree
β       βββ main.go
βββ makefile
βββ testdata
β   βββ sample1.md
β   βββ sample2.md
βββ tree.go

8 directories, 0 files
```

##### *make directories and files*
```console
22:15:59 > gtree template
- gtree
        - cmd
                - gtree
                        - main.go
        - testdata
                - sample1.md
                - sample2.md
        - makefile
        - tree.go
22:16:13 > gtree template | gtree mkdir -e .go -e .md -e makefile
22:16:19 > tree gtree/
gtree/
βββ cmd
β   βββ gtree
β       βββ main.go
βββ makefile
βββ testdata
β   βββ sample1.md
β   βββ sample2.md
βββ tree.go

3 directories, 5 files
```

##### *dry run*
Does not create a file and directory.

```console
12:40:49 > gtree template | gtree mkdir --dry-run -e .go -e .md -e makefile
gtree
βββ cmd
β   βββ gtree
β       βββ main.go
βββ testdata
β   βββ sample1.md
β   βββ sample2.md
βββ makefile
βββ tree.go

4 directories, 5 files
```

β colored output<br>
![](cli_mkdir_dryrun.png)


Any invalid file or directory name will result in an error.

```console
23:20:04 > gtree mkdir --dry-run --ts <<EOS
- root
  - aa
  - bb
    - b/b
EOS
invalid node name: b/b
```

```console
23:27:27 > gtree mkdir --dry-run --ts <<EOS
- /root
  - aa
  - bb
    - bb
EOS
invalid path: /root/aa
```

---

## Documents
- [Markdownε½’εΌγ?ε₯εγγtreeγεΊεγγCLI](https://zenn.dev/ddddddo/articles/ad97623a004496)
- [Goγ§treeγθ‘¨ηΎγγ](https://zenn.dev/ddddddo/articles/8cd85c68763f2e)
- [Markdownε½’εΌγ?ε₯εγγγγ‘γ€γ«/γγ£γ¬γ―γγͺγηζγγCLI/Goγγγ±γΌγΈ](https://zenn.dev/ddddddo/articles/460d12e8c07763)
