package render

import (
	"github.com/veandco/go-sdl2/sdl"
	"image"
	"math"
)

type FrameBuffer struct {
	Title  string
	Bounds image.Rectangle

	Window   *sdl.Window
	Renderer *sdl.Renderer
}

func (fb *FrameBuffer) Run() {
	fb.Window, _ = sdl.CreateWindow(fb.Title, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(fb.Bounds.Dx()), int32(fb.Bounds.Dy()), sdl.WINDOW_SHOWN)
	fb.Renderer, _ = sdl.CreateRenderer(fb.Window, -1, sdl.RENDERER_ACCELERATED)
	//
	//defer fb.Renderer.Destroy()

	sdl.Do(func() {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				break
			case *sdl.KeyboardEvent:
				break
			}
		}

		fb.Renderer.Clear()

		_ = fb.Renderer.SetDrawColor(255, 125, 125, 255)
		for i := 0; i < 36; i++ {
			x := math.Cos(float64(i)*10) * 25
			y := math.Sin(float64(i)*10) * 25
			_ = fb.Renderer.DrawPointF(float32(float64(fb.Bounds.Dx()/2)+x), float32(float64(fb.Bounds.Dy()/2)+y))
			//}
		}
	})

	sdl.Do(func() {
		fb.Renderer.Present()
		sdl.Delay(1000)
	})

}

func NewFrameBuffer(w, h int, title string) *FrameBuffer {
	return &FrameBuffer{
		Bounds: image.Rect(0, 0, w, h),
		Title:  title,
	}

}
