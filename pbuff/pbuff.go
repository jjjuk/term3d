package pbuff

import (
	"fmt"
	"log"
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
		return 0, fmt.Errorf("position (%v,%v) is out of range. max position is (%v,%v)", x, y, c.xSize, c.ySize)
	}
	return y*c.xSize + x, nil
}

func (c *PlaneBuffer) Set(x, y uint, char byte) {
	pos, err := c.getLinPos(x, y)
	if err != nil {
		log.Panic(err)
	}
	c.Buffer[pos] = char
}

func (c *PlaneBuffer) ToString() string {
	return string(c.Buffer)
}
