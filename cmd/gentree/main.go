package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/ddddddO/gentree"
)

// These variables are set in build step
var (
	Version  = "unset"
	Revision = "unset"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("failed to gentree...\nplease review the file format.\nhint: %s\n", err)
			os.Exit(1)
		}
	}()

	var (
		f                         string
		isTwoSpaces, isFourSpaces bool
	)
	flag.StringVar(&f, "f", "", "markdown file path")
	flag.BoolVar(&isTwoSpaces, "ts", false, "for indent two spaces")
	flag.BoolVar(&isFourSpaces, "fs", false, "for indent four spaces")
	flag.Parse()

	if isTwoSpaces && isFourSpaces {
		fmt.Errorf("%s", `choose either "ts" or "fs".`)
		os.Exit(1)
	}

	var input io.Reader
	if f == "" || f == "-" {
		input = os.Stdin
	} else {
		filePath, err := filepath.Abs(f)
		if err != nil {
			fmt.Errorf("%+v", err)
			os.Exit(1)
		}
		input, err = os.Open(filePath)
		if err != nil {
			fmt.Errorf("%+v", err)
			os.Exit(1)
		}
		defer input.(*os.File).Close()
	}

	fmt.Println(gentree.Execute(input, isTwoSpaces, isFourSpaces))
}
