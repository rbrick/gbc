package render

import (
	"github.com/rbrick/gbc/mathutil"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"image"
	"log"
	"math"
	"os"
	"sync"
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

	if err = ttf.Init(); err != nil {
		log.Panicln("failed to initialize SDL TTF:", err)
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

	font, err := ttf.OpenFont("comic.ttf", 24)

	if err != nil {
		log.Panicln("failed to open font:", err)
	}

	defer font.Close()

	isRunning := true
	var xpos float64 = 0
	z := float64(10)
	decrement := true

	segments := float64(32)
	//var movement int32 = 1
	//msg := "Baxter"
	//w, h, _ := font.SizeUTF8(msg)
	//textSurface, _ := font.RenderUTF8Solid(msg, sdl.Color(color.RGBA{
	//	R: 255,
	//	G: 0,
	//	B: 0,
	//	A: 255,
	//}))
	//texture, _ := fb.Renderer.CreateTextureFromSurface(textSurface)
	centerX := float64(fb.Bounds.Dx() / 2)
	centerY := float64(fb.Bounds.Dy() / 2)
	// main loop
	for isRunning {

		// handle events, in this case escape key and close window
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				isRunning = false
				os.Exit(0)
			case *sdl.KeyboardEvent:
				if t.Keysym.Sym == sdl.K_ESCAPE {
					isRunning = false

					os.Exit(0)
				}
			}
		}

		fb.Renderer.SetDrawColor(0xFF, 0xFF, 0xFF, 0xFF)
		fb.Renderer.Clear()

		var wg sync.WaitGroup
		wg.Add(1)

		go func() {
			sdl.Do(func() {

				outerRadius := float64(256)

				angle := 360 / segments

				for i := segments; i > 0; i-- {
					cos := math.Cos(mathutil.ToRadians(i * angle))
					sin := math.Sin(mathutil.ToRadians(i * angle))

					x := cos * 100
					y := sin * 100

					travelX := cos * outerRadius
					travelY := sin * outerRadius

					//v := mathutil.ToRadians((i-1)*angle)
					nextX := math.Cos(mathutil.ToRadians((i-1)*angle)) * outerRadius
					nextY := math.Sin(mathutil.ToRadians((i-1)*angle)) * outerRadius

					points := make([]sdl.FPoint, 4)

					points[0] = sdl.FPoint{X: float32(x + (centerX)), Y: float32(y) + float32(centerY)} // the tip
					points[1] = sdl.FPoint{X: float32(travelX + centerX), Y: float32(travelY) + float32(centerY)}
					points[2] = sdl.FPoint{X: float32(nextX + centerX), Y: float32(nextY) + float32(centerY)}
					points[3] = points[0]

					fb.Renderer.SetDrawColor(0xFF, 0, 0, 0xFF)
					fb.Renderer.DrawLinesF(points)

					//fb.Renderer.CopyF(texture, nil, &sdl.FRect{X: float32((x/2)+nextX + centerX), Y: float32(y/2+nextX + centerY), W: float32(w), H: float32(h)})
				}

				wg.Done()
			})
		}()

		wg.Wait()

		xpos++

		if z > 2 && decrement {
			z -= 0.05
		} else {

			if z <= 10 {
				decrement = false
				z += 0.05
			} else {
				decrement = true
			}
		}

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
