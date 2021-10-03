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
		showVersion           bool
		mdFilepath            string
		twoSpaces, fourSpaces bool
		watching              bool
	)
	flag.BoolVar(&showVersion, "v", false, "current gtree version")
	flag.StringVar(&mdFilepath, "f", "", "markdown file path")
	flag.BoolVar(&twoSpaces, "ts", false, "for indent two spaces")
	flag.BoolVar(&fourSpaces, "fs", false, "for indent four spaces")
	flag.BoolVar(&watching, "w", false, "watching input file")
	flag.Parse()

	if showVersion {
		fmt.Printf("gtree version %s / revision %s\n", Version, Revision)
		return
	}
	if twoSpaces && fourSpaces {
		fmt.Errorf("%s", `choose either "ts" or "fs".`)
		os.Exit(1)
	}

	if mdFilepath == "" || mdFilepath == "-" {
		if err := execute(os.Stdout, os.Stdin, twoSpaces, fourSpaces); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		return
	}

	filePath, err := filepath.Abs(mdFilepath)
	if err != nil {
		fmt.Errorf("%+v", err)
		os.Exit(1)
	}

	if !watching {
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Errorf("%+v", err)
			os.Exit(1)
		}
		defer file.Close()

		if err := execute(os.Stdout, file, twoSpaces, fourSpaces); err != nil {
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

				_ = execute(os.Stdout, file, twoSpaces, fourSpaces)
			}
		}()
	}
}

func execute(out io.Writer, in io.Reader, twoSpaces, fourSpaces bool) error {
	var err error
	switch {
	case twoSpaces:
		err = gtree.Execute(out, in, gtree.IndentTwoSpaces())
	case fourSpaces:
		err = gtree.Execute(out, in, gtree.IndentFourSpaces())
	default:
		err = gtree.Execute(out, in)
	}
	return err
}
