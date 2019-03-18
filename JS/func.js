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