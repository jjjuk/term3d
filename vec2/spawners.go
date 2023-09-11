package vec2

func AddTwo(a, b *Vec2) *Vec2 {
	return &Vec2{a[0] + b[0], a[1] + b[1]}
}
func SubTwo(a, b *Vec2) *Vec2 {
	return &Vec2{a[0] - b[0], a[1] - b[1]}
}
func MulTwo(a, b *Vec2) *Vec2 {
	return &Vec2{a[0] * b[0], a[1] * b[1]}
}
func DivTwo(a, b *Vec2) *Vec2 {
	return &Vec2{a[0] / b[0], a[1] / b[1]}
}
