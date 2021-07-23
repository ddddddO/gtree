package adapter

import (
	"io"

	"github.com/ddddddO/gtree"
)

type Tab struct {
	Data io.Reader
}

func (tab *Tab) Execute() string {
	conf := gtree.Config{}
	return gtree.Execute(tab.Data, conf)
}

type TwoSpaces struct {
	Data io.Reader
}

func (ts *TwoSpaces) Execute() string {
	conf := gtree.Config{
		IsTwoSpaces: true,
	}
	return gtree.Execute(ts.Data, conf)
}

type FourSpaces struct {
	Data io.Reader
}

func (fs *FourSpaces) Execute() string {
	conf := gtree.Config{
		IsFourSpaces: true,
	}
	return gtree.Execute(fs.Data, conf)
}
