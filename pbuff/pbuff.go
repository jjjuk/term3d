package pbuff

import (
	"errors"
	"sync"
)

type PlaneBuffer struct {
	Buffer []byte
	mu     sync.Mutex
	xSize  uint
	ySize  uint
}

func New(x, y uint) *PlaneBuffer {
	return &PlaneBuffer{Buffer: make([]byte, x*y, x*y), xSize: x, ySize: y}
}

func (c *PlaneBuffer) getLinPos(x, y uint) (uint, error) {
	if x > c.xSize || y > c.ySize {
		return 0, errors.New("Position is out of range")
	}
	return y*c.xSize + x, nil
}

func (c *PlaneBuffer) Set(x, y uint, char byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	pos, err := c.getLinPos(x, y)
	if err != nil {
		return
	}
	c.Buffer[pos] = char
}

func (c *PlaneBuffer) ToString() string {
	return string(c.Buffer)
}
