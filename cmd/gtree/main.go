package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/ddddddO/gtree/v6"
)

// These variables are set in build step
var (
	Version  = "unset"
	Revision = "unset"
)

func main() {
	var (
		isVersion                 bool
		f                         string
		isTwoSpaces, isFourSpaces bool
		isWatching                bool
	)
	flag.BoolVar(&isVersion, "v", false, "current gtree version")
	flag.StringVar(&f, "f", "", "markdown file path")
	flag.BoolVar(&isTwoSpaces, "ts", false, "for indent two spaces")
	flag.BoolVar(&isFourSpaces, "fs", false, "for indent four spaces")
	flag.BoolVar(&isWatching, "w", false, "watching input file")
	flag.Parse()

	if isVersion {
		fmt.Printf("gtree version %s / revision %s\n", Version, Revision)
		return
	}
	if isTwoSpaces && isFourSpaces {
		fmt.Errorf("%s", `choose either "ts" or "fs".`)
		os.Exit(1)
	}

	if f == "" || f == "-" {
		if err := execute(os.Stdout, os.Stdin, isTwoSpaces, isFourSpaces); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		return
	}

	filePath, err := filepath.Abs(f)
	if err != nil {
		fmt.Errorf("%+v", err)
		os.Exit(1)
	}

	if !isWatching {
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Errorf("%+v", err)
			os.Exit(1)
		}
		defer file.Close()

		if err := execute(os.Stdout, file, isTwoSpaces, isFourSpaces); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		return
	}

	ticker := time.NewTicker(1 * time.Second)
	var preFileModTime time.Time
	for range ticker.C {
		func() {
			file, err := os.Open(filePath)
			if err != nil {
				fmt.Errorf("%+v", err)
				os.Exit(1)
			}
			defer file.Close()

			fileInfo, err := file.Stat()
			if err != nil {
				fmt.Errorf("%+v", err)
				os.Exit(1)
			}

			if fileInfo.ModTime() != preFileModTime {
				preFileModTime = fileInfo.ModTime()

				_ = execute(os.Stdout, file, isTwoSpaces, isFourSpaces)
			}
		}()
	}
}

func execute(out io.Writer, in io.Reader, isTwoSpaces, isFourSpaces bool) error {
	var err error
	switch {
	case isTwoSpaces:
		err = gtree.Execute(out, in, gtree.IndentTwoSpaces())
	case isFourSpaces:
		err = gtree.Execute(out, in, gtree.IndentFourSpaces())
	default:
		err = gtree.Execute(out, in)
	}
	return err
}
