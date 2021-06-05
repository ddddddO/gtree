# gentree

[![ci](https://github.com/ddddddO/gentree/actions/workflows/ci.yaml/badge.svg)](https://github.com/ddddddO/gentree/actions/workflows/ci.yaml)

markdown to tree.


## Demo

### input

```sh
23:23:26 > cat testdata/sample3.md
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
	- eee
```

### output

```
23:25:09 > cat testdata/sample3.md | gentree
root
├── dddd
│   └── kkkkkkk
│       └── lllll
│           ├── ffff
│           ├── LLL
│           │   └── WWWWW
│           │       └── ZZZZZ
│           └── ppppp
│               └── KKK
│                   └── 1111111
│                       └── AAAAAAA
└── eee
```



## Description
- CLI.
- When you enter the markdown file, the tree command result is output.
- Create the markdown file by referring to the file in the `testdata/` directory.
    - The directory hierarchy is represented by indentation.
    - Indent should be unified by one of the following.
        - Tab
        - Two half-width spaces（required option: `-ts`）
        - Four half-width spaces（required option: `-fs`）

## Installation
```sh
go get github.com/ddddddO/gentree
```

or, download from [here](https://github.com/ddddddO/gentree/releases).


## Usage

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

- or `gentree -f testdata/sample1.md`<br>
- or `cat testdata/sample1.md | gentree -f -`

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
