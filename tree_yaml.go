package gtree

import (
	"io"

	"gopkg.in/yaml.v2"
)

type yamlTree struct {
	*tree
}

// noop
func (yt *yamlTree) grow() error {
	return nil
}

func (yt *yamlTree) spread(w io.Writer) error {
	enc := yaml.NewEncoder(w)

	for _, root := range yt.roots {
		if err := enc.Encode(root); err != nil {
			return err
		}
	}
	return nil
}
