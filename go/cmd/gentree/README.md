# gentree

## description
- cliです。
- マークダウン形式の入力から、treeコマンドで出力するような結果が出力されます。
- マークダウンはタブで整形してください。
- バグがあります。とは言っても手でツリーを0から作成するよりは幾分ましっちゃましと思います。

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

or `gentree -f testdata/sample1.md` or `cat testdata/sample1.md | gentree -f -`

```sh
23:43:06 > cat testdata/sample2.md | gentree
root
├── child1
├── child2
│   └── chilchil
├── dddd
│   ├── kkkkkkk
│   │   └── lllll
│   │       ├── ffff
│   │       └── ppppp
│   └── oooo
└── eee
```

---
- TODO: Fix bug

```sh
00:16:59 > cat testdata/sample3.md | gentree
root
├── dddd
│   └── kkkkkkk
│       └── lllll
│           ├── ffff
│   │   │   ├── LLL
│   │   │   │   └── WWWWW
│   │   │   │       └── ZZZZZ
│           └── ppppp
│               └── KKK
│                   └── 1111111
│                       └── AAAAAAA
└── eee
```
