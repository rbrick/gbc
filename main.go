package main

import (
	"fmt"
	"github.com/rbrick/gbc/hardware"
	"log"
)

func main() {

	fmt.Println("bytes:", hardware.HighRAMEnd-hardware.HighRAMStart)

	gameboyColor := hardware.NewGBC()

	if err := gameboyColor.LoadCartridge("zelda.gb"); err != nil {
		log.Panicln(err)
	}

	fmt.Println(gameboyColor.CartridgeName())
}
