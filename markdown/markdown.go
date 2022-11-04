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
