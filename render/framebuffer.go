package render

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"image"
	"log"
	"math"
	"os"
	"sync"
)

var (
	EMPTY_COLOR = sdl.Color{}
)

type FrameBuffer struct {
	Title  string
	Bounds image.Rectangle

	Window   *sdl.Window
	Renderer *sdl.Renderer

	Pixels []sdl.Color

	RunningMutex sync.Mutex
	Running      bool
}

func (fb *FrameBuffer) Run() int {

	fb.Running = true
	var err error

	sdl.Do(func() {
		fb.Window, err = sdl.CreateWindow(fb.Title, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(fb.Bounds.Dx()), int32(fb.Bounds.Dy()), sdl.WINDOW_OPENGL)
	})

	if err != nil {
		log.Fatalf("Failed to create window: %s\n", err)
		return 1
	}

	defer func() {
		sdl.Do(func() {
			fb.Window.Destroy()
		})
	}()

	sdl.Do(func() {
		fb.Renderer, err = sdl.CreateRenderer(fb.Window, -1, sdl.RENDERER_ACCELERATED)
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
		return 2
	}
	defer func() {
		sdl.Do(func() {
			fb.Renderer.Destroy()
		})
	}()

	sdl.Do(func() {
		fb.Renderer.Clear()
	})

	for fb.Running {
		sdl.Do(func() {
			for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
				switch event.(type) {
				case *sdl.QuitEvent:
					fb.RunningMutex.Lock()
					fb.Running = false
					fb.RunningMutex.Unlock()
				}
			}

			fb.Renderer.Clear()
			fb.Renderer.SetDrawColor(0, 0, 0, 0xFF)
			fb.Renderer.FillRect(&sdl.Rect{
				X: 0,
				Y: 0,
				W: int32(fb.Bounds.Dx()),
				H: int32(fb.Bounds.Dy()),
			})
		})

		for i, c := range fb.Pixels {

			if c == EMPTY_COLOR {
				fb.Renderer.SetDrawColor(0, 0, 0, 0xFF)
				continue
			}

			x := i / fb.Bounds.Dx()
			y := i % fb.Bounds.Dx()

			sdl.Do(func() {
				fb.Renderer.SetDrawColor(c.R, c.G, c.B, c.A)
				fb.Renderer.FillRect(&sdl.Rect{H: 4, W: 4, X: int32(x), Y: int32(y)})
			})

		}

		sdl.Do(func() {
			fb.Renderer.Present()
			sdl.Delay(1000 / 60)
		})
	}

	return 0
}

func (fb *FrameBuffer) Set(x, y int, c sdl.Color) {
	index := int(math.Abs(float64(x*fb.Bounds.Dx() + y)))

	fb.Pixels[index] = c
}

func NewFrameBuffer(w, h int, title string) *FrameBuffer {
	return &FrameBuffer{
		Bounds: image.Rect(0, 0, w, h),
		Title:  title,
		Pixels: make([]sdl.Color, w*h),
	}

}
