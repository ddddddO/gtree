## 参考
https://note.mu/erukiti/n/n02217beab2ef
SD３月号　P62~87 (P76まで)

## exec command
node 'JSfile'

## google map
https://developers.google.com/maps/documentation/javascript/tutorial

---
# 基礎
#### クラスとファイル分割
- クラス定義後に`export default クラス名`で、クラスをエクスポートして他のファイルから使用する
- エクスポートされたクラスは、インポートするファイルの先頭で`import クラス名 from "./ファイル名"` .jsは省略可
- クラス以外(文字列、数値、関数とか)もエクスポート可能
- `export default`は、ファイルに１つのみ。そのファイルがインポートされた場合、自動的に`export default 値`の値がインポートされる。
- `名前付きエクスポート`： 複数の値をエクスポート可能。インポート時は複数指定する。


#### 配列操作メソッド
- push
- forEach
```JavaScript
const chars = ["A", "B", "C"];

chars.forEach((char)=> {
  console.log(char);
});
```
- find
コールバック関数の処理部分に記述した条件式に合う`1つ目の`要素を配列の中から取り出すメソッド。配列の要素がオブジェクトの場合も使用可能。
```javascript
const numbers = [1, 2, 3, 4, 5, 6];

const foundNumber = numbers.find((number)=> {
  return number%3 === 0;
});
```
- filter
記述した条件に合う要素のみを取り出して新しい配列を作成するメソッド
findと使い方は同じ。
- map
配列の全要素に対して同一の処理を実行し、その結果の配列を返す。
findと使い方は同じ。
