package vec2_test

import (
	"ter3d/vec2"
	"testing"

	"github.com/chewxy/math32"
)

func TestScalarDivision(t *testing.T) {
	vec := vec2.New(6, 6)
	vec.DivScal(2)
	x, y := vec.GetRoundedXY()

	if x != 3 || y != 3 {
		t.Fatalf(`Expected vec to be (2, 2), got (%v, %v)`, x, y)
	}
}

func TestRotation(t *testing.T) {
	vec := vec2.New(1, 3)

	vec.Rotate(90*(math32.Pi/180), false)

	x, y := vec.GetRoundedXY()

	if x != -3 || y != 1 {
		t.Fatalf(`Expected vec to be (-3, 1), got (%v, %v)`, x, y)
	}

}
