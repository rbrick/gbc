package memory

//Memory represents data that can always be read from.
type Memory interface {
	Read(addr uint16) byte
	Raw() []byte
}

//WritableMemory represents an area of memory that can be written to
type WritableMemory interface {
	Memory
	Write(addr uint16, data byte)
}

//RAM represents the Random Access Memory. This is memory that can be written to and read from
type RAM struct {
	// the memory array we will be reading from and writing too
	memory []byte
}

func (r *RAM) Write(addr uint16, data byte) {
	r.memory[addr] = data // assign the data
}

func (r *RAM) Read(addr uint16) byte {
	return r.memory[addr] // read the data from the array
}

func (r *RAM) Raw() []byte {
	return r.memory
}

func AllocateRAM(memorySize int) WritableMemory {
	return &RAM{memory: make([]byte, memorySize)}
}

//ROM represents the Read-only Memory. This is memory that can only be read from
// Basically it is immutable.
type ROM struct {
	memory []byte
}

func (r *ROM) Read(addr uint16) byte {
	return r.memory[addr]
}

func (r *ROM) Raw() []byte {
	return r.memory
}

func AllocateROM(data []byte) Memory {
	return &ROM{memory: data}
}
