package main

import (
	"fmt"
	"github.com/rbrick/gbc/hardware"
	"github.com/rbrick/gbc/render"
	"github.com/veandco/go-sdl2/sdl"
	"log"
	"sync"
)

func main() {

	fmt.Println("bytes:", hardware.HighRAMEnd-hardware.HighRAMStart)

	gameboyColor := hardware.NewGBC()

	if err := gameboyColor.LoadCartridge("zelda.gb"); err != nil {
		log.Panicln(err)
	}

	fmt.Println(gameboyColor.CartridgeName())

	fb := render.NewFrameBuffer(800, 600, "Test")

	sdl.Main(func() {
		fb.Run()
	})

	var wg sync.WaitGroup
	wg.Add(1)

	wg.Wait()
}
