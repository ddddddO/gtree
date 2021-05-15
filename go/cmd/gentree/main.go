package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// s := `aa`
	// r := strings.NewReader(s)
	r := os.Stdin
	fmt.Println(gen(r))
}

func gen(input io.Reader) string {
	var scanner *bufio.Scanner

	switch v := input.(type) {
	case *strings.Reader:
		fmt.Println("sr")
		scanner = bufio.NewScanner(input)
		_ = v
	case *os.File:
		fmt.Println("os")
		scanner = bufio.NewScanner(input)
	default:
		panic("unsupported type")
	}

	output := ""
	for scanner.Scan() {
		row := scanner.Text()
		// NOTE: test case 3から、次の行を見ないと計算できなくなる。
		//       そのため、一度入力を全て読み込む必要がある。
		//       読み込んだ入力を、例えばx軸・y軸な二次元配列で表して、各座標となるノード用の構造体も用意して、、な感じになると思う。以下な感じ。
		//       https://play.golang.org/p/Ey_T-Xw2MHi

		converted := convert(row)
		output += converted + "\n"
	}
	return strings.TrimSpace(output)
}

func convert(row string) string {
	fmt.Println("↓↓↓↓↓↓")

	converted := ""

	tabCnt := 0
	for _, r := range row {
		// https://ja.wikipedia.org/wiki/ASCII
		switch r {
		case 45: // -
			continue
		case 32: // space
			continue
		case 9: // tab
			tabCnt++
		default: // directry or file name char
			converted += convertTab(tabCnt) + " " + string(r)
			tabCnt = 0
		}

		fmt.Println(r, string(r))
	}

	fmt.Println("↑↑↑↑↑↑")

	return converted
}

const convertedTab = "└" + "─" + "─"

func convertTab(cnt int) string {
	converted := ""
	if cnt == 0 {
		return converted
	}

	for i := 0; i < cnt-1; i++ {
		converted += "    "
	}
	converted += convertedTab
	return converted
}
