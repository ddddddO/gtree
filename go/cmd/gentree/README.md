# gentree

## description
- cliです。
- マークダウン形式の入力から、treeコマンドで出力するような結果が出力されます。
- マークダウンはタブで整形してください。

## installation
```sh
go get github.com/ddddddO/work/go/cmd/gentree
```

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

---

- TODO: Fix bug

```sh
19:17:55 > cat testdata/sample2.md | gentree
root
├── child1
├── child2
│   └── chilchil
├── dddd
│   ├── kkkkkkk
│   │   └── lllll
            ├── ffff
            └── ppppp
│   └── oooo
└── eee
```
