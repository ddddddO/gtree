package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/ddddddO/gtree"
	"github.com/ddddddO/gtree/example/like_cli/adapter"
)

func main() {
	fmt.Printf("exapmle\n\n")

	dataTab := strings.NewReader(strings.TrimSpace(`
- a
	- i
		- u
			- k
			- kk
		- t
	- e
		- o
	- g`))

	tab := &adapter.Tab{
		Data: dataTab,
	}

	dataTwoSpaces := strings.NewReader(strings.TrimSpace(`
- a
  - i
    - u
      - k
      - kk
    - t
  - e
    - o
  - g`))

	spacesTwo := &adapter.TwoSpaces{
		Data: dataTwoSpaces,
	}

	dataFourSpaces := strings.NewReader(strings.TrimSpace(`
- a
    - i
        - u
            - k
            - kk
        - t
    - e
        - o
    - g`))

	spacesFour := &adapter.FourSpaces{
		Data: dataFourSpaces,
	}

	outputer := []adapter.Outputer{
		tab,
		spacesTwo,
		spacesFour,
	}

	for _, or := range outputer {
		if err := or.Output(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	if err := sampleWalker(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func sampleWalker() error {
	fmt.Println("\nWalker Sample...")
	fmt.Println()

	src := strings.TrimSpace(`
- a
  - i
    - u
      - k
  - kk
    - t
- e
  - o
    - g`)

	callback := func(wn *gtree.WalkerNode) error {
		fmt.Println(wn.Row())
		return nil
	}

	if err := gtree.Walk(strings.NewReader(src), callback); err != nil {
		return err
	}
	// Output:
	// a
	// ├── i
	// │   └── u
	// │       └── k
	// └── kk
	// 		└── t
	// e
	// └── o
	// 		└── g

	fmt.Println("----------------------------------------")

	callback2 := func(wn *gtree.WalkerNode) error {
		fmt.Println("WalkerNode's methods called...")
		fmt.Printf("\tName   : %s\n", wn.Name())
		fmt.Printf("\tBranch : %s\n", wn.Branch())
		fmt.Printf("\tRow    : %s\n", wn.Row())
		fmt.Printf("\tPath   : %s\n", wn.Path())
		return nil
	}

	if err := gtree.Walk(strings.NewReader(src), callback2); err != nil {
		return err
	}
	// Output:
	// WalkerNode's methods called...
	// 	Name   : a
	// 	Branch :
	// 	Row    : a
	// 	Path   : a
	// WalkerNode's methods called...
	// 	Name   : i
	// 	Branch : ├──
	// 	Row    : ├── i
	// 	Path   : a/i
	// WalkerNode's methods called...
	// 	Name   : u
	// 	Branch : │   └──
	// 	Row    : │   └── u
	// 	Path   : a/i/u
	// WalkerNode's methods called...
	// 	Name   : k
	// 	Branch : │       └──
	// 	Row    : │       └── k
	// 	Path   : a/i/u/k
	// WalkerNode's methods called...
	// 	Name   : kk
	// 	Branch : └──
	// 	Row    : └── kk
	// 	Path   : a/kk
	// WalkerNode's methods called...
	// 	Name   : t
	// 	Branch :     └──
	// 	Row    :     └── t
	// 	Path   : a/kk/t
	// WalkerNode's methods called...
	// 	Name   : e
	// 	Branch :
	// 	Row    : e
	// 	Path   : e
	// WalkerNode's methods called...
	// 	Name   : o
	// 	Branch : └──
	// 	Row    : └── o
	// 	Path   : e/o
	// WalkerNode's methods called...
	// 	Name   : g
	// 	Branch :     └──
	// 	Row    :     └── g
	// 	Path   : e/o/g

	fmt.Println("\nWalker Sample...end")
	fmt.Println()

	return nil
}
