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

// Output is ...
func (tab *Tab) Output() error {
	buf := &strings.Builder{}
	if err := gtree.Output(buf, tab.Data); err != nil {
		return err
	}
	fmt.Printf("%s\n\n", buf.String())
	return nil
}

// TwoSpaces is ...
type TwoSpaces struct {
	Data io.Reader
}

// Output is ...
func (ts *TwoSpaces) Output() error {
	buf := &bytes.Buffer{}
	if err := gtree.Output(buf, ts.Data, gtree.IndentTwoSpaces()); err != nil {
		return err
	}
	fmt.Printf("%s\n\n", buf.String())
	return nil
}

// FourSpaces is ...
type FourSpaces struct {
	Data io.Reader
}

// Output is ...
func (fs *FourSpaces) Output() error {
	return gtree.Output(os.Stdout, fs.Data, gtree.IndentFourSpaces())
}
