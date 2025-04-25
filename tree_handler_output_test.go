package gtree_test

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/ddddddO/gtree"
	tu "github.com/ddddddO/gtree/testutil"
	"go.uber.org/goleak"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}

func TestOutput_detecting_goroutineleak(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(500*time.Millisecond))
	defer cancel()
	w := io.Discard
	r := strings.NewReader(tu.TwentyThousandRoots)
	if gotErr := gtree.OutputFromMarkdown(w, r, gtree.WithMassive(ctx)); gotErr != nil {
		if gotErr != context.DeadlineExceeded {
			t.Errorf("\ngotErr: \n%v\nwantErr: \n%v", gotErr, context.DeadlineExceeded)
		}
	}
}

func TestOutput_json_detecting_goroutineleak(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(500*time.Millisecond))
	defer cancel()
	w := io.Discard
	r := strings.NewReader(tu.TwentyThousandRoots)
	if gotErr := gtree.OutputFromMarkdown(w, r, gtree.WithEncodeJSON(), gtree.WithMassive(ctx)); gotErr != nil {
		if gotErr != context.DeadlineExceeded {
			t.Errorf("\ngotErr: \n%v\nwantErr: \n%v", gotErr, context.DeadlineExceeded)
		}
	}
}

func TestOutput_yaml_detecting_goroutineleak(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(500*time.Millisecond))
	defer cancel()
	w := io.Discard
	r := strings.NewReader(tu.TwentyThousandRoots)
	if gotErr := gtree.OutputFromMarkdown(w, r, gtree.WithEncodeYAML(), gtree.WithMassive(ctx)); gotErr != nil {
		if gotErr != context.DeadlineExceeded {
			t.Errorf("\ngotErr: \n%v\nwantErr: \n%v", gotErr, context.DeadlineExceeded)
		}
	}
}

func TestOutput_toml_detecting_goroutineleak(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(500*time.Millisecond))
	defer cancel()
	w := io.Discard
	r := strings.NewReader(tu.TwentyThousandRoots)
	if gotErr := gtree.OutputFromMarkdown(w, r, gtree.WithEncodeTOML(), gtree.WithMassive(ctx)); gotErr != nil {
		if gotErr != context.DeadlineExceeded {
			t.Errorf("\ngotErr: \n%v\nwantErr: \n%v", gotErr, context.DeadlineExceeded)
		}
	}
}

func TestOutput_dryrun_detecting_goroutineleak(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(500*time.Millisecond))
	defer cancel()
	w := io.Discard
	r := strings.NewReader(tu.TwentyThousandRoots)
	if gotErr := gtree.OutputFromMarkdown(w, r, gtree.WithDryRun(), gtree.WithMassive(ctx)); gotErr != nil {
		if gotErr != context.DeadlineExceeded {
			t.Errorf("\ngotErr: \n%v\nwantErr: \n%v", gotErr, context.DeadlineExceeded)
		}
	}
}

type in struct {
	input   io.Reader
	options []gtree.Option
}

type out struct {
	output string
	err    error
}

