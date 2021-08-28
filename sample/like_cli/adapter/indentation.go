package adapter

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/ddddddO/gtree/v6"
)

type Tab struct {
	Data io.Reader
}

func (tab *Tab) Execute() error {
	buf := &strings.Builder{}
	if err := gtree.Execute(buf, tab.Data); err != nil {
		return err
	}
	fmt.Printf("%s\n\n", buf.String())
	return nil
}

type TwoSpaces struct {
	Data io.Reader
}

func (ts *TwoSpaces) Execute() error {
	buf := &bytes.Buffer{}
	if err := gtree.Execute(buf, ts.Data, gtree.IndentTwoSpaces()); err != nil {
		return err
	}
	fmt.Printf("%s\n\n", buf.String())
	return nil
}

type FourSpaces struct {
	Data io.Reader
}

func (fs *FourSpaces) Execute() error {
	return gtree.Execute(os.Stdout, fs.Data, gtree.IndentFourSpaces())
}
