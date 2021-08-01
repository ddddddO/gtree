package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/ddddddO/gtree/v5"
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

	conf := gtree.Config{
		IsTwoSpaces:  isTwoSpaces,
		IsFourSpaces: isFourSpaces,
	}

	if f == "" || f == "-" {
		if err := gtree.Execute(os.Stdout, os.Stdin, conf); err != nil {
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

		if err := gtree.Execute(os.Stdout, file, conf); err != nil {
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
				gtree.Execute(os.Stdout, file, conf)
			}
		}()
	}
}