func TestOutputFromMarkdown(t *testing.T) {
	tests := []struct {
		name string
		in   in
		out  out
	}{
		{
			name: "case(succeeded/has a child)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- a
	- b`)),
			},
			out: out{
				output: strings.TrimPrefix(`
a
└── b
`, "\n"),
				err: nil,
			},
		},
		{
			name: "case(succeeded/has a child nest)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- a
	- b
		- c`))},

			out: out{
				output: strings.TrimPrefix(`
a
└── b
    └── c
`, "\n"),
				err: nil,
			},
		},
		{
			name: "case(succeeded/has children)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- a
	- b
	- c`))},
			out: out{
				output: strings.TrimPrefix(`
a
├── b
└── c
`, "\n"),
				err: nil,
			},
		},
		{
			name: "case(succeeded/has children deeply)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- a
	- b
		- c
			- d
			- e
			- f`))},
			out: out{
				output: strings.TrimPrefix(`
a
└── b
    └── c
        ├── d
        ├── e
        └── f
`, "\n"),
				err: nil,
			},
		},
		{
			name: "case(succeeded/has children complexly)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- a
	- i
		- u
			- k
			- kk
		- t
	- e
		- o
	- g`))},
			out: out{
				output: strings.TrimPrefix(`
a
├── i
│   ├── u
│   │   ├── k
│   │   └── kk
│   └── t
├── e
│   └── o
└── g
`, "\n"),
				err: nil,
			},
		},

		{
			name: "case(succeeded/same name on the same hierarchy)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- a
	- same_a
		- same_b
			- k
			- kk
		- t
	- p
		- q
	- same_a
		- o
		- same_b
			- ppp
	- g`))},
			out: out{
				output: strings.TrimPrefix(`
a
├── same_a
│   ├── same_b
│   │   ├── k
│   │   ├── kk
│   │   └── ppp
│   ├── t
│   └── o
├── p
│   └── q
└── g
`, "\n"),
				err: nil,
			},
		},

		{
			name: "case(succeeded/very deeply)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
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
	- eee`))},
			out: out{
				output: strings.TrimPrefix(`
root
├── dddd
│   └── kkkkkkk
│       └── lllll
│           ├── ffff
│           ├── LLL
│           │   └── WWWWW
│           │       └── ZZZZZ
│           └── ppppp
│               └── KKK
│                   └── 1111111
│                       └── AAAAAAA
└── eee
`, "\n"),
				err: nil,
			},
		},
		{
			name: "case(succeeded/indent 2spaces)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- a
  - i
    - u
      - k
      - kk
    - t
  - e
    - o
  - g`)),
			},
			out: out{
				output: strings.TrimPrefix(`
a
├── i
│   ├── u
│   │   ├── k
│   │   └── kk
│   └── t
├── e
│   └── o
└── g
`, "\n"),
				err: nil,
			},
		},
		{
			name: "case(succeeded/indent 4spaces)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- a
    - i
        - u
            - k
            - kk
        - t
    - e
        - o
    - g`)),
			},
			out: out{
				output: strings.TrimPrefix(`
a
├── i
│   ├── u
│   │   ├── k
│   │   └── kk
│   └── t
├── e
│   └── o
└── g
`, "\n"),
				err: nil,
			},
		},
		{
			name: "case(succeeded/node name 1space & -)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- root dir aaa
	- child-dir`))},
			out: out{
				output: strings.TrimPrefix(`
root dir aaa
└── child-dir
`, "\n"),
				err: nil,
			},
		},
		{
			name: "case(succeeded/same node name)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- parent
	- child
		- chilchil
		- chilchil
		- chilchil
	- child`))},
			out: out{
				output: strings.TrimPrefix(`
parent
└── child
    └── chilchil
`, "\n"),
				err: nil,
			},
		},
		{
			name: "case(succeeded/input byte)",
			in: in{
				input: bytes.NewBufferString(strings.TrimSpace(`
- a
	- b`)),
			},
			out: out{
				output: strings.TrimPrefix(`
a
└── b
`, "\n"),
				err: nil,
			},
		},
		{
			name: "case(succeeded/multi root)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- a
	- i
		- u
			- k
			- kk
		- t
	- e
		- o
	- g
- a
	- i
		- u
			- k
			- kk
		- t
	- e
		- o
	- g`))},
			out: out{
				output: strings.TrimPrefix(`
a
├── i
│   ├── u
│   │   ├── k
│   │   └── kk
│   └── t
├── e
│   └── o
└── g
a
├── i
│   ├── u
│   │   ├── k
│   │   └── kk
│   └── t
├── e
│   └── o
└── g
`, "\n"),
				err: nil,
			},
		},
		{
			name: "case(succeeded/multi root/massive)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- a
	- i
		- u
			- k
			- kk
		- t
	- e
		- o
	- g
- a
	- i
		- u
			- k
			- kk
		- t
	- e
		- o
	- g`)),
				options: []gtree.Option{gtree.WithMassive(context.Background())},
			},
			out: out{
				output: strings.TrimPrefix(`
a
├── i
│   ├── u
│   │   ├── k
│   │   └── kk
│   └── t
├── e
│   └── o
└── g
a
├── i
│   ├── u
│   │   ├── k
│   │   └── kk
│   └── t
├── e
│   └── o
└── g
`, "\n"),
				err: nil,
			},
		},
		{
			name: "case(empty node name)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- a
	-`)),
			},
			out: out{
				output: "",
				err:    gtree.ExportErrEmptyText,
			},
		},
		/*{
					// TODO: fixme
					name: "case(incorrect input format(input 4spaces indent / tab mode))",
					in: in{
						input: strings.NewReader(strings.TrimSpace(`
		- a
		    - b`)),
					},
					out: out{
						output: "",
						err:    gtree.ExportErrIncorrectFormat,
					},
				},*/
		{
			name: "case(bufio.Scanner err)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(fmt.Sprintf(`
- a
	- %s`, strings.Repeat("a", 64*1024)))),
			},
			out: out{
				output: "",
				err:    bufio.ErrTooLong,
			},
		},
		{
			name: "case(massive/bufio.Scanner err)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(fmt.Sprintf(`
- a
	- %s`, strings.Repeat("a", 64*1024)))),
				options: []gtree.Option{gtree.WithMassive(context.Background())},
			},
			out: out{
				output: "",
				err:    bufio.ErrTooLong,
			},
		},
		{
			name: "case(succeeded/input markdown file)",
			in: in{
				input: tu.PrepareMarkdownFile(t)},
			out: out{
				output: strings.TrimPrefix(`
Artiodactyla
├── Artiofabula
│   ├── Cetruminantia
│   │   ├── Whippomorpha
│   │   │   ├── Hippopotamidae
│   │   │   └── Cetacea
│   │   └── Ruminantia
│   └── Suina
└── Tylopoda
Carnivora
├── Feliformia
└── Caniformia
    ├── Canidae
    └── Arctoidea
        ├── Ursidae
        └── x
            ├── Pinnipedia
            └── Musteloidea
                ├── Ailuridae
                └── x
                    ├── Mephitidae
                    └── x
                        ├── Procyonidae
                        └── Mustelidae
`, "\n"),
				err: nil,
			},
		},
		{
			name: "case(succeeded/indent 2spaces and cutom branch format)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- a
  - i
    - u
      - k
      - kk
    - t
  - e
    - o
  - g`)),
				options: []gtree.Option{
					gtree.WithBranchFormatIntermedialNode("+->", ":   "),
					gtree.WithBranchFormatLastNode("+->", "    "),
				},
			},
			out: out{
				output: strings.TrimPrefix(`
a
+-> i
:   +-> u
:   :   +-> k
:   :   +-> kk
:   +-> t
+-> e
:   +-> o
+-> g
`, "\n"),
				err: nil,
			},
		},
		{
			name: "case(succeeded/dry run/no error)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- a
	- b`)),
				options: []gtree.Option{
					gtree.WithDryRun(),
				},
			},
			out: out{
				output: strings.TrimPrefix(`
a
└── b

2 directories, 0 files
`, "\n"),
				err: nil,
			},
		},
		{
			name: "case(dry run/invalid node name)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- a
	- b/c`)),
				options: []gtree.Option{
					gtree.WithDryRun(),
				},
			},
			out: out{
				output: "",
				err:    errors.New("invalid node name: b/c"),
			},
		},
		{
			name: "case(input format error)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- a
	 - b`)),
				options: []gtree.Option{
					gtree.WithDryRun(),
				},
			},
			out: out{
				output: "",
				err:    gtree.ExportErrIncorrectFormat("	 - b"),
			},
		},
		{
			name: "case(succeeded/tab on the way)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- a	prev tab
	- b	prev tab`)),
			},
			out: out{
				output: strings.TrimPrefix(`
a	prev tab
└── b	prev tab
`, "\n"),
				err: nil,
			},
		},
		{
			// 複数Rootブロックを指定すべきだが、実装上、出力の順番が保証されないため1Rootで実施
			name: "case(succeeded/when massive root)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- a
	- b
		- c`)),
				options: []gtree.Option{
					gtree.WithMassive(context.Background()),
				},
			},
			out: out{
				output: strings.TrimPrefix(`
a
└── b
    └── c
`, "\n"),
				err: nil,
			},
		},
		{
			// 複数Rootブロックを指定すべきだが、実装上、出力の順番が保証されないため1Rootで実施
			name: "case(succeeded/when massive root and dryrun)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- a
	- b
		- z
		- c
	- y`)),
				options: []gtree.Option{
					gtree.WithMassive(context.Background()),
					gtree.WithDryRun(),
					gtree.WithFileExtensions([]string{"c"}),
				},
			},
			out: out{
				output: strings.TrimPrefix(`
a
├── b
│   ├── z
│   └── c
└── y

4 directories, 1 files
`, "\n"),
				err: nil,
			},
		},
		{
			// 複数Rootブロックを指定すべきだが、実装上、出力の順番が保証されないため1Rootで実施
			name: "case(succeeded/when massive root and json)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- a
	- b
		- c`)),
				options: []gtree.Option{
					gtree.WithMassive(context.Background()),
					gtree.WithEncodeJSON(),
				},
			},
			out: out{
				output: `{"value":"a","children":[{"value":"b","children":[{"value":"c","children":null}]}]}` + "\n",
				err:    nil,
			},
		},
		{
			// 複数Rootブロックを指定すべきだが、実装上、出力の順番が保証されないため1Rootで実施
			name: "case(succeeded/when massive root and yaml)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- a
	- b
		- c`)),
				options: []gtree.Option{
					gtree.WithMassive(context.Background()),
					gtree.WithEncodeYAML(),
				},
			},
			out: out{
				output: strings.TrimSpace(`
value: a
children:
- value: b
  children:
  - value: c
    children: []
`) + "\n",
				err: nil,
			},
		},
		{
			// 複数Rootブロックを指定すべきだが、実装上、出力の順番が保証されないため1Rootで実施
			name: "case(succeeded/when massive root and toml)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- a
	- b
		- c`)),
				options: []gtree.Option{
					gtree.WithMassive(context.Background()),
					gtree.WithEncodeTOML(),
				},
			},
			out: out{
				output: strings.TrimSpace(`
value = 'a'

[[children]]
value = 'b'

[[children.children]]
value = 'c'
children = []
`) + "\n",
				err: nil,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			out := &bytes.Buffer{}
			gotErr := gtree.OutputFromMarkdown(out, tt.in.input, tt.in.options...)
			gotOutput := out.String()

			if gotOutput != tt.out.output {
				t.Errorf("\ngot: \n%s\nwant: \n%s", gotOutput, tt.out.output)
			}
			if gotErr != nil {
				if gotErr.Error() != tt.out.err.Error() {
					t.Errorf("\ngotErr: \n%v\nwantErr: \n%v", gotErr, tt.out.err)
				}
			}

			if file, ok := tt.in.input.(*os.File); ok {
				file.Close()
			}
		})
	}
}

func TestOutput_encodeJSON(t *testing.T) {
	tests := []struct {
		name string
		in   in
		out  out
	}{
		{
			name: "case(tab spaces & multi root & output json)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- a
	- i
		- u
			- k
			- kk
		- t
	- e
		- o
	- g
- a
	- i
		- u
			- k
			- kk
		- t
	- e
		- o
	- g`)),
				options: []gtree.Option{gtree.WithEncodeJSON()},
			},
			out: out{
				output: strings.TrimPrefix(`
{"value":"a","children":[{"value":"i","children":[{"value":"u","children":[{"value":"k","children":null},{"value":"kk","children":null}]},{"value":"t","children":null}]},{"value":"e","children":[{"value":"o","children":null}]},{"value":"g","children":null}]}
{"value":"a","children":[{"value":"i","children":[{"value":"u","children":[{"value":"k","children":null},{"value":"kk","children":null}]},{"value":"t","children":null}]},{"value":"e","children":[{"value":"o","children":null}]},{"value":"g","children":null}]}
`, "\n"),
				err: nil,
			},
		},
		{
			name: "case(indent 2spaces & output json)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- a
  - i
    - u
      - k
      - kk
    - t
  - e
    - o
  - g`)),
				options: []gtree.Option{gtree.WithEncodeJSON()},
			},
			out: out{
				output: strings.TrimPrefix(`
{"value":"a","children":[{"value":"i","children":[{"value":"u","children":[{"value":"k","children":null},{"value":"kk","children":null}]},{"value":"t","children":null}]},{"value":"e","children":[{"value":"o","children":null}]},{"value":"g","children":null}]}
`, "\n"),
				err: nil,
			},
		},
		{
			name: "case(indent 4spaces & output json)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- a
    - i
        - u
            - k
            - kk
        - t
    - e
        - o
    - g`)),
				options: []gtree.Option{gtree.WithEncodeJSON()},
			},
			out: out{
				output: strings.TrimPrefix(`
{"value":"a","children":[{"value":"i","children":[{"value":"u","children":[{"value":"k","children":null},{"value":"kk","children":null}]},{"value":"t","children":null}]},{"value":"e","children":[{"value":"o","children":null}]},{"value":"g","children":null}]}
`, "\n"),
				err: nil,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			out := &bytes.Buffer{}
			gotErr := gtree.OutputFromMarkdown(out, tt.in.input, tt.in.options...)
			gotOutput := out.String()

			if gotOutput != tt.out.output {
				t.Errorf("\ngot: \n%s\nwant: \n%s", gotOutput, tt.out.output)
			}
			if gotErr != tt.out.err {
				t.Errorf("\ngotErr: \n%v\nwantErr: \n%v", gotErr, tt.out.err)
			}
		})
	}
}

