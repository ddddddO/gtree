package main

// https://developers.google.com/sheets/api/quickstart/go

import (
	"flag"
	"fmt"
	"os"

	"github.com/ddddddO/work/gointernal/read"
)

func main() {
	var (
		spreadsheetURL      string
		spreadsheetPageName string
	)

	// errorHandling ErrorHandling ..?
	readFlgSet := flag.NewFlagSet("read", flag.ExitOnError)

	readFlgSet.StringVar(&spreadsheetURL, "url", "", "please input spreadsheet url")
	readFlgSet.StringVar(&spreadsheetPageName, "page", "", "please input spreadsheet page name")
	readFlgSet.Usage = func() {
		fmt.Println("read :for reading to google spread sheet")
		fmt.Printf("  --%s: %s\n", readFlgSet.Lookup("url").Name, readFlgSet.Lookup("url").Usage)
		fmt.Printf("  --%s: %s\n", readFlgSet.Lookup("page").Name, readFlgSet.Lookup("page").Usage)
	}

	if len(os.Args) == 1 {
		readFlgSet.Usage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "read":
		readFlgSet.Parse(os.Args[2:])
		if err := read.Run(spreadsheetURL, spreadsheetPageName); err != nil {
			fmt.Println(err)
			fmt.Println()
			readFlgSet.Usage()
			os.Exit(1)
		}

	default:
		readFlgSet.Usage()
		os.Exit(1)
	}

}
