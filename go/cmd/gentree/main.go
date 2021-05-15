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

	// https://ja.wikipedia.org/wiki/ASCII
	for _, r := range row {
		switch r {
		case 45: // -
			continue
		case 32: // space
			continue
		case 9: // tab
			converted += "└" + "─" + "─"
		default: // directry or file name
			converted += " " + string(r)
		}

		fmt.Println(r, string(r))
	}

	fmt.Println("↑↑↑↑↑↑")

	return converted
}
