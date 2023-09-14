package pbuff

import (
	"errors"
)

type PlaneBuffer struct {
	Buffer []byte
	xSize  uint
	ySize  uint
}

func New(x, y uint) *PlaneBuffer {
	return &PlaneBuffer{Buffer: make([]byte, x*y), xSize: x, ySize: y}
}

func (c *PlaneBuffer) getLinPos(x, y uint) (uint, error) {
	if x > c.xSize || y > c.ySize {
		return 0, errors.New("position is out of range")
	}
	return y*c.xSize + x, nil
}

func (c *PlaneBuffer) Set(x, y uint, char byte) {
	pos, err := c.getLinPos(x, y)
	if err != nil {
		return // maybe panic
	}
	c.Buffer[pos] = char
}

func (c *PlaneBuffer) ToString() string {
	return string(c.Buffer)
}
