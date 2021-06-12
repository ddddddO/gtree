# gentree

[![ci](https://github.com/ddddddO/gentree/actions/workflows/ci.yaml/badge.svg)](https://github.com/ddddddO/gentree/actions/workflows/ci.yaml)

markdown to tree.


## Demo
```sh
12:24:56 > gentree -ts << EOS
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
├── CLI or Library.
├── Given a markdown file or format, the result of the tree command is printed.
├── `gentree` does not temporarily create directories or files.
└── Create markdown file by referring to the file in the `testdata/` directory.
    ├── Hierarchy is represented by indentation.
    └── Indentation should be unified by one of the following.
        ├── Tab
        ├── Two half-width spaces（required: `-ts`）
        └── Four half-width spaces（required: `-fs`）
```

---

## Installation(CLI)
```sh
go get github.com/ddddddO/gentree/cmd/gentree
```

or, download from [here](https://github.com/ddddddO/gentree/releases).


## Usage(CLI)

```sh
19:17:07 > cat testdata/sample1.md | gentree
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
├── gentree -f testdata/sample1.md
└── cat testdata/sample1.md | gentree -f -
```

---

- Usage other than representing a directory.

```sh
16:31:42 > cat testdata/sample2.md | gentree
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
01:15:25 > cat testdata/sample4.md | gentree -ts
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
01:16:46 > cat testdata/sample5.md | gentree -fs
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

## Installation(Library)
```sh
go get github.com/ddddddO/gentree
```

## Usage(Library)

```go
package main

import (
	"bytes"
	"strings"

	"github.com/ddddddO/gentree"
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
		isTwoSpaces bool = false  // `true` when indentation is two half-width spaces
		isFourSpaces bool = false // `true` when indentation is four half-width spaces
	)

	output := gentree.Execute(buf, isTwoSpaces, isFourSpaces)
	fmt.Println(output)

// output
//
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
