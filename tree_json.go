package gtree

import (
	"encoding/json"
	"io"
)

type jsonTree struct {
	*tree
}

// noop
func (jt *jsonTree) grow() error {
	return nil
}

func (jt *jsonTree) spread(w io.Writer) error {
	enc := json.NewEncoder(w)

	for _, root := range jt.roots {
		if err := enc.Encode(root); err != nil {
			return err
		}
	}
	return nil
}
