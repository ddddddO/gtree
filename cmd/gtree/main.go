package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/ddddddO/gtree"
)

// These variables are set in build step
var (
	Version  = "unset"
	Revision = "unset"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("failed to gtree...\nplease review the file format.\nhint: %s\n", err)
			os.Exit(1)
		}
	}()

	var (
		isVersion                 bool
		f                         string
		isTwoSpaces, isFourSpaces bool
	)
	flag.BoolVar(&isVersion, "v", false, "current gtree version")
	flag.StringVar(&f, "f", "", "markdown file path")
	flag.BoolVar(&isTwoSpaces, "ts", false, "for indent two spaces")
	flag.BoolVar(&isFourSpaces, "fs", false, "for indent four spaces")
	flag.Parse()

	if isVersion {
		fmt.Printf("gtree version %s / revision %s\n", Version, Revision)
		return
	}

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

	conf := gtree.Config{
		IsTwoSpaces:  isTwoSpaces,
		IsFourSpaces: isFourSpaces,
	}

	output, err := gtree.Execute(input, conf)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(output)
}
