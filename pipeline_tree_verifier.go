package gtree

import (
	"context"
	"fmt"
)

type defaultVerifierPipeline struct{}

func newVerifierPipeline() verifierPipeline {
	return &defaultVerifierPipeline{}
}

func (dv *defaultVerifierPipeline) verify(ctx context.Context, root <-chan *Node) <-chan error {
	fmt.Println("in verify pipeline")
	return nil
}
