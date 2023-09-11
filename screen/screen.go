package screen

/*
#include <windows.h>
*/
import "C"

import (
	"log"
	"os"

	"golang.org/x/term"

	"ter3d/pbuff"
	"ter3d/vec2"
	"ter3d/win"
)

type Screen struct {
	terminal *win.Win
	width    uint
	height   uint
	ratio    float32

	Frame *pbuff.PlaneBuffer
}

func New() *Screen {
	width, height, err := term.GetSize(int(os.Stdout.Fd()))

	if err != nil {
		log.Fatal(err)
	}

	return &Screen{
		terminal: win.New(),
		width:    uint(width),
		height:   uint(height),
		ratio:    float32(width) / float32(height),

		Frame: pbuff.New(uint(width), uint(height)),
	}
}

func (s *Screen) GetPixelVector(x, y uint) *vec2.Vec2 {
	return vec2.New((float32(x) / float32(s.width)), (float32(y) / float32(s.height)))
}

type RenderFunction func(pixel *vec2.Vec2, w, h uint) byte

func (s *Screen) Render(fn RenderFunction) {
	for y := uint(0); y < s.height; y++ {

		for x := uint(0); x < s.width; x++ {
			pixel := s.GetPixelVector(x, y)
			s.Frame.Set(x, y, fn(pixel, s.width, s.height))
		}

	}
}

func (s *Screen) DrawFrame() {
	s.terminal.WriteTermOutput(s.Frame.ToString())
}
