package instructions

import (
	"github.com/markelmencia/gogb/cpu"
	"github.com/markelmencia/gogb/emulator"
)

/* AUXILIARY FUNCTIONS */

// Returns true if there is a carry
// in bit 7 of the resulting value of a + b.
func hasCarry(a, b byte) bool {
	// To check the 7-th bit carry we can simply check if the result
	// of the sum of these two values exceeds 0xFF (8-bit overflow).

	// The casting to uint16 is necessary because if the sum of two
	// bytes overflows it resets back to zero.
	result := uint16(a) + uint16(b)
	return result > 0xFF
}

// Returns true if there is a carry
// in bit 3 of the resulting value of a + b.
func hasHalfCarry(a, b byte) bool {
	// Masks the lower 4 bits of a and b
	lo4a := a & 0x0F
	lo4b := b & 0x0F

	// If the sum of lo4a + lo4b overflows
	// in bit 3 (the number is bigger than 0xF),
	// it means that there was a carry in bit 3.

	result := lo4a + lo4b
	return result > 0x0F
}

/* SM89 INSTRUCTIONS */

// LD r, r': Load register (register) (8-Bit)
//
// Loads the value of r' into r.
func LDrr(dst, src cpu.Halve, emu emulator.Emulation) {
	v := emu.CPU.GetHalve(src)
	emu.CPU.SetHalve(dst, v)
	emu.CPU.PC++
}

// LD r, n: Load register (immediate)
//
// Loads n (the value in memory next to the instruction)
// into register r.
func LDra(dst cpu.Halve, emu emulator.Emulation) {
	emu.CPU.PC++
	v := emu.RAM.GetByte(emu.CPU.PC)
	emu.CPU.SetHalve(dst, v)
	emu.CPU.PC++
}

// LD r, (HL): Load register (indirect HL)
//
// Loads the memory value in the index inside register
// HL (16 bits) into r.
func LDrHL(dst cpu.Halve, emu emulator.Emulation) {
	a := emu.CPU.HL
	v := emu.RAM.GetByte(a)
	emu.CPU.SetHalve(dst, v)
	emu.CPU.PC++
}

// LD (HL), r: Load from register (indirect HL)
//
// Writes the value in register r into the memory
// address specified in HL.
func LDHLr(src cpu.Halve, emu emulator.Emulation) {
	a := emu.CPU.HL
	v := emu.CPU.GetHalve(src)
	emu.RAM.SetByte(v, a)
	emu.CPU.PC++
}

// LD (HL), n: Load from immediate data (indirect HL)
//
// Writes the value of the memory address next to
// the instruction into the memory address specified in HL.
func LDHLn(emu emulator.Emulation) {
	a := emu.CPU.HL
	emu.CPU.PC++
	v := emu.RAM.GetByte(emu.CPU.PC)
	emu.RAM.SetByte(v, a)
	emu.CPU.PC++
}

// LD A, (BC): Load accumulator (indirect BC)
//
// Loads the memory value specified in BC into A.
func LDaBC(emu emulator.Emulation) {
	a := emu.CPU.BC
	v := emu.RAM.GetByte(a)
	emu.CPU.SetHalve(cpu.A, v)
	emu.CPU.PC++
}

// LD A, (DE): Load accumulator (indirect DE)
//
// Loads the memory value specified in DE into A.
func LDaDE(emu emulator.Emulation) {
	a := emu.CPU.DE
	v := emu.RAM.GetByte(a)
	emu.CPU.SetHalve(cpu.A, v)
	emu.CPU.PC++
}

// LD (BC), A: Load accumulator (indirect BC)
//
// Writes the value of register A into the
// address specified in BC.
func LDBCa(emu emulator.Emulation) {
	a := emu.CPU.BC
	v := emu.CPU.GetHalve(cpu.A)
	emu.RAM.SetByte(v, a)
	emu.CPU.PC++
}

// LD (DE), A: Load accumulator (indirect DE)
//
// Writes the value of register A into the
// address specified in BC.
func LDDEa(emu emulator.Emulation) {
	a := emu.CPU.DE
	v := emu.CPU.GetHalve(cpu.A)
	emu.RAM.SetByte(v, a)
	emu.CPU.PC++
}

