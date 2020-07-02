package render

import (
	"github.com/rbrick/gbc/mathutil"
	"github.com/veandco/go-sdl2/sdl"
	"image"
	"log"
	"math"
	"time"
)

type FrameBuffer struct {
	Title  string
	Bounds image.Rectangle

	Window   *sdl.Window
	Renderer *sdl.Renderer
}

func (fb *FrameBuffer) Run() int {
	err := sdl.Init(sdl.INIT_EVERYTHING)

	if err != nil {
		log.Panicln("failed to initialize SDL", err)
	}

	fb.Window, err = sdl.CreateWindow(fb.Title, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(fb.Bounds.Dx()), int32(fb.Bounds.Dy()), sdl.WINDOW_SHOWN)

	if err != nil {
		log.Panicln("failed to create window:", err)
	}

	defer fb.Window.Destroy()

	fb.Renderer, err = sdl.CreateRenderer(fb.Window, -1, sdl.RENDERER_ACCELERATED)

	if err != nil {
		log.Panicln("failed to create renderer for window:", err)
	}
	//
	defer fb.Renderer.Destroy()
	isRunning := true
	var xpos float64 = 0
	//var movement int32 = 1

	// main loop
	for isRunning {
		// handle events, in this case escape key and close window
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				isRunning = false
			case *sdl.KeyboardEvent:
				if t.Keysym.Sym == sdl.K_ESCAPE {
					isRunning = false
				}
			}
		}

		fb.Renderer.SetDrawColor(0xFF, 0xFF, 0xFF, 0xFF)
		fb.Renderer.Clear()

		centerX := float64(fb.Bounds.Dx() / 2)
		centerY := float64(fb.Bounds.Dy() / 2)
		outerRadius := float64(256)

		segments := float64(36)
		angle := 360 / segments

		for i := segments; i > 0; i-- {
			cos := math.Cos(mathutil.ToRadians(i * angle))
			sin := math.Sin(mathutil.ToRadians(i*angle + xpos))

			x := cos * 100
			y := sin * 100

			travelX := cos * outerRadius
			travelY := sin * outerRadius

			v := mathutil.ToRadians((i-1)*angle + xpos)
			nextX := math.Cos(v) * outerRadius
			nextY := math.Sin(v) * outerRadius

			points := make([]sdl.FPoint, 4)

			points[0] = sdl.FPoint{X: float32(x + centerX), Y: float32(y + centerY)} // the tip
			points[1] = sdl.FPoint{X: float32(travelX + centerX), Y: float32(travelY + centerY)}
			points[2] = sdl.FPoint{X: float32(nextX + centerX), Y: float32(nextY + centerY)}
			points[3] = points[0]

			fb.Renderer.SetDrawColor(0xFF, 0, 0, 0xFF)
			fb.Renderer.DrawLinesF(points)

		}
		xpos++

		fb.Renderer.Present()
		//sdl.Delay(500)

		time.Sleep(50 * time.Millisecond)
	}

	return 0

}

func NewFrameBuffer(w, h int, title string) *FrameBuffer {
	return &FrameBuffer{
		Bounds: image.Rect(0, 0, w, h),
		Title:  title,
	}

}
