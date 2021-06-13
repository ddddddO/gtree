package adapter

import (
	"io"

	"github.com/ddddddO/gentree"
)

const (
	enableTwoSpaces  = true
	enableFourSpaces = true
)

type Tab struct {
	Data io.Reader
}

func (tab *Tab) Execute() string {
	return gentree.Execute(tab.Data, !enableTwoSpaces, !enableFourSpaces)
}

type TwoSpaces struct {
	Data io.Reader
}

func (ts *TwoSpaces) Execute() string {
	return gentree.Execute(ts.Data, enableTwoSpaces, !enableFourSpaces)
}

type FourSpaces struct {
	Data io.Reader
}

func (fs *FourSpaces) Execute() string {
	return gentree.Execute(fs.Data, !enableTwoSpaces, enableFourSpaces)
}
