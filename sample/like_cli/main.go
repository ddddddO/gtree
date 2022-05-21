package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/ddddddO/gtree/sample/like_cli/adapter"
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
}
