package gtree

import (
	"io"

	"github.com/vmihailenco/msgpack/v5"
)

type msgpkTree struct {
	*tree
}

// noop
func (mt *msgpkTree) grow() treeer {
	return mt
}

func (mt *msgpkTree) expand(w io.Writer) error {
	enc := msgpack.NewEncoder(w)

	for _, root := range mt.roots {
		if err := enc.Encode(root); err != nil {
			return err
		}
	}
	return nil
}
