package gtree

import (
	"encoding/json"
	"io"
)

type jsonTree struct {
	*tree
}

// noop
func (jt *jsonTree) grow() treeer {
	return jt
}

func (jt *jsonTree) expand(w io.Writer) error {
	enc := json.NewEncoder(w)

	for _, root := range jt.roots {
		if err := enc.Encode(root); err != nil {
			return err
		}
	}
	return nil
}
