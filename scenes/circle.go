package scenes

import (
	"ter3d/geometry/d2/circle"
	"ter3d/screen"
	"ter3d/vec2"

	"github.com/chewxy/math32"
)

func MovingCircle() (screen.RenderFunction, FrameFunction) {
	circle := circle.New(vec2.New(0.6, 0), 0.15)

	origin := vec2.New(0, 0)

	return func(pixel *vec2.Vec2, w, h uint, r float32) byte {

			var light int

			minDist := (1 / float32(h)) / 16

			dist := circle.DistanceToPoint(pixel)

			light = int(math32.Round((dist / (minDist * 8)))) - 1

			if light > 16 {
				light = 16
			}
			if light < 0 {
				light = 0
			}

			return shades[light]
		}, func() {
			circle.RotateByOrigin(origin, math32.Pi/64, false)
		}
}
