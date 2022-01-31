package gtree

import (
	"io"

	toml "github.com/pelletier/go-toml/v2"
)

type tomlTree struct {
	*tree
}

func (tt *tomlTree) spread(w io.Writer) error {
	enc := toml.NewEncoder(w)

	for _, root := range tt.roots {
		if err := enc.Encode(root); err != nil {
			return err
		}
	}
	return nil
}
