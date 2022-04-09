# gtree

[![GitHub release](https://img.shields.io/github/release/ddddddO/gtree.svg)](https://github.com/ddddddO/gtree/releases) [![Go Reference](https://pkg.go.dev/badge/github.com/ddddddO/gtree)](https://pkg.go.dev/github.com/ddddddO/gtree) [![ci](https://github.com/ddddddO/gtree/actions/workflows/ci.yaml/badge.svg)](https://github.com/ddddddO/gtree/actions/workflows/ci.yaml) [![codecov](https://codecov.io/gh/ddddddO/gtree/branch/master/graph/badge.svg?token=JLGSLF33RH)](https://codecov.io/gh/ddddddO/gtree) [![Go Report Card](https://goreportcard.com/badge/github.com/ddddddO/gtree)](https://goreportcard.com/report/github.com/ddddddO/gtree) [![License](https://img.shields.io/badge/License-BSD_2--Clause-orange.svg)](https://github.com/ddddddO/gtree/blob/master/LICENSE) [![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go#uncategorized)

Output tree🌳 or Make directories(files)📁 from Markdown or Programmatically. Provide CLI and Go Package.

⚠It is incompatible with v1.4.3 and earlier versions⚠

```
# Description
├── Output tree from markdown or programmatically.
│   ├── Output format is tree or yaml or toml or json.
│   └── Default tree.
├── Make directories from markdown or programmatically.
│   ├── It is possible to dry run.
│   └── You can use `-e` flag to make specified extensions as file.
├── Output a markdown template that can be used with either `output` subcommand or `mkdir` subcommand.
└── Provide CLI and Go Package.
```

(outputted by `cat testdata/sample0.md | gtree output --fs`)

## CLI
[read me!](https://github.com/ddddddO/gtree/blob/master/README_CLI.md#cli)


## Package(1) / like CLI
[read me!](https://github.com/ddddddO/gtree/blob/master/README_Package_1.md#package1--like-cli)


## Package(2) / generate a tree programmatically
[read me!](https://github.com/ddddddO/gtree/blob/master/README_Package_2.md#package2--generate-a-tree-programmatically)

---

## Documents
- [Markdown形式の入力からtreeを出力するCLI](https://zenn.dev/ddddddo/articles/ad97623a004496)
- [Goでtreeを表現する](https://zenn.dev/ddddddo/articles/8cd85c68763f2e)
