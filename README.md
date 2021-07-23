# gtree

[![ci](https://github.com/ddddddO/gtree/actions/workflows/ci.yaml/badge.svg)](https://github.com/ddddddO/gtree/actions/workflows/ci.yaml) [![codecov](https://codecov.io/gh/ddddddO/gtree/branch/master/graph/badge.svg?token=JLGSLF33RH)](https://codecov.io/gh/ddddddO/gtree) [![GitHub release](https://img.shields.io/github/release/ddddddO/gtree.svg)](https://github.com/ddddddO/gtree/releases)

markdown to tree.


## Demo
```sh
12:24:56 > gtree -ts << EOS
- root
  - parent_a
  - parent_b
    - child_a
      - 1
      - 2
        - a
          - 1
    - child_b
      - 1
        - a
  - parent_c
    - child_a
  - parent_d
EOS
root
├── parent_a
├── parent_b
│   ├── child_a
│   │   ├── 1
│   │   └── 2
│   │       └── a
│   │           └── 1
│   └── child_b
│       └── 1
│           └── a
├── parent_c
│   └── child_a
└── parent_d
```

## Description
```
├── CLI or Package.
├── Given a markdown file or format, the result of the tree command is printed.
├── `gtree` does not temporarily create directories or files.
└── Create markdown file by referring to the file in the `testdata/` directory.
    ├── Hierarchy is represented by indentation.
    └── Indentation should be unified by one of the following.
        ├── Tab
        ├── Two half-width spaces（required: `-ts`）
        └── Four half-width spaces（required: `-fs`）
```

---

## As CLI

### Installation
```sh
go get github.com/ddddddO/gtree/cmd/gtree
```

or, download from [here](https://github.com/ddddddO/gtree/releases).


### Usage

```sh
19:17:07 > cat testdata/sample1.md | gtree
a
├── vvv
│   └── jjj
├── kggg
│   ├── kkdd
│   └── tggg
├── edddd
│   └── orrr
└── gggg
```

#### OR
```
├── gtree -f testdata/sample1.md
└── cat testdata/sample1.md | gtree -f -
```

---

- Usage other than representing a directory.

```sh
16:31:42 > cat testdata/sample2.md | gtree
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

```sh
01:15:25 > cat testdata/sample4.md | gtree -ts
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

```sh
01:16:46 > cat testdata/sample5.md | gtree -fs
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

```sh
13:06:26 > cat testdata/sample6.md | gtree
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
---

## As Package

### Installation
```sh
go get github.com/ddddddO/gtree
```

### Usage

```go
package main

import (
	"bytes"
	"strings"

	"github.com/ddddddO/gtree"
)

func main() {
	buf := bytes.NewBufferString(strings.TrimSpace(`
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

	var (
		isTwoSpaces  bool = false // `true` when indentation is two half-width spaces
		isFourSpaces bool = false // `true` when indentation is four half-width spaces
	)

	conf := gtree.Config{
		IsTwoSpaces: isTwoSpaces,
		IsFourSpaces: isFourSpaces,
	}

	output, err := gtree.Execute(buf, conf)
	if err != nil {
		panic(err)
	}

	fmt.Println(output)

	// output

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
}

```
