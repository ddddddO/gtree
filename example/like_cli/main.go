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
		fmt.Printf("\tName     : %s\n", wn.Name())
		fmt.Printf("\tBranch   : %s\n", wn.Branch())
		fmt.Printf("\tRow      : %s\n", wn.Row())
		fmt.Printf("\tLevel    : %d\n", wn.Level())
		fmt.Printf("\tPath     : %s\n", wn.Path())
		fmt.Printf("\tHasChild : %t\n", wn.HasChild())
		return nil
	}

	if err := gtree.Walk(strings.NewReader(src), callback2); err != nil {
		return err
	}
	// Output:
	// WalkerNode's methods called...
	// 	Name     : a
	// 	Branch   :
	// 	Row      : a
	// 	Level    : 1
	// 	Path     : a
	// 	HasChild : true
	// WalkerNode's methods called...
	// 	Name     : i
	// 	Branch   : ├──
	// 	Row      : ├── i
	// 	Level    : 2
	// 	Path     : a/i
	// 	HasChild : true
	// WalkerNode's methods called...
	// 	Name     : u
	// 	Branch   : │   └──
	// 	Row      : │   └── u
	// 	Level    : 3
	// 	Path     : a/i/u
	// 	HasChild : true
	// WalkerNode's methods called...
	// 	Name     : k
	// 	Branch   : │       └──
	// 	Row      : │       └── k
	// 	Level    : 4
	// 	Path     : a/i/u/k
	// 	HasChild : false
	// WalkerNode's methods called...
	// 	Name     : kk
	// 	Branch   : └──
	// 	Row      : └── kk
	// 	Level    : 2
	// 	Path     : a/kk
	// 	HasChild : true
	// WalkerNode's methods called...
	// 	Name     : t
	// 	Branch   :     └──
	// 	Row      :     └── t
	// 	Level    : 3
	// 	Path     : a/kk/t
	// 	HasChild : false
	// WalkerNode's methods called...
	// 	Name     : e
	// 	Branch   :
	// 	Row      : e
	// 	Level    : 1
	// 	Path     : e
	// 	HasChild : true
	// WalkerNode's methods called...
	// 	Name     : o
	// 	Branch   : └──
	// 	Row      : └── o
	// 	Level    : 2
	// 	Path     : e/o
	// 	HasChild : true
	// WalkerNode's methods called...
	// 	Name     : g
	// 	Branch   :     └──
	// 	Row      :     └── g
	// 	Level    : 3
	// 	Path     : e/o/g
	// 	HasChild : false

	fmt.Println("\nWalker Sample...end")
	fmt.Println()

	return nil
}
