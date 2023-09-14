package screen

import (
	"log"
	"os"
	"sync"

	"golang.org/x/term"

	"ter3d/pbuff"
	"ter3d/terminal"
	"ter3d/vec2"
)

const termSymRatio = 0.5

type Screen struct {
	terminal *terminal.Win
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
		terminal: terminal.NewWinTerm(),
		width:    uint(width),
		height:   uint(height),
		ratio:    float32(width) / float32(height),

		Frame: pbuff.New(uint(width), uint(height)),
	}
}

func (s *Screen) getVX(x uint) float32 {
	return (-1 + float32(x)/(float32(s.width)/2))
}

func (s *Screen) getVY(y uint) float32 {
	return (1 - float32(y)/(float32(s.height)/2))
}

func (s *Screen) getPixelVector(x, y uint) *vec2.Vec2 {
	return vec2.New(s.getVX(x)*s.ratio*termSymRatio, s.getVY(y))
}

type RenderFunction func(pixel *vec2.Vec2, w, h uint, r float32) byte

func (s *Screen) Render(fn RenderFunction) {
	var wg sync.WaitGroup
	for y := uint(0); y < s.height; y++ {
		for x := uint(0); x < s.width; x++ {
			wg.Add(1)
			go func(x, y uint) {
				defer wg.Done()
				pixel := s.getPixelVector(x, y)
				s.Frame.Set(x, y, fn(pixel, s.width, s.height, s.ratio))
			}(x, y)
		}
	}
	wg.Wait()
}

func (s *Screen) DrawFrame(onFrame func()) {
	s.terminal.WriteTermOutput(s.Frame.ToString())
	onFrame()
}
