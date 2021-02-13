// 関数宣言よりも前に関数呼び出しを行えることを「巻き上げ」というらしい。
fooo(1, 3);

function fooo(x, y) {
    console.log(x + y);
}

// 関数式
const barr = function (x, y) {
    console.log(x - y);
}

barr(100, 20);

// 等値比較(ふつうは===を使う。が、undefined(またはnull)チェックは ==null でおこなう)
let num = 9
let chr = "9"

if (num == chr) {
    console.log("型変換が実行される")
}

if (!(num === chr)) {
    console.log("型変換が実行されない")
}

let s = null
if (s == null) {
    console.log("null desu")
}

if (s === null) {
    console.log("null と判定できる？")
} else {
    console.log("null と判定できませんでした")
}

let p = undefined
if (p == null) {
    console.log("null?")
} else {
    console.log("nulnul?")
}

// こちらだと判定できない
if (p === null) {
    console.log("nulnulnul?")
} else {
    console.log("nulnulnulnul?")
}