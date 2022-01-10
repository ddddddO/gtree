package gtree

import (
	"io"

	"gopkg.in/yaml.v2"
)

type yamlTree struct {
	*tree
}

func (yt *yamlTree) expand(w io.Writer) error {
	enc := yaml.NewEncoder(w)

	for _, root := range yt.roots {
		if err := enc.Encode(root); err != nil {
			return err
		}
	}
	return nil
}
