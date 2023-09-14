package circle

import (
	"ter3d/vec2"

	"github.com/chewxy/math32"
)

type Circle struct {
	cen *vec2.Vec2
	r   float32
}

func New(cen *vec2.Vec2, r float32) *Circle {
	return &Circle{cen: cen, r: r}
}

func (c *Circle) DistanceToPoint(p *vec2.Vec2) float32 {
	prod := vec2.SubTwo(p, c.cen)
	return math32.Sqrt(prod[0]*prod[0]+prod[1]*prod[1]) - c.r
}

func (c *Circle) RotateByOrigin(or *vec2.Vec2, ang float32, clockwise bool) {
	c.cen.RotateByPoint(or, ang, clockwise)
}
