[OPENLOGI Advent Calendar 2023](https://qiita.com/advent-calendar/2023/openlogi) 2日目の記事です。
https://qiita.com/advent-calendar/2023/openlogi

この記事では、私が育てている OSS にまつわる歴史・嬉しかったことを書こうと思います。

# 前置き
Linux や Windows で `tree` コマンドがあるかと思いますが、大雑把に言うと、そのコマンドの出力結果と同等の出力を得ることができる CLI・ライブラリ・Web の開発をしています（[gtree](https://github.com/ddddddO/gtree)）。
この章では、 CLI・Web の説明は割愛し、ライブラリに絞って概説します（もし興味が出た方は、リポジトリの README を参照いただければと思います）。

ライブラリですが、主に以下の関数を利用できます。
|No|関数名|機能|
|--|--|--|
|1|**OutputProgrammably**|ツリーを出力できます|
|2|MkdirProgrammably|ディレクトリを生成できます|
|3|VerifyProgrammably|ディレクトリを検証できます|
|4|WalkProgrammably|木を構成する各ノードに対してユーザー定義関数を実行できます|

4つの関数はいずれも、ユーザーの Go プログラム内で木を構成してもらって、その木に対して何らか処理をするというものです。
No1の **OutputProgrammably** 関数を使ったサンプルはこのような感じです。

:::details サンプル

```go:main.go
package main

import (
	"fmt"
	"os"

	"github.com/ddddddO/gtree"
)

func main() {
	var root *gtree.Node = gtree.NewRoot("root")
	root.Add("child 1").Add("child 2")
	root.Add("child 1").Add("child 3")
	child4 := root.Add("child 4")

	var child7 *gtree.Node = child4.Add("child 5").Add("child 6").Add("child 7")
	child7.Add("child 8")

	if err := gtree.OutputProgrammably(os.Stdout, root); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// Output:
	// root
	// ├── child 1
	// │   ├── child 2
	// │   └── child 3
	// └── child 4
	//     └── child 5
	//         └── child 6
	//             └── child 7
	//                 └── child 8
}
```
プログラム上で構成した木を、標準出力にツリーとして出力するというものです。
:::

また、ツリーを出力する他の Go ライブラリはいくつか[^1]ありますが、それらと比較しこのライブラリの特徴を挙げると以下です。
:::details 特徴
|gtree|他の Go ライブラリ|
|--|--|
|ユーザー側で木を組み立てるときは、**`Add` メソッドのみ**が使える|ノード（または枝）を装飾するための便利メソッドが用意されているものがある|
||tree コマンドで提供される様々なフラグと同様の機能が実装されているものがある|
||木を組み立てるときに最低２種のメソッド呼び出しが必要になるものがある（ユーザー側で、そのノードが子を持つか持たないか判定して、いずれかのメソッドを呼ぶイメージ）|
|Add メソッドはチェーンでき、木構造の**階層を降りながら**ノードを作る|チェーンでノードを作れるメソッドはあるが、階層を降りずに同一階層でノードを作るものがある|
|Add メソッドでノード追加時に、同一階層に同じ名前のノードが既にある場合、**既存のノードが返る**|同一階層に同じ名前のノードがあっても別ノードとして返すものがある|

こうして見ると、機能面では gtree は劣っているかもしれません。しかし、木を構成する（そしてユーザー側で構成した木を使って何らか処理させる）という点において、simple & easy かと思います。
とはいえ実装する際にはサンプルコードやドキュメント類を見ていただかないと難しいとも思っています🙇
:::


そして、この OSS ですが開発が始まったのが 2021/05 でした。なので、現在（2023/12）までで2年以上の付き合いになります。

:::details 歴史・宣伝媒体

作り始めたきっかけは、過去に職場メンバーの「treeを手動で作るのは手間なんだけど便利なツールないだろうか？」という呟きを耳にしたことです。最初は CLI として作っていて、当時の[メモ（リンク先の gtree セクション）](https://scrapbox.io/ddddddo/useful_tools)によると、2日で動くものはできていたようです。
しかし、色々欲が出て、ライブラリとしても提供できそうと考えたり、リファクタリングや機能追加、[試してみたいアーキテクチャ実装](https://github.com/ddddddO/gtree?tab=readme-ov-file#eg-gtreepipeline_treego)、README の推敲等をかなりしました（次は、Go1.22 で試験的に導入されている [`range-over-function`](https://go.dev/wiki/RangefuncExperiment) を使った実装を導入したいと[考えています](https://github.com/ddddddO/gtree/issues/274)）。

Go で Wasm 化できるなら Web 上でも楽に[サービス提供](https://ddddddo.github.io/gtree/)できそうなど、この OSS では色々遊ぶことができたと思っています。[Wasm について社内 LT で発表](https://github.com/ddddddO/web-editor/blob/main/docs/LT.pdf)するきっかけにもなりました（本記事から逸れるのですが、Wasm バイナリをブラウザでキャッシュしていれば、ユーザーから一番近いエッジ（手元の端末）で処理が完結するわけですし、これはエッジコンピューティングなんでしょうか？）。

また、機能追加の度に宣伝もしていました（Reddit, Slack#Gophersワークスペース, X, HackerNews ...etc）。あまり目立たない OSS は、ちょくちょく宣伝するとか、著名な人に取り上げてもらうなどないと多くの人に知ってもらうのは厳しいと実感しています。
他の人に知ってもらうという点では、上記の刹那的な宣伝だけでなく、永続的に残るようなところにも宣伝できるとなお良さそうです。例えば、以下です。
- Zenn
- [awesome-go](https://github.com/avelino/awesome-go)
    - awesome-xxx は色々あるので、OSS の特徴に合った awesome-xxx に宣伝するといいかもしれません。
    - ただ、それぞれのawesomeリポジトリで要件があったりします。

:::

こうして2年以上過ごしてきたわけですが...

# 2つのCLIツールが誕生していました🎉!!

最近（2023/09）まで上で紹介したライブラリは使われている気配がありませんでしたが、利用して頂ける方が現れました！以下が、その OSS です

|⭐[**_orangekame3/stree_**](https://github.com/orangekame3/stree)|⭐[**_owlinux1000/gcstree_**](https://github.com/owlinux1000/gcstree)|
|--|--|
|・Amazon S3 の bucket をツリー表示してくれる CLI で、読みは「エスツリー」<br>・[紹介記事](https://future-architect.github.io/articles/20230926a/)|・Google Cloud Storage（gcs） の bucket をツリー表示してくれる CLI|
|![](https://storage.googleapis.com/zenn-user-upload/71eab3ebb43a-20231209.png =300x)|![](https://storage.googleapis.com/zenn-user-upload/e5059127d8c9-20231209.png =310x)|


どちらの OSS も、クラウドの storage bucket をツリー表示してくれるもので、私にはその発想はなくハッとしました。
AWS / GCP が提供している公式のコマンド（`aws s3`/`gcloud storage`）には、現状ツリーを出力するという機能はありませんし、bucket 内が多量でもディレクトリ指定できるため、tree したいときに重宝するツールです👍（インストールも楽ですし、もっと広まればいいなと思います🙏）

gtree ライブラリは、ローカルPC上で他のコマンドと併用した何らかのユーティリティ作成に役立つかな...くらいに考えていましたが（[別記事](https://zenn.dev/ddddddo/articles/8cd85c68763f2e#%E3%81%A9%E3%82%93%E3%81%AA%E3%81%A8%E3%81%8D%E3%81%AB%E4%BD%BF%E3%81%88%E3%82%8B%E3%81%AE%EF%BC%9F)）、クラウドは広く利用されるようになっていますので、このような公益性のあるツールに組み込んでいただき、とても嬉しく思っています。

また、これらの OSS に対して、PullRequest を投げたのですが[^2]親切に対応していただきました！ありがとうございました！私がバグを埋め込んでしまった PullRequest があったのですが、修正を高速にレビュー・マージいただき感謝です🙇

# 結び
[gtree](https://github.com/ddddddO/gtree) の剪定を2年以上してきましたが、「あまり他の人に使ってもらえなくても、これはこれで楽しい」という気持ちでやっていました（また、「[自分を救うプログラミング](https://sizu.me/naoya/posts/7vxkuwvowo0z)」 にありますが、自分を癒すための側面もあったと思います）。

とは言え、開発している OSS が、他の人に利用してもらっているとわかるとやっぱり嬉しいものです。社内でも、良いと伝えられた時はとても嬉しかったです。

嬉しいですし、自分にはなかった発想を知れたり、PullRequest や X 上のライブ感のあるやり取りが楽しかったり、ライブラリをもう少し進化させられないか？と考える機会も出てきたりと良い体験になりました。

:::details （ライブラリの進化についてどんな風に考えたか）
これら CLI を実際に使ってみて、
- ツリーの出力から各行を選択してオブジェクトのダウンロードができるともっと嬉しいかも？
- なら、gtree で木構造の各ノードのパスを返せばクラウドの SDK なり WebAPI なりでダウンロードしやすくなるのではないか？
- [xlab/treeprint](https://github.com/xlab/treeprint#iterating-over-the-tree-nodes) は、ユーザー定義関数を使える仕組みを提供してる、これ良さそう
- せっかくだからパス以外にも、ノードが返した方が良さそうなデータは利用できるようにしよう
	- 誤った使い方は防ぎたいから、ノードの参照ではなく値で返そう

と考えました（[検討issue](https://github.com/ddddddO/gtree/issues/252)）。

その結果、[新しく関数（WalkProgrammably func）を追加](https://github.com/ddddddO/gtree#walkprogrammably-func)しました。柔軟に利用できるのではと思っています。
:::

改めて、 [**orangekame3/stree**](https://github.com/orangekame3/stree) と [**owlinux1000/gcstree**](https://github.com/owlinux1000/gcstree) の開発者に感謝致します（楽しかったです！）。
そして、きっかけを与えてくれた過去の同僚にも感謝です。この OSS を通じて色々と縁が生まれました。

最後まで記事を読んでいただき、ありがとうございました！

# ほんとに最後に
よければぜひ！SRE / CRE も絶賛募集中です！（2024/11/18 現在）
https://herp.careers/v1/openlogi/requisition-groups/486b8b01-6cf9-4434-8601-381c9c092e0d

[OPENLOGI Advent Calendar 2023](https://qiita.com/advent-calendar/2023/openlogi) 3日目は、[kt-tsutsumi](https://qiita.com/kt-tsutsumi) さんの「[**CREチームへのお問い合わせをscikit-learnでラベリングしてみた**](https://qiita.com/kt-tsutsumi/items/0b1276062093a334a599)」です！

[^1]: 比較の参考にしたライブラリは、 [xlab/treeprint](https://github.com/xlab/treeprint), [a8m/tree](https://github.com/a8m/tree), [bayashi/go-proptree](https://github.com/bayashi/go-proptree) です。利用者が多く安心なライブラリや、木構造を辿ってユーザー定義関数を実行出来たり、ノード説明用のメソッドがあったりと、大変刺激になりました。ありがとうございます。

[^2]: https://github.com/orangekame3/stree/pull/18 や、https://github.com/owlinux1000/gcstree/pull/1 など。https://github.com/owlinux1000/gcstree/pull/6 はバグ修正をサッと取り込んでいただけた PullRequest