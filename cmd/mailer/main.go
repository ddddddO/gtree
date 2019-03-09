package main

import (
	"flag"
	"fmt"
	"os"
)

var target string

func main() {
	flag.StringVar(&target, "target", "", "destination mail address")
	flag.Parse()

	if len(target) == 0 {
		tFlag := flag.Lookup("target")
		fmt.Printf("%s : %s\n", tFlag.Name, tFlag.Usage)
		os.Exit(1)
	}

	m := NewMail()
	if err := m.send(target); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
