package main

import (
	"fmt"
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

	twoSpaces := &adapter.TwoSpaces{
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

	fourSpaces := &adapter.FourSpaces{
		Data: dataFourSpaces,
	}

	outputer := []adapter.Outputer{
		tab,
		twoSpaces,
		fourSpaces,
	}

	for _, or := range outputer {
		if err := or.Output(); err != nil {
			panic(err)
		}
	}
}
