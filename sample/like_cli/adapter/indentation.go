package adapter

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/ddddddO/gtree"
)

// Tab is ...
type Tab struct {
	Data io.Reader
}

// Execute is ...
func (tab *Tab) Execute() error {
	buf := &strings.Builder{}
	if err := gtree.Execute(buf, tab.Data); err != nil {
		return err
	}
	fmt.Printf("%s\n\n", buf.String())
	return nil
}

// TwoSpaces is ...
type TwoSpaces struct {
	Data io.Reader
}

// Execute is ...
func (ts *TwoSpaces) Execute() error {
	buf := &bytes.Buffer{}
	if err := gtree.Execute(buf, ts.Data, gtree.IndentTwoSpaces()); err != nil {
		return err
	}
	fmt.Printf("%s\n\n", buf.String())
	return nil
}

// FourSpaces is ...
type FourSpaces struct {
	Data io.Reader
}

// Execute is ...
func (fs *FourSpaces) Execute() error {
	return gtree.Execute(os.Stdout, fs.Data, gtree.IndentFourSpaces())
}
