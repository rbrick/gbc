package hardware

//Registers represents all the registers the GBC contains.
// GBC contains eight 8-bit registers ( A, B, C, D, E, F, H, L) and two 16-bit registers (SP (StackPointer), PC (Program Counter))
type Registers struct {
	A, B, C, D, E, F, H, L uint8
	SP, PC                 uint16
}

//CPU represents the central processing unit of the GBC
type CPU struct {
	Regs Registers
}
