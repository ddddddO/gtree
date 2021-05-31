# gentree

[![ci](https://github.com/ddddddO/gentree/actions/workflows/ci.yaml/badge.svg)](https://github.com/ddddddO/gentree/actions/workflows/ci.yaml)


## demo

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



## description
- CLI.
- When you enter the markdown file, the tree command result is output.
- Create the markdown file by referring to the file in the `testdata/` directory.
    - The directory hierarchy is represented by indentation.
    - Indent should be unified by one of the following.
        - Tab
        - Two half-width spaces（required option: `-ts`）
        - Four half-width spaces（required option: `-fs`）

## installation
```sh
go get github.com/ddddddO/gentree
```

or, download from https://github.com/ddddddO/gentree/releases.


## how to use

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

```sh
16:31:42 > cat testdata/sample2.md | gentree
k8s_resources
├── (Tier3)
│   └── (Tier2)
│       └── (Tier1)
│           └── (Tier0)
├── Deployments
│   └── ReplicaSet
│       └── Pod
│           └── containers
├── CronJob
│   └── Job
│       └── Pod
│           └── containers
├── (empty1)
│   └── DaemonSet
│       └── Pod
│           └── containers
└── (empty2)
    └── StatefulSet
        └── Pod
            └── containers
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
