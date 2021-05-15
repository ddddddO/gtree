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
		output += scanner.Text() + "\n"
	}
	return output
}
