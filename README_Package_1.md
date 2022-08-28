# Package(1) / like CLI

## Installation

Go version requires 1.18 or later.

```console
$ go get github.com/ddddddO/gtree
```

## Usage
### *Output* func

```go
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/ddddddO/gtree"
)

func main() {
	// input Markdown is tab indented
	r1 := bytes.NewBufferString(strings.TrimSpace(`
- root
	- dddd
		- kkkkkkk
			- lllll
				- ffff
				- LLL
					- WWWWW
						- ZZZZZ
				- ppppp
					- KKK
						- 1111111
							- AAAAAAA
	- eee`))
	if err := gtree.Output(os.Stdout, r1); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// Output:
	// root
	// ├── dddd
	// │   └── kkkkkkk
	// │       └── lllll
	// │           ├── ffff
	// │           ├── LLL
	// │           │   └── WWWWW
	// │           │       └── ZZZZZ
	// │           └── ppppp
	// │               └── KKK
	// │                   └── 1111111
	// │                       └── AAAAAAA
	// └── eee

	// input Markdown is two spaces indented
	r2 := bytes.NewBufferString(strings.TrimSpace(`
- a
  - i
    - u
      - k
      - kk
    - t
  - e
    - o
  - g`))
	// When indentation is four spaces, use WithIndentFourSpaces func instead of WithIndentTwoSpaces func.
	// and, you can customize branch format.
	if err := gtree.Output(os.Stdout, r2,
		gtree.WithIndentTwoSpaces(),
		gtree.WithBranchFormatIntermedialNode("+->", ":   "),
		gtree.WithBranchFormatLastNode("+->", "    "),
	); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// Output:
	// a
	// +-> i
	// :   +-> u
	// :   :   +-> k
	// :   :   +-> kk
	// :   +-> t
	// +-> e
	// :   +-> o
	// +-> g
}

```

- You can also output JSON/YAML/TOML.
  - `gtree.WithEncodeJSON()`
  - `gtree.WithEncodeTOML()`
  - `gtree.WithEncodeYAML()`

---

### *Mkdir* func

- `gtree.Mkdir` func makes directories.
	- You can use `gtree.WithFileExtensions` func to make specified extensions as file.
