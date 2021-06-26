package adapter

import (
	"io"

	"github.com/ddddddO/gentree"
)

type Tab struct {
	Data io.Reader
}

func (tab *Tab) Execute() string {
	conf := gentree.Config{}
	return gentree.Execute(tab.Data, conf)
}

type TwoSpaces struct {
	Data io.Reader
}

func (ts *TwoSpaces) Execute() string {
	conf := gentree.Config{
		IsTwoSpaces: true,
	}
	return gentree.Execute(ts.Data, conf)
}

type FourSpaces struct {
	Data io.Reader
}

func (fs *FourSpaces) Execute() string {
	conf := gentree.Config{
		IsFourSpaces: true,
	}
	return gentree.Execute(fs.Data, conf)
}
