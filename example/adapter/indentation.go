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
	output, _ := gtree.Execute(tab.Data, conf)
	return output
}

type TwoSpaces struct {
	Data io.Reader
}

func (ts *TwoSpaces) Execute() string {
	conf := gtree.Config{
		IsTwoSpaces: true,
	}
	output, _ := gtree.Execute(ts.Data, conf)
	return output
}

type FourSpaces struct {
	Data io.Reader
}

func (fs *FourSpaces) Execute() string {
	conf := gtree.Config{
		IsFourSpaces: true,
	}
	output, _ := gtree.Execute(fs.Data, conf)
	return output
}
