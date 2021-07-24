package adapter

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/ddddddO/gtree"
)

type Tab struct {
	Data io.Reader
}

func (tab *Tab) Execute() error {
	conf := gtree.Config{}
	buf := &strings.Builder{}
	if err := gtree.Execute(buf, tab.Data, conf); err != nil {
		return err
	}
	fmt.Printf("%s\n\n", buf.String())
	return nil
}

type TwoSpaces struct {
	Data io.Reader
}

func (ts *TwoSpaces) Execute() error {
	conf := gtree.Config{
		IsTwoSpaces: true,
	}
	buf := &bytes.Buffer{}
	if err := gtree.Execute(buf, ts.Data, conf); err != nil {
		return err
	}
	fmt.Printf("%s\n\n", buf.String())
	return nil
}

type FourSpaces struct {
	Data io.Reader
}

func (fs *FourSpaces) Execute() error {
	conf := gtree.Config{
		IsFourSpaces: true,
	}
	return gtree.Execute(os.Stdout, fs.Data, conf)
}
