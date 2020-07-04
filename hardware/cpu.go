package hardware

const (
	FlagCarryPosition = iota + 4
	FlagHalfCarryPosition
	FlagSubtractPosition
	FlagZeroPosition
)

type FlagRegister uint8

func (f FlagRegister) Test(position int) bool {
	return (f>>position)&1 != 0
}

func (f *FlagRegister) Set(position int, v bool) {
	if v {
		*f = *f | (1 << position)
	} else {
		*f = FlagRegister(uint8(*f) & uint8(^position))
	}

}

//Registers represents all the registers the GBC contains.
// GBC contains eight 8-bit registers ( A, B, C, D, E, F, H, L) and two 16-bit registers (SP (StackPointer), PC (Program Counter))
// While the GameBoy doesn't have 16-bit registers, games can group two registers together for "virtual" 16-bit registers
type Registers struct {
	A, B, C, D, E, H, L uint8
	Flags               FlagRegister // F = flags
	SP, PC              uint16
}

func (r *Registers) ReadGrouped(r0, r1 uint8) uint16 {
	return uint16(r0)<<8 | uint16(r1)
}

func (r *Registers) SetGrouped(r0, r1 *uint8, value uint16) {
	*r0 = uint8((value & 0xFF00) >> 8)
	*r1 = uint8((value & 0xFF) >> 8)
}

func (r *Registers) ReadAF() uint16 {
	return r.ReadGrouped(r.A, uint8(r.Flags))
}

func (r *Registers) SetAF(val uint16) {
	r.SetGrouped(&r.A, (*uint8)(&r.Flags), val)
}

func (r *Registers) ReadBC() uint16 {
	return r.ReadGrouped(r.B, r.C)
}

func (r *Registers) SetBC(val uint16) {
	r.SetGrouped(&r.B, &r.C, val)
}

func (r *Registers) ReadDE() uint16 {
	return r.ReadGrouped(r.D, r.E)
}

func (r *Registers) SetDE(val uint16) {
	r.SetGrouped(&r.D, &r.E, val)
}

func (r *Registers) ReadHL() uint16 {
	return r.ReadGrouped(r.H, r.L)
}

func (r *Registers) SetHL(val uint16) {
	r.SetGrouped(&r.H, &r.L, val)
}

//CPU represents the central processing unit of the GBC
type CPU struct {
	Regs Registers
}

func (c *CPU) HandleInterrupts(gbc *GBC) {

}

func (c *CPU) Step(gbc *GBC) {
	c.HandleInterrupts(gbc)

	opcode := c.Read(gbc)

	switch opcode {
	// GIANT CASE STATEMENT? no.
	case 0x00: // NO-OP
		return
	}
}

func (c *CPU) Read(gbc *GBC) byte {
	b := gbc.ReadMemory(c.Regs.PC)
	c.Regs.PC += 1
	return b
}
