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

	// 0x8000-0x97FF is the Tile RAM
	// 0x9800-0x9FFF - Background Map
	VideoRAMStart = 0x8000
	VideoRAMEnd   = 0xA000

	ExternalRAMBankStart = 0xA000
	ExternalRAMBankEnd   = 0xC000

	WorkingRAMStart = 0xC000
	WorkingRAMEnd   = 0xE000

	// These addresses appear to access the same memory
	EchoRAMStart = 0xE000
	EchoRAMEnd   = 0xFE00

	SpriteAttribMemoryStart = 0xFE00 // Also known as OAM
	SpriteAttribMemoryEnd   = 0xFEA0

	UnusableSection1Start = 0xFEA0
	UnusableSection1End   = 0xFF00

	// I definitely want to see if i can emulate link cables
	IOPortsStart = 0xFF00
	IOPortsEnd   = 0xFF80

	HighRAMStart = 0xFF80
	HighRAMEnd   = 0xFFFF

	InterruptEnableRegister = 0xFFFF

	ExternalRAMSize        = ExternalRAMBankEnd - ExternalRAMBankStart
	VideoRAMSize           = VideoRAMEnd - VideoRAMStart
	WorkingRAMSize         = WorkingRAMEnd - WorkingRAMStart
	SpriteAttribMemorySize = SpriteAttribMemoryEnd - SpriteAttribMemoryStart
	HighRAMSize            = HighRAMEnd - HighRAMStart
)

//GBC represents the hardware device of the GameBoy Color
type GBC struct {
	//Cartridge is the current cartridge file that is loaded
	//Cartridge Memory ranges from address 0x0000 to 0x8000
	Cartridge memory.Memory
	//ExternalRAM is ram that is on the cartridge. If the game had extra RAM it'll map it to this (8kB)
	ExternalRAM memory.WritableMemory
	//WorkingRAM The RAM that the game boy allows a game to use (8kB)
	WorkingRAM memory.WritableMemory
	//VideoRAM Contains tile/pixel data and the background map (8kB)
	VideoRAM memory.WritableMemory
	//
	HighRAM memory.WritableMemory
}

//ReadMemory wraps the various read memory structs we have in the GBC struct.
//This will check the address and verify it will be accessing the correct memory location
func (g *GBC) ReadMemory(addr uint16) byte {
	if addr < ROMBankEnd {
		return g.Cartridge.Read(addr)
	} else if addr >= ExternalRAMBankStart && addr < ExternalRAMBankEnd {
		return g.ExternalRAM.Read(addr - ExternalRAMBankStart)
	} else if addr >= VideoRAMStart && addr < VideoRAMEnd {
		return g.VideoRAM.Read(addr - VideoRAMStart)
	} else if (addr >= WorkingRAMStart && addr < WorkingRAMEnd) || (addr >= EchoRAMStart && addr < EchoRAMEnd) {
		// WorkingRAM and EchoRAM are basically the same thing. EchoRAM shouldn't be used but if it is, just treat it
		// as working RAM
		if addr >= EchoRAMStart && addr < EchoRAMEnd {
			return g.WorkingRAM.Read(addr - WorkingRAMStart)
		}
		return g.WorkingRAM.Read(addr - WorkingRAMStart)
	} else if addr >= SpriteAttribMemoryStart && addr < SpriteAttribMemoryEnd {
		return 0 // TODO: Implement GPU
	} else if addr >= HighRAMStart && addr < HighRAMEnd {
		return g.HighRAM.Read(addr - HighRAMStart)
	}
	return 0
}

//WriteMemory wraps the various writable memory structs we have in the GBC struct.
//This will check the address and verify it will be writing to the correct memory location
func (g *GBC) WriteMemory(addr uint16, data byte) {
	if addr >= ExternalRAMBankStart && addr < ExternalRAMBankEnd {
		g.ExternalRAM.Write(addr-ExternalRAMBankStart, data)
	} else if addr >= VideoRAMStart && addr < VideoRAMEnd {
		g.VideoRAM.Write(addr-VideoRAMStart, data)
	} else if (addr >= WorkingRAMStart && addr < WorkingRAMEnd) || (addr >= EchoRAMStart && addr < EchoRAMEnd) {
		// WorkingRAM and EchoRAM are basically the same thing. EchoRAM shouldn't be used but if it is, just treat it
		// as working RAM
		if addr >= EchoRAMStart && addr < EchoRAMEnd {
			g.WorkingRAM.Write(addr-WorkingRAMStart, data)
		} else {
			g.WorkingRAM.Write(addr-WorkingRAMStart, data)
		}
	} else if addr >= SpriteAttribMemoryStart && addr < SpriteAttribMemoryEnd {
		//-1 // TODO: Implement GPU
	} else if addr >= HighRAMStart && addr < HighRAMEnd {
		g.HighRAM.Write(addr-HighRAMStart, data)
	}
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

func (g *GBC) CartridgeName() string {
	return string(g.Cartridge.Raw()[0x0134:0x0142])
}

//NewGBC creates a new instance of the GameBoy Color
func NewGBC() *GBC {
	return &GBC{
		ExternalRAM: memory.AllocateRAM(ExternalRAMSize),
		WorkingRAM:  memory.AllocateRAM(WorkingRAMSize),
		VideoRAM:    memory.AllocateRAM(VideoRAMSize),
		HighRAM:     memory.AllocateRAM(HighRAMSize),
	}
}
