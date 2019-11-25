console.log("HELLO WORLD");

console.log(3 / 2);
console.log(2 ** 53 == 2**53 + 1);
console.log();

// 文字列はプリミティブ(=イミュータブル)
var str = "foobar";
var repStr = str.replace("foo", "hoge");
console.log(str);
console.log(repStr);
console.log();

// オブジェクト(基本、連想配列。あと配列)
var obj = {foo: 999, bar: "ああああ"};
console.log(obj);
console.log(obj.foo);
console.log(obj["bar"]);
console.log(obj.nnn);
console.log()

obj["buzz"] = "BUZZ";
console.log(obj);
console.log()

obj.buzz = "fizz";
console.log(obj);
console.log()

// オブジェクトの要素追加
obj.fizz = null; 
console.log(obj)
console.log()

var list = ["AA", 3, 000, true, {foo: "FOOOO"}];
console.log(list);
console.log(list[4].foo);
console.log(list[4]["foo"]);
console.log()

// オブジェクト参照型挙動
var ref1 = {foo: 123};
var ref2 = ref1;
ref2.foo = 456;
console.log(ref1.foo);
console.log();

// プロトタイプ
var arr = [];
console.log(arr instanceof Array);
var s = "";
console.log(s instanceof String);
console.log();

// for-in
var obj1 = {foo: 123, bar: "gggg", buzz: 999};
for (var key in obj1) {
    console.log(key)
}
console.log()
