package main

import (
	"ter3d/scenes"
	"ter3d/screen"
	"time"
)

func main() {

	screen := screen.New()
	render, onFrame := scenes.MovingLinesWithGradient()
	for {
		start := time.Now().Unix()
		screen.Render(render)
		screen.DrawFrame(onFrame)

		ms := start + 16 - time.Now().Unix() // ~120fps (16 ms per frame)
		time.Sleep(time.Duration(ms) * time.Millisecond)
	}
}
