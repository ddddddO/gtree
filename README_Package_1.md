## Package(1) / like CLI

### Installation
```console
go get github.com/ddddddO/gtree
```

### Usage

```go
package main

import (
	"bytes"
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
		panic(err)
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
	// When indentation is four spaces, use IndentFourSpaces func instead of IndentTwoSpaces func.
	// and, you can customize branch format.
	if err := gtree.Output(os.Stdout, r2,
		gtree.IndentTwoSpaces(),
		gtree.BranchFormatIntermedialNode("+->", ":   "),
		gtree.BranchFormatLastNode("+->", "    "),
	); err != nil {
		panic(err)
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
  - `gtree.EncodeJSON()`
  - `gtree.EncodeTOML()`
  - `gtree.EncodeYAML()`
