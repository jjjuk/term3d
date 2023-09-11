package line2_test

import (
	line2 "ter3d/geometry/d2/line"
	"ter3d/vec2"
	"testing"
)

func TestLineCenter(t *testing.T) {
	line := line2.New(vec2.New(3, 0), vec2.New(1, 4))

	center := line.Center()

	x, y := center.GetRoundedXY()

	if x != 2 || y != 2 {
		t.Fatalf("Expected center to be (2,2) got (%v,%v)", x, y)
	}
}
