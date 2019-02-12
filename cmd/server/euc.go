package main

import (
	"log"
	"strconv"
	"strings"
)

// FIXME: 原因ここではなさそうだけど、数字がclient側に謎文字で出力される
// NOTE: 原因は、EUC文字コード変換が行われる。それで面白いのでそのままにする。
// EUC変換表: http://charset.7jp.net/euc.html
// 例: in: 33 → out: !
func convEUCHandler(msg string) string {
	i, err := strconv.Atoi(strings.TrimRight(msg, "\n"))
	if err != nil {
		log.Fatal(err)
	}

	return string(i)
}
