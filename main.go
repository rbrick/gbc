package main

import (
	"fmt"
	"github.com/rbrick/gbc/hardware"
	"github.com/rbrick/gbc/render"
	"github.com/veandco/go-sdl2/sdl"
	"log"
	"os"
)

func main() {

	fmt.Println("bytes:", hardware.HighRAMEnd-hardware.HighRAMStart)

	gameboyColor := hardware.NewGBC()

	if err := gameboyColor.LoadCartridge("zelda.gb"); err != nil {
		log.Panicln(err)
	}

	fmt.Println(gameboyColor.CartridgeName())

	framebuffer := render.NewFrameBuffer(420, 420, "Simple Physics")

	for i := 0; i < 420; i++ {
		framebuffer.Set(int(i), int(i), sdl.Color{
			R: 255,
			G: 0,
			B: 0,
			A: 255,
		})
	}

	exitCode := 0

	sdl.Main(func() {
		exitCode = framebuffer.Run()
	})

	os.Exit(exitCode)

}
