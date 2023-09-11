package vec2

import "github.com/chewxy/math32"

type Vec2 [2]float32

func New(x, y float32) *Vec2 {
	return &Vec2{x, y}
}

func (v *Vec2) Add(ov *Vec2) *Vec2 {
	v[0] += ov[0]
	v[1] += ov[1]
	return v
}

func (v *Vec2) Sub(ov *Vec2) *Vec2 {
	v[0] -= ov[0]
	v[1] -= ov[1]
	return v
}

func (v *Vec2) Mul(ov *Vec2) *Vec2 {
	v[0] *= ov[0]
	v[1] *= ov[1]
	return v
}

func (v *Vec2) Div(ov *Vec2) *Vec2 {
	v[0] /= ov[0]
	v[1] /= ov[1]
	return v
}

func (v *Vec2) Abs() *Vec2 {
	v[0] = math32.Abs(v[0])
	v[1] = math32.Abs(v[1])
	return v
}

func (v *Vec2) AddScal(s float32) *Vec2 {
	v[0] += s
	v[1] += s
	return v
}

func (v *Vec2) SubScal(s float32) *Vec2 {
	v[0] -= s
	v[1] -= s
	return v
}

func (v *Vec2) MulScal(s float32) *Vec2 {
	v[0] *= s
	v[1] *= s
	return v
}

func (v *Vec2) DivScal(s float32) *Vec2 {
	v[0] /= s
	v[1] /= s
	return v
}

func (v *Vec2) GetRoundedXY() (float32, float32) {
	return math32.Round(v[0]), math32.Round(v[1])
}

func (v *Vec2) Cross(ov *Vec2) float32 {
	return (v[0] * ov[1]) - (ov[0] * v[1])
}

func (v *Vec2) Length() float32 {
	return math32.Sqrt(v[0]*v[0] + v[1]*v[1])
}
