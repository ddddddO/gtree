### refs

- dot 言語[https://qiita.com/rubytomato@github/items/51779135bc4b77c8c20d]
- dot 言語ドキュメント[http://www.graphviz.org/documentation/]
- dot から svg へ
  `dot -Tsvg sample.dot -o sample.svg`

- graphviz インストール[https://packages.debian.org/jessie/graphviz]
  `sudo apt install graphviz -y`
  `dot -V`
- 深さ優先探索[https://qiita.com/drken/items/4a7869c5e304883f539b]

- 「Go + WebAssembly に入門して Dockerfile の依存グラフを図にしてくれるサービスを作ったので、知見とハマりポイントを共有します。」[https://qiita.com/po3rin/items/b964ad8c655c648e65ff]
- go の dot ライブラリ[https://godoc.org/github.com/gonum/graph/encoding/dot]

- livedoor のまとめランキングの個々のスレッドは、大体以下で抜けそう

```
スレッドセレクタ
->レスセレクタ

.article-body-more
->.t_b
.article-body-more
->.t_h b
.article-body-more
->.t_b
.article-body-more
->.t_b
.article-body-more
->strong span
.article-body-more
->.t_b
.article-body-more
->.t_b
.article-body-more
->.t_b
.article-body-more
->.t_b
.article-body-more
->b
.article-body
->.t_b
#articlebody
->.t_b
#articlebody
->b
.entrybody
->div#resid1 b span ...「resid<N>」が厳しい
.more_body
->.t_b
```

- キーワード
  - 抽出したスレッド、自然言語処理、感情分析、共感、各ノードの感情で色つけてリンク飛べるようにしたり
