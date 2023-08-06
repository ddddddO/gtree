package markdown

import (
	"testing"
)

// TODO: テストもうちょいちゃんとする

func TestParser_ParseTab(t *testing.T) {
	tests := map[string]struct {
		rows    []string
		wants   []*Markdown
		wantErr error
	}{
		"hyphen only": {
			[]string{
				"- aaabbb",
				"- cc",
				"	- bb",
			},
			[]*Markdown{
				{hierarchy: 1, text: "aaabbb"},
				{hierarchy: 1, text: "cc"},
				{hierarchy: 2, text: "bb"},
			},
			nil,
		},
		"sharp included": {
			[]string{
				"# aaabbb",
				"- cc",
				"- oo",
				"	- ff",
				"		- pp",
				"			* ooo",
				"- bb",
				"# kk",
				"+ ll",
				"	+ jj",
			},
			[]*Markdown{
				{hierarchy: 1, text: "aaabbb"},
				{hierarchy: 2, text: "cc"},
				{hierarchy: 2, text: "oo"},
				{hierarchy: 3, text: "ff"},
				{hierarchy: 4, text: "pp"},
				{hierarchy: 5, text: "ooo"},
				{hierarchy: 2, text: "bb"},
				{hierarchy: 1, text: "kk"},
				{hierarchy: 2, text: "ll"},
				{hierarchy: 3, text: "jj"},
			},
			nil,
		},
		"sharp included 2": {
			[]string{
				"- oo",
				"	- ff",
				"		- pp",
				"# aaabb",
				"- cc",
				"- oo",
				"	- ff",
				"		- pp",
				"- bb",
				"### kk",
				"- ll",
			},
			[]*Markdown{
				{hierarchy: 1, text: "oo"},
				{hierarchy: 2, text: "ff"},
				{hierarchy: 3, text: "pp"},
				{hierarchy: 1, text: "aaabb"},
				{hierarchy: 2, text: "cc"},
				{hierarchy: 2, text: "oo"},
				{hierarchy: 3, text: "ff"},
				{hierarchy: 4, text: "pp"},
				{hierarchy: 2, text: "bb"},
				{hierarchy: 1, text: "kk"},
				{hierarchy: 2, text: "ll"},
			},
			nil,
		},
	}

	for name, tt := range tests {
		tt := tt
		parser := NewParser()

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			for i := range tt.rows {
				ret, err := parser.Parse(tt.rows[i])
				if err != tt.wantErr {
					t.Errorf("\ngot: \n%v\nwant: \n%v", err, tt.wantErr)
				}
				if *ret != *tt.wants[i] {
					t.Errorf("\ngot: \n%v\nwant: \n%v", ret, tt.wants[i])
				}
			}
		})
	}
}

func TestParser_ParseTwoSpaces(t *testing.T) {
	tests := map[string]struct {
		rows    []string
		wants   []*Markdown
		wantErr error
	}{
		"hyphen only": {
			[]string{
				"- aaabbb",
				"- cc",
				"  - bb",
			},
			[]*Markdown{
				{hierarchy: 1, text: "aaabbb"},
				{hierarchy: 1, text: "cc"},
				{hierarchy: 2, text: "bb"},
			},
			nil,
		},
		"sharp included": {
			[]string{
				"# aaabbb",
				"- cc",
				"- oo",
				"  - ff",
				"    - pp",
				"- bb",
				"# kk",
				"- ll",
			},
			[]*Markdown{
				{hierarchy: 1, text: "aaabbb"},
				{hierarchy: 2, text: "cc"},
				{hierarchy: 2, text: "oo"},
				{hierarchy: 3, text: "ff"},
				{hierarchy: 4, text: "pp"},
				{hierarchy: 2, text: "bb"},
				{hierarchy: 1, text: "kk"},
				{hierarchy: 2, text: "ll"},
			},
			nil,
		},
		"sharp included 2": {
			[]string{
				"- oo",
				"  - ff",
				"    - pp",
				"# aaabb",
				"- cc",
				"- oo",
				"  - ff",
				"    - pp",
				"- bb",
				"### kk",
				"- ll",
			},
			[]*Markdown{
				{hierarchy: 1, text: "oo"},
				{hierarchy: 2, text: "ff"},
				{hierarchy: 3, text: "pp"},
				{hierarchy: 1, text: "aaabb"},
				{hierarchy: 2, text: "cc"},
				{hierarchy: 2, text: "oo"},
				{hierarchy: 3, text: "ff"},
				{hierarchy: 4, text: "pp"},
				{hierarchy: 2, text: "bb"},
				{hierarchy: 1, text: "kk"},
				{hierarchy: 2, text: "ll"},
			},
			nil,
		},
	}

	for name, tt := range tests {
		tt := tt
		parser := NewParser()

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			for i := range tt.rows {
				ret, err := parser.Parse(tt.rows[i])
				if err != tt.wantErr {
					t.Errorf("\ngot: \n%v\nwant: \n%v", err, tt.wantErr)
				}
				if *ret != *tt.wants[i] {
					t.Errorf("\ngot: \n%v\nwant: \n%v", ret, tt.wants[i])
				}
			}
		})
	}
}

func TestParser_ParseFourSpaces(t *testing.T) {
	tests := map[string]struct {
		rows    []string
		wants   []*Markdown
		wantErr error
	}{
		"hyphen only": {
			[]string{
				"- aaabbb",
				"- cc",
				"    - bb",
			},
			[]*Markdown{
				{hierarchy: 1, text: "aaabbb"},
				{hierarchy: 1, text: "cc"},
				{hierarchy: 2, text: "bb"},
			},
			nil,
		},
		"sharp included": {
			[]string{
				"# aaabbb",
				"- cc",
				"- oo",
				"    - ff",
				"        - pp",
				"- bb",
				"# kk",
				"- ll",
			},
			[]*Markdown{
				{hierarchy: 1, text: "aaabbb"},
				{hierarchy: 2, text: "cc"},
				{hierarchy: 2, text: "oo"},
				{hierarchy: 3, text: "ff"},
				{hierarchy: 4, text: "pp"},
				{hierarchy: 2, text: "bb"},
				{hierarchy: 1, text: "kk"},
				{hierarchy: 2, text: "ll"},
			},
			nil,
		},
		"sharp included 2": {
			[]string{
				"- oo",
				"    - ff",
				"        - pp",
				"# aaabb",
				"- cc",
				"- oo",
				"    - ff",
				"        - pp",
				"- bb",
				"### kk",
				"- ll",
			},
			[]*Markdown{
				{hierarchy: 1, text: "oo"},
				{hierarchy: 2, text: "ff"},
				{hierarchy: 3, text: "pp"},
				{hierarchy: 1, text: "aaabb"},
				{hierarchy: 2, text: "cc"},
				{hierarchy: 2, text: "oo"},
				{hierarchy: 3, text: "ff"},
				{hierarchy: 4, text: "pp"},
				{hierarchy: 2, text: "bb"},
				{hierarchy: 1, text: "kk"},
				{hierarchy: 2, text: "ll"},
			},
			nil,
		},
	}

	for name, tt := range tests {
		tt := tt
		parser := NewParser()

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			for i := range tt.rows {
				ret, err := parser.Parse(tt.rows[i])
				if err != tt.wantErr {
					t.Errorf("\ngot: \n%v\nwant: \n%v", err, tt.wantErr)
				}
				if *ret != *tt.wants[i] {
					t.Errorf("\ngot: \n%v\nwant: \n%v", ret, tt.wants[i])
				}
			}
		})
	}
}
