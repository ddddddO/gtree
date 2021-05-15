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
