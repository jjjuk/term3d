package scenes

import (
	line2 "ter3d/geometry/d2/line"
	"ter3d/screen"
	"ter3d/vec2"

	"github.com/chewxy/math32"
)

const shades = " .:!/r(l1Z4H9W8$@"

type FrameFunction func()

func MovingLinesWithGradient() (screen.RenderFunction, FrameFunction) {
	line1 := line2.New(vec2.New(-1, -1), vec2.New(1, 1))
	line2 := line2.New(vec2.New(-1, 1), vec2.New(1, -1))

	origin := vec2.New(0, 0)

	return func(pixel *vec2.Vec2, w, h uint, r float32) byte {

			var light int

			minDist := (1 / float32(h)) / 16

			dist1 := line1.DistanceToPoint(pixel)
			dist2 := line2.DistanceToPoint(pixel)

			light1 := int(math32.Round((dist1 / (minDist * 8)))) - 1
			light2 := int(math32.Round((dist2 / (minDist * 8)))) - 1

			if light1 < light2 {
				light = light1
			} else {
				light = light2
			}

			if light > 16 {
				light = 16
			}
			if light < 0 {
				light = 0
			}
			return shades[light]
		}, func() {
			line1.RotateByOrigin(origin, math32.Pi/64, false)
			line2.RotateByOrigin(origin, math32.Pi/256, true)
		}
}

func MovingLines() (screen.RenderFunction, FrameFunction) {
	line1 := line2.New(vec2.New(-1, -1), vec2.New(1, 1))
	line2 := line2.New(vec2.New(-1, 1), vec2.New(1, -1))

	origin := vec2.New(0, 0)

	return func(pixel *vec2.Vec2, w, h uint, r float32) byte {

			var light int

			minDist := (1 / float32(h)) / 16

			dist1 := line1.DistanceToPoint(pixel)
			dist2 := line2.DistanceToPoint(pixel)

			light1 := 16 - int(math32.Round((dist1 / (minDist * 2))))
			light2 := 16 - int(math32.Round((dist2 / (minDist * 2))))

			if light1 > light2 {
				light = light1
			} else {
				light = light2
			}

			if light > 16 {
				light = 16
			}
			if light < 0 {
				light = 0
			}
			return shades[light]
		}, func() {
			line1.RotateByOrigin(origin, math32.Pi/64, false)
			line2.RotateByOrigin(origin, math32.Pi/256, true)
		}
}
