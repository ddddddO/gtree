package gtree

import (
	"strings"
)

type fileConsiderer struct {
	extensions []string
}

func newFileConsiderer(extensions []string) *fileConsiderer {
	return &fileConsiderer{
		extensions: extensions,
	}
}

func (fc *fileConsiderer) isFile(current *Node) bool {
	if current.hasChild() {
		return false
	}

	for _, e := range fc.extensions {
		if strings.HasSuffix(current.name, e) {
			return true
		}
	}
	return false
}
