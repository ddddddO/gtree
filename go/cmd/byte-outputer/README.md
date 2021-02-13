inファイルを読込み、byteスライスにし(それを標準出力に出力し)、byteスライスの一部を変換し、outファイルを出力する

inファイル
```
abc
de f
```

標準出力
```
=============
[97 98 99 13 10 100 101 32 102]
=============
```

↓

byteスライスの0番目を、97 -> 122 に変更

↓

outファイル
```
zbc
de f
```


# [ASCII](https://ja.wikipedia.org/wiki/ASCII)
97～99(a～c), 100～102(d～f), 122(z), 32(スペース) は、ASCII印字可能文字  
13(\r), 10(\n)は、ASCII制御文字(参考：https://ja.stackoverflow.com/questions/12897/%E6%94%B9%E8%A1%8C%E3%81%AE-n%E3%81%A8-r-n%E3%81%AE%E9%81%95%E3%81%84%E3%81%AF%E4%BD%95%E3%81%A7%E3%81%99%E3%81%8B)
