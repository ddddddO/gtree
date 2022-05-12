package gtree

import (
	"sync"
)

type counter struct {
	n  uint
	mu sync.Mutex
}

func newCounter() *counter {
	return &counter{}
}

func (c *counter) next() uint {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.n += 1
	return c.n
}

func (c *counter) reset() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.n = 0
}

func (c *counter) current() uint {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.n
}
