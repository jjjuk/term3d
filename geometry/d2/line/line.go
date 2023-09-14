package line2

import (
	"ter3d/vec2"

	"github.com/chewxy/math32"
)

type Line2 [2]*vec2.Vec2

func New(x, y *vec2.Vec2) *Line2 {
	return &Line2{x, y}
}

/*
Function finds distance between given line and 2d vector (as point).

Formula:
https://math.stackexchange.com/questions/2757318/distance-between-a-point-and-a-line-defined-by-2-points
*/
func (l *Line2) DistanceToPoint(p *vec2.Vec2) float32 {
	return math32.Abs(vec2.SubTwo(l[1], l[0]).Cross(vec2.SubTwo(p, l[0]))) /
		vec2.SubTwo(l[1], l[0]).Length()
}

func (l *Line2) Center() *vec2.Vec2 {
	return vec2.AddTwo(l[0], l[1]).DivScal(2)
}

func (l *Line2) RotateByOrigin(or *vec2.Vec2, ang float32, clockwise bool) {
	l[0].RotateByPoint(or, ang, clockwise)
	l[1].RotateByPoint(or, ang, clockwise)
}

func (l *Line2) RotateByCenter(ang float32, clockwise bool) {
	l.RotateByOrigin(l.Center(), ang, clockwise)
}
