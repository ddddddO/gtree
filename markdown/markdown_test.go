package markdown

import (
	"testing"
)

func Test_Markdown(t *testing.T) {
	tests := map[string]struct {
		hierarchy uint
		text      string
	}{
		"hierarchy root": {
			hierarchy: 1,
			text:      "Root node",
		},
		"hierarchy 2": {
			hierarchy: 2,
			text:      "child2",
		},
		"hierarchy 3": {
			hierarchy: 3,
			text:      "child3",
		},
	}

	for name, tt := range tests {
		tt := tt
		md := &Markdown{hierarchy: tt.hierarchy, text: tt.text}

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			if md.Hierarchy() != tt.hierarchy {
				t.Errorf("\ngot: \n%d\nwant: \n%d", md.Hierarchy(), tt.hierarchy)
			}
			if md.Text() != tt.text {
				t.Errorf("\ngot: \n%s\nwant: \n%s", md.Text(), tt.text)
			}
		})
	}
}

func Test_IsSymbol(t *testing.T) {
	tests := map[string]struct {
		symbol string
		want   bool
	}{
		"# is symbol": {
			symbol: sharp,
			want:   true,
		},
		"- is symbol": {
			symbol: hyphen,
			want:   true,
		},
		"* is symbol": {
			symbol: asterisk,
			want:   true,
		},
		"+ is symbol": {
			symbol: plus,
			want:   true,
		},
		"space is not symbol": {
			symbol: space,
			want:   false,
		},
		"tab is not symbol": {
			symbol: tab,
			want:   false,
		},
		"a is not symbol": {
			symbol: "a",
			want:   false,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			if IsSymbol(tt.symbol) != tt.want {
				t.Errorf("\ngot: \n%t\nwant: \n%t", IsSymbol(tt.symbol), tt.want)
			}
		})
	}
}
