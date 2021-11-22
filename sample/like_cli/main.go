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

	executors := []adapter.Executor{
		tab,
		twoSpaces,
		fourSpaces,
	}

	for _, ex := range executors {
		if err := ex.Execute(); err != nil {
			panic(err)
		}
	}
}
