```
├── c.go      // goからcを呼び出す試しコード
├── exe       // githubにpushしたくない実行可能ファイル置き場
├── hello.c   // コピペ。初c
└── primer    // cの基本構文を試しながら実行したコードの置き場
```

## メモ
- ハードウェア寄りの言語
- 高速
- 色々な言語の基礎
- ポインタ：アドレスを格納するための変数。メモリの節約に使える。
    - 参照渡し
    - malloc/free/NULL
- [ファイル分割](https://c-lang.sevendays-study.com/day7.html)
    - コンパイル時に、呼び出し先の関数が定義してある分割したソースファイル(.c)の指定も必要
        - e.g.) `gcc -o exe/caller primer/libCaller.c primer/library/calc.c`
