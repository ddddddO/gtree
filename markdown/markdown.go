package markdown

type Markdown struct {
	hierarchy uint
	text      string
}

func (m *Markdown) Hierarchy() uint {
	return m.hierarchy
}

func (m *Markdown) Text() string {
	return m.text
}

const (
	sharp = "#"

	hyphen   = "-"
	asterisk = "*"
	plus     = "+"

	space = " "
	tab   = "\t"
)

// TODO: runeへ
var symbols = map[string]struct{}{
	sharp:    {},
	hyphen:   {},
	asterisk: {},
	plus:     {},
}

func IsSymbol(k rune) bool {
	_, ok := symbols[string(k)]
	return ok
}