func TestOutput_encodeTOML(t *testing.T) {
	tests := []struct {
		name string
		in   in
		out  out
	}{
		{
			name: "case(succeeded/tab spaces & multi root & output toml)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- a
	- i
		- u
			- k
			- kk
		- t
	- e
		- o
	- g
- a
	- i
		- u
			- k
			- kk
		- t
	- e
		- o
	- g`)),
				options: []gtree.Option{gtree.WithEncodeTOML()},
			},
			out: out{
				output: strings.TrimPrefix(`
value = 'a'

[[children]]
value = 'i'

[[children.children]]
value = 'u'

[[children.children.children]]
value = 'k'
children = []

[[children.children.children]]
value = 'kk'
children = []

[[children.children]]
value = 't'
children = []

[[children]]
value = 'e'

[[children.children]]
value = 'o'
children = []

[[children]]
value = 'g'
children = []
value = 'a'

[[children]]
value = 'i'

[[children.children]]
value = 'u'

[[children.children.children]]
value = 'k'
children = []

[[children.children.children]]
value = 'kk'
children = []

[[children.children]]
value = 't'
children = []

[[children]]
value = 'e'

[[children.children]]
value = 'o'
children = []

[[children]]
value = 'g'
children = []
`, "\n"),
				err: nil,
			},
		},
		{
			name: "case(succeeded/indent 2spaces & output toml)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- a
  - i
    - u
      - k
      - kk
    - t
  - e
    - o
  - g`)),
				options: []gtree.Option{gtree.WithEncodeTOML()},
			},
			out: out{
				output: strings.TrimPrefix(`
value = 'a'

[[children]]
value = 'i'

[[children.children]]
value = 'u'

[[children.children.children]]
value = 'k'
children = []

[[children.children.children]]
value = 'kk'
children = []

[[children.children]]
value = 't'
children = []

[[children]]
value = 'e'

[[children.children]]
value = 'o'
children = []

[[children]]
value = 'g'
children = []
`, "\n"),
				err: nil,
			},
		},
		{
			name: "case(succeeded/indent 4spaces & output toml)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- a
    - i
        - u
            - k
            - kk
        - t
    - e
        - o
    - g`)),
				options: []gtree.Option{gtree.WithEncodeTOML()},
			},
			out: out{
				output: strings.TrimPrefix(`
value = 'a'

[[children]]
value = 'i'

[[children.children]]
value = 'u'

[[children.children.children]]
value = 'k'
children = []

[[children.children.children]]
value = 'kk'
children = []

[[children.children]]
value = 't'
children = []

[[children]]
value = 'e'

[[children.children]]
value = 'o'
children = []

[[children]]
value = 'g'
children = []
`, "\n"),
				err: nil,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			out := &bytes.Buffer{}
			gotErr := gtree.OutputFromMarkdown(out, tt.in.input, tt.in.options...)
			gotOutput := out.String()

			if gotOutput != tt.out.output {
				t.Errorf("\ngot: \n%s\nwant: \n%s", gotOutput, tt.out.output)
			}
			if gotErr != tt.out.err {
				t.Errorf("\ngotErr: \n%v\nwantErr: \n%v", gotErr, tt.out.err)
			}
		})
	}
}

func TestOutput_encodeYAML(t *testing.T) {
	tests := []struct {
		name string
		in   in
		out  out
	}{
		{
			name: "case(succeeded/tab spaces & multi root & output yaml)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- a
	- i
		- u
			- k
			- kk
		- t
	- e
		- o
	- g
- a
	- i
		- u
			- k
			- kk
		- t
	- e
		- o
	- g`)),
				options: []gtree.Option{gtree.WithEncodeYAML()},
			},
			out: out{
				output: strings.TrimPrefix(`
value: a
children:
- value: i
  children:
  - value: u
    children:
    - value: k
      children: []
    - value: kk
      children: []
  - value: t
    children: []
- value: e
  children:
  - value: o
    children: []
- value: g
  children: []
---
value: a
children:
- value: i
  children:
  - value: u
    children:
    - value: k
      children: []
    - value: kk
      children: []
  - value: t
    children: []
- value: e
  children:
  - value: o
    children: []
- value: g
  children: []
`, "\n"),
				err: nil,
			},
		},
		{
			name: "case(succeeded/indent 2spaces & output yaml)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- a
  - i
    - u
      - k
      - kk
    - t
  - e
    - o
  - g`)),
				options: []gtree.Option{gtree.WithEncodeYAML()},
			},
			out: out{
				output: strings.TrimPrefix(`
value: a
children:
- value: i
  children:
  - value: u
    children:
    - value: k
      children: []
    - value: kk
      children: []
  - value: t
    children: []
- value: e
  children:
  - value: o
    children: []
- value: g
  children: []
`, "\n"),
				err: nil,
			},
		},
		{
			name: "case(succeeded/indent 4spaces & output yaml)",
			in: in{
				input: strings.NewReader(strings.TrimSpace(`
- a
    - i
        - u
            - k
            - kk
        - t
    - e
        - o
    - g`)),
				options: []gtree.Option{gtree.WithEncodeYAML()},
			},
			out: out{
				output: strings.TrimPrefix(`
value: a
children:
- value: i
  children:
  - value: u
    children:
    - value: k
      children: []
    - value: kk
      children: []
  - value: t
    children: []
- value: e
  children:
  - value: o
    children: []
- value: g
  children: []
`, "\n"),
				err: nil,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			out := &bytes.Buffer{}
			gotErr := gtree.OutputFromMarkdown(out, tt.in.input, tt.in.options...)
			gotOutput := out.String()

			if gotOutput != tt.out.output {
				t.Errorf("\ngot: \n%s\nwant: \n%s", gotOutput, tt.out.output)
			}
			if gotErr != tt.out.err {
				t.Errorf("\ngotErr: \n%v\nwantErr: \n%v", gotErr, tt.out.err)
			}
		})
	}
}

// TODO: config.go用にtest.goあってもいいんじゃないか
func TestOutput_nilctx(t *testing.T) {
	w := io.Discard
	r := strings.NewReader(tu.SingleRoot)
	if gotErr := gtree.OutputFromMarkdown(w, r, gtree.WithMassive(nil)); gotErr != nil {
		t.Errorf("\ngotErr: \n%v\nwantErr: \n%v", gotErr, nil)
	}
}

// TODO: config.go用にtest.goあってもいいんじゃないか
func TestOutput_nilopt(t *testing.T) {
	w := io.Discard
	r := strings.NewReader(tu.SingleRoot)
	var emptyOpt gtree.Option
	if gotErr := gtree.OutputFromMarkdown(w, r, emptyOpt); gotErr != nil {
		t.Errorf("\ngotErr: \n%v\nwantErr: \n%v", gotErr, nil)
	}
}
