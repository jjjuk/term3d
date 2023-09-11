package vec2

import "github.com/chewxy/math32"

/*
Rotates vector by a (0,0) base.

Clockwise rotation: clockwise = true

Counter clockwise rotation: clockwise = false
*/
func (v *Vec2) Rotate(ang float32, clockwise bool) {
	var x, y float32
	s, c := math32.Sincos(ang)

	if clockwise {
		x = v[0]*c + v[1]*s
		y = -v[0]*s + v[1]*c
	} else {
		x = v[0]*c - v[1]*s
		y = v[0]*s + v[1]*c
	}

	v[0] = x
	v[1] = y
}

/*
Rotates vector by a base of a given point.

Clockwise rotation: clockwise = true

Counter clockwise rotation: clockwise = false
*/
func (v *Vec2) RotateByPoint(or *Vec2, ang float32, clockwise bool) {
	v.Sub(or)
	v.Rotate(ang, clockwise)
	v.Add(or)
}
