package hardware

import (
	"github.com/rbrick/gbc/memory"
	"io/ioutil"
)

// This is the layout of the memory map.
// The GameBoy Color has a 64kB memory space
// (0x0000-0xFFFF)

// Taken from http://marc.rawer.de/Gameboy/Docs/GBCPUman.pdf
const (
	ROMBankStart = 0x0000
	ROMBankEnd   = 0x8000

	VideoRAMStart = 0x8000
	VideoRAMEnd   = 0xA000

	SwitchableRAMBankStart = 0xA000
	SwitchableRAMBankEnd   = 0xC000

	Internal8kBRAMStart = 0xC000
	Internal8kBRAMEnd   = 0xE000

	// These addresses appear to access the same memory
	EchoInternal8kBRAMStart = 0xE000
	EchoInternal8kBRAMEnd   = 0xFE00

	SpriteAttribMemoryStart = 0xFE00 // Also known as OAM
	SpriteAttribMemoryEnd   = 0xFEA0

	UnusableSection1Start = 0xFEA0
	UnusableSection1End   = 0xFF00

	// I definitely want to see if i can emulate link cables
	IOPortsStart = 0xFF00
	IOPortsEnd   = 0xFF4C

	UnusableSection2Start = 0xFF4C
	UnusableSection2End   = 0xFF80

	HighRAMStart = 0xFF80
	HighRAMEnd   = 0xFFFE

	InterruptEnableRegister = 0xFFFF
)

//GBC represents the hardware device of the GameBoy Color
type GBC struct {
	//Cartridge is the current cartridge file that is loaded
	//Cartridge Memory ranges from address 0x0000 to 0x8000
	Cartridge memory.Memory
}

//ReadMemory wraps the various read memory structs we have in the GBC struct.
//This will check the address and verify it will be accessing the correct memory location
func (g *GBC) ReadMemory(addr uint16) byte {
	if addr < ROMBankEnd {
		return g.Cartridge.Read(addr)
	}
	return 0
}

//WriteMemory wraps the various writable memory structs we have in the GBC struct.
//This will check the address and verify it will be writing to the correct memory location
func (g *GBC) WriteMemory(addr uint16, data byte) {
	return
}

//Reset will reset the GBC.
//This will be used for loading/swapping cartridges if games are running
func (g *GBC) Reset() {
	return
}

//LoadCartridge loads a GBC game from a file.
func (g *GBC) LoadCartridge(path string) error {
	data, err := ioutil.ReadFile(path)

	if err != nil {
		return err
	}

	g.Cartridge = memory.AllocateROM(data)
	return nil
}

//NewGBC creates a new instance of the GameBoy Color
func NewGBC() *GBC {
	return &GBC{}
}
