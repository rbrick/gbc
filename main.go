package main

import (
	"github.com/rbrick/gbc/hardware"
	"log"
)

func main() {
	gameboyColor := hardware.NewGBC()

	if err := gameboyColor.LoadCartridge("zelda.gbc"); err != nil {
		log.Panicln(err)
	}
}