// LD A, (nn): Load accumulator (direct)
//
// Loads into A the memory data of the
// address obtained from the next two
// RAM values of the instruction.
func LDAnn(emu emulator.Emulation) {
	emu.CPU.PC++
	nLo := emu.RAM.GetByte(emu.CPU.PC)
	emu.CPU.PC++
	nHi := emu.RAM.GetByte(emu.CPU.PC)

	a := uint16(nHi)<<8 | uint16(nLo)
	v := emu.RAM.GetByte(a)

	emu.CPU.SetHalve(cpu.A, v)
	emu.CPU.PC++
}

// LD (nn), A: Load from accumulator (direct)
//
// Writes into the memory address
// specified by the next two RAM bytes
// of the instruction the value of A.
func LDnnA(emu emulator.Emulation) {
	emu.CPU.PC++
	nLo := emu.RAM.GetByte(emu.CPU.PC)
	emu.CPU.PC++
	nHi := emu.RAM.GetByte(emu.CPU.PC)

	a := uint16(nHi)<<8 | uint16(nLo)
	v := emu.CPU.GetHalve(cpu.A)

	emu.RAM.SetByte(v, a)
	emu.CPU.PC++
}

// LDH A, (C): Load accumulator (indirect 0xFF00+C)
//
// Loads the value in memory of the address 0xFF00 + C
// into A.
func LDHaC(emu emulator.Emulation) {
	a := 0xFF00 | uint16(emu.CPU.GetHalve(cpu.C))
	v := emu.RAM.GetByte(a)
	emu.CPU.SetHalve(cpu.A, v)
	emu.CPU.PC++
}

// LDH (C), A: Load from accumulator (indirect 0xFF00+C)
//
// Loads the value of A into the memory address 0xFF00 + C.
func LDHCa(emu emulator.Emulation) {
	a := 0xFF00 | uint16(emu.CPU.GetHalve(cpu.C))
	v := emu.CPU.GetHalve(cpu.A)
	emu.RAM.SetByte(v, a)
	emu.CPU.PC++
}

// LDH A, (n): Load accumulator (indirect 0xFF00+n)
//
// Loads the value memory in 0xFF00 + n (next value
// in memory from the instruction) into A.
func LDHAn(emu emulator.Emulation) {
	emu.CPU.PC++
	a := 0xFF00 | uint16(emu.RAM.GetByte(emu.CPU.PC))
	v := emu.RAM.GetByte(a)
	emu.CPU.SetHalve(cpu.A, v)
	emu.CPU.PC++
}

// LDH (n), A: Load from accumulator (indirect 0xFF00+n)
//
// Loads the value of A into the memory address 0xFF + n.
func LDHnA(emu emulator.Emulation) {
	emu.CPU.PC++
	a := 0xFF00 | uint16(emu.RAM.GetByte(emu.CPU.PC))
	v := emu.CPU.GetHalve(cpu.A)
	emu.RAM.SetByte(v, a)
	emu.CPU.PC++
}

// LD A, (HL-): Load accumulator (indirect HL, decrement)
//
// Loads the memory value in the specified index at HL
// into the register A. Then, HL is decremented by 1.
func LDaHLm(emu emulator.Emulation) {
	a := emu.CPU.HL
	v := emu.RAM.GetByte(a)
	emu.CPU.SetHalve(cpu.A, v)
	emu.CPU.HL--
	emu.CPU.PC++
}

// LD (HL-), A: Load from accumulator (indirect HL, decrement)
//
// Loads into the memory position in HL
// the value in register A, then decrements HL.
func LDHLam(emu emulator.Emulation) {
	a := emu.CPU.HL
	v := emu.CPU.GetHalve(cpu.A)
	emu.RAM.SetByte(v, a)
	emu.CPU.HL--
	emu.CPU.PC++
}

// LD A, (HL+): Load accumulator (indirect HL, increment)
//
// Loads the memory value in the specified index at HL
// into the register A. Then, HL is incremented by 1.
func LDaHLp(emu emulator.Emulation) {
	a := emu.CPU.HL
	v := emu.RAM.GetByte(a)
	emu.CPU.SetHalve(cpu.A, v)
	emu.CPU.HL++
	emu.CPU.PC++
}

// LD (HL-), A: Load from accumulator (indirect HL, increment)
//
// Loads into the memory position in HL
// the value in register A, then increments HL.
func LDHLap(emu emulator.Emulation) {
	a := emu.CPU.HL
	v := emu.CPU.GetHalve(cpu.A)
	emu.RAM.SetByte(v, a)
	emu.CPU.HL++
	emu.CPU.PC++
}

// LD rr, nn: Load 16-bit register / register pair
//
// Loads into rr the immediate data in the next
// two registers from the instruction.
func LDrrnn(rr cpu.Register, emu emulator.Emulation) {
	emu.CPU.PC++
	nLo := emu.RAM.GetByte(emu.CPU.PC)
	emu.CPU.PC++
	nHi := emu.RAM.GetByte(emu.CPU.PC)
	v := uint16(nHi)<<8 | uint16(nLo)
	emu.CPU.SetReg(rr, v)
	emu.CPU.PC++
}

// LD (nn), SP: Load from stack pointer (direct)
//
// Loads into the memory address defined in nn the
// value inside SP.
func LDnnSP(emu emulator.Emulation) {
	emu.CPU.PC++
	nLo := emu.RAM.GetByte(emu.CPU.PC)
	emu.CPU.PC++
	nHi := emu.RAM.GetByte(emu.CPU.PC)

	a := uint16(nHi)<<8 | uint16(nLo)
	v := emu.CPU.GetReg(cpu.SP)

	emu.RAM.Set16Bit(v, a)
	emu.CPU.PC++
}

// LD SP, HL: Load stack pointer from HL
//
// Loads the value in HL into SP
func LDSPHL(emu emulator.Emulation) {
	v := emu.CPU.HL
	emu.CPU.SP = v
	emu.CPU.PC++
}

// PUSH rr: Push to stack
//
// Pushes the value of register rr to
// the stack
func PUSHrr(rr cpu.Register, emu emulator.Emulation) {
	a := emu.CPU.GetReg(cpu.SP)
	v := emu.CPU.GetReg(rr)
	emu.CPU.SP--
	emu.RAM.Set16Bit(v, a)
	emu.CPU.SP--
	emu.CPU.PC++
}

// POP rr: Pop from stack
//
// Pops from the stack into rr
func POPrr(rr cpu.Register, emu emulator.Emulation) {
	a := emu.CPU.GetReg(cpu.SP)
	v := emu.RAM.Get16Bit(a)
	emu.CPU.SP += 2
	emu.CPU.SetReg(rr, v)
	emu.CPU.PC++
}

// LD HL, SP+e: Load HL from adjusted stack pointer
//
// Loads the sum of e (next value to the instruction in memory) and SP
// into HL.
func LDHLSPpe(emu emulator.Emulation) {
	emu.CPU.PC++
	e := int8(emu.RAM.GetByte(emu.CPU.PC)) // casted so its signed
	// Here we cast into int 32 to respect e's signed value
	v := int32(emu.CPU.GetReg(cpu.SP)) + int32(e)

	emu.CPU.SetReg(cpu.HL, uint16(v)) // And now we recast it to uint16

	// Gets the lower byte of SP
	loSP := byte(emu.CPU.GetReg(cpu.SP) & cpu.LOW_MASK)
	// Gets the unsigned value of e
	loe := byte(e)

	// Sets the flag if it has a carry
	emu.CPU.SetFlag(hasCarry(loSP, loe), cpu.FlagC)
	// Sets the flag if it has a half carry
	emu.CPU.SetFlag(hasHalfCarry(loSP, loe), cpu.FlagH)
	// The rest of the flags will never be set
	emu.CPU.SetFlag(false, cpu.FlagZ)
	emu.CPU.SetFlag(false, cpu.FlagN)

	emu.CPU.PC++
}
