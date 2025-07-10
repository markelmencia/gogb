package instructions

import (
	"github.com/markelmencia/gogb/cpu"
	"github.com/markelmencia/gogb/emulator"
)

/* AUXILIARY FUNCTIONS */

// Calculates a + b and returns it, along
// with two boolean values that are true
// if the 7th and 3rd bits have a carry
// respectively, meaning:
//
// Returns:
// (result, hasCarry, hasHalfCarry)
//
// This function is useful because it
// calculates the sum of two bytes
// along with their significant carries,
// which are relevant due to flags C and H
// of the CPU.
func sum(a, b byte) (byte, bool, bool) {
	// Calculates the result of the sum
	// The casting to uint16 is necessary because if the sum of two
	// bytes overflows it resets back to zero.
	result := uint16(a) + uint16(b)
	// To check the 7th bit carry we can simply check if the result
	// of the sum of these two values exceeds 0xFF (8-bit overflow).
	carry := result > 0xFF

	// To check the 3rd bit carry we have to mask the lower 4 bits of
	// the operands and check if its sum exceeds 0xF (4-bit overflow).
	lo4a := a & 0x0F
	lo4b := b & 0x0F
	halfCarry := lo4a+lo4b > 0x0F

	return byte(result), carry, halfCarry // result is casted back to byte
}

// Returns the subtraction of a - b along
// with its half and full carries (borrows).
// Meaning:
//
// Returns: (result, hasCarry, hasHalfCarry)
func sub(a, b byte) (byte, bool, bool) {
	// Calculates the result
	result := a - b
	// If a < b, a carry will happen in bit 7
	carry := a < b

	// Masks the lowest 4 bits in order to
	// calculate the carry in bit 3
	lo4a := a & 0x0F
	lo4b := b & 0x0F

	// If lo4a < lo4b, a carry will happen in bit 3
	halfCarry := lo4a < lo4b

	return result, carry, halfCarry
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
// Loads the value in HL into SP.
func LDSPHL(emu emulator.Emulation) {
	v := emu.CPU.HL
	emu.CPU.SP = v
	emu.CPU.PC++
}

// PUSH rr: Push to stack
//
// Pushes the value of register rr to
// the stack.
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
// Pops from the stack into rr.
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

	// Calculates if there was a carry in the 7th and 3rd bit
	_, hasCarry, hasHalfCarry := sum(loSP, loe)

	// Sets the flag if it has a carry
	emu.CPU.SetFlag(hasCarry, cpu.FlagC)
	// Sets the flag if it has a half carry
	emu.CPU.SetFlag(hasHalfCarry, cpu.FlagH)
	// The rest of the flags will never be set
	emu.CPU.SetFlag(false, cpu.FlagZ)
	emu.CPU.SetFlag(false, cpu.FlagN)

	emu.CPU.PC++
}

// ADD r: Add (register)
//
// Loads into register A the value of A + the value of
// the specified register (r).
func ADDr(r cpu.Halve, emu emulator.Emulation) {
	v, hasC, hasH := sum(emu.CPU.GetHalve(cpu.A), emu.CPU.GetHalve(r))
	emu.CPU.SetHalve(cpu.A, v)

	emu.CPU.SetFlag(v == 0, cpu.FlagZ)
	emu.CPU.SetFlag(false, cpu.FlagN)
	emu.CPU.SetFlag(hasC, cpu.FlagC)
	emu.CPU.SetFlag(hasH, cpu.FlagH)

	emu.CPU.PC++
}

// ADD (HL): Add (indirect HL)
//
// Loads into register A the value of A + the value of
// in memory in address HL.
func ADDHL(emu emulator.Emulation) {
	a := emu.CPU.GetReg(cpu.HL)
	v, hasC, hasH := sum(emu.CPU.GetHalve(cpu.A), emu.RAM.GetByte(a))
	emu.CPU.SetHalve(cpu.A, v)

	emu.CPU.SetFlag(v == 0, cpu.FlagZ)
	emu.CPU.SetFlag(false, cpu.FlagN)
	emu.CPU.SetFlag(hasC, cpu.FlagC)
	emu.CPU.SetFlag(hasH, cpu.FlagH)

	emu.CPU.PC++
}

// ADD n: Add (immediate)
//
// Loads into register A the value of A + the value of
// in memory next to the instruction.
func ADDn(emu emulator.Emulation) {
	emu.CPU.PC++
	a := emu.CPU.PC
	v, hasC, hasH := sum(emu.CPU.GetHalve(cpu.A), emu.RAM.GetByte(a))
	emu.CPU.SetHalve(cpu.A, v)

	emu.CPU.SetFlag(v == 0, cpu.FlagZ)
	emu.CPU.SetFlag(false, cpu.FlagN)
	emu.CPU.SetFlag(hasC, cpu.FlagC)
	emu.CPU.SetFlag(hasH, cpu.FlagH)

	emu.CPU.PC++
}

// ADC r: Add with carry (register)
//
// Loads into register A the value of A + the value of
// register r + the carry flag.
func ADCr(r cpu.Halve, emu emulator.Emulation) {
	var f byte = 0
	if emu.CPU.IsFlag(cpu.FlagC) {
		f = 1
	}

	// First we compute the register sum, then we add 1 if the C flag
	// was set
	v1, hasC1, hasH1 := sum(emu.CPU.GetHalve(cpu.A), emu.CPU.GetHalve(r))
	v2, hasC2, hasH2 := sum(v1, f)

	emu.CPU.SetHalve(cpu.A, v2)

	emu.CPU.SetFlag(v2 == 0, cpu.FlagZ)
	emu.CPU.SetFlag(false, cpu.FlagN)
	// Since two operations were performed, we must set the carry
	// flags if a respective carry happened in any of them
	emu.CPU.SetFlag(hasC1 || hasC2, cpu.FlagC)
	emu.CPU.SetFlag(hasH1 || hasH2, cpu.FlagH)

	emu.CPU.PC++
}

// ADC (HL): Add with carry (indirect HL)
//
// Loads into register A the value of A + the value in address HL +
// the carry flag.
func ADCHL(emu emulator.Emulation) {
	var f byte = 0
	if emu.CPU.IsFlag(cpu.FlagC) {
		f = 1
	}
	a := emu.CPU.GetReg(cpu.HL)

	// First we compute the register sum, then we add 1 if the C flag
	// was set
	v1, hasC1, hasH1 := sum(emu.CPU.GetHalve(cpu.A), emu.RAM.GetByte(a))
	v2, hasC2, hasH2 := sum(v1, f)

	emu.CPU.SetHalve(cpu.A, v2)

	emu.CPU.SetFlag(v2 == 0, cpu.FlagZ)
	emu.CPU.SetFlag(false, cpu.FlagN)
	// Since two operations were performed, we must set the carry
	// flags if a respective carry happened in any of them
	emu.CPU.SetFlag(hasC1 || hasC2, cpu.FlagC)
	emu.CPU.SetFlag(hasH1 || hasH2, cpu.FlagH)

	emu.CPU.PC++
}

// ADC n: Add with carry (immediate)
//
// Loads into register A the value of A + the value in memory
// next to the instruction + the carry flag.
func ADCn(emu emulator.Emulation) {
	emu.CPU.PC++
	var f byte = 0
	if emu.CPU.IsFlag(cpu.FlagC) {
		f = 1
	}
	a := emu.CPU.PC

	// First we compute the register sum, then we add 1 if the C flag
	// was set
	v1, hasC1, hasH1 := sum(emu.CPU.GetHalve(cpu.A), emu.RAM.GetByte(a))
	v2, hasC2, hasH2 := sum(v1, f)

	emu.CPU.SetHalve(cpu.A, v2)

	emu.CPU.SetFlag(v2 == 0, cpu.FlagZ)
	emu.CPU.SetFlag(false, cpu.FlagN)
	// Since two operations were performed, we must set the carry
	// flags if a respective carry happened in any of them
	emu.CPU.SetFlag(hasC1 || hasC2, cpu.FlagC)
	emu.CPU.SetFlag(hasH1 || hasH2, cpu.FlagH)

	emu.CPU.PC++
}

// SUB r: Subtract (register)
//
// Loads into register A the value of A - the value in
// register r.
func SUBr(r cpu.Halve, emu emulator.Emulation) {
	v, carry, halfCarry := sub(emu.CPU.GetHalve(cpu.A), emu.CPU.GetHalve(r))
	emu.CPU.SetHalve(cpu.A, v)

	emu.CPU.SetFlag(v == 0, cpu.FlagZ)
	emu.CPU.SetFlag(true, cpu.FlagN)
	emu.CPU.SetFlag(halfCarry, cpu.FlagH)
	emu.CPU.SetFlag(carry, cpu.FlagC)

	emu.CPU.PC++
}

// SUB (HL): Subtract (indirect HL)
//
// Loads into register A the value of A - the value in
// memory in address HL
func SUBHL(emu emulator.Emulation) {
	a := emu.CPU.GetReg(cpu.HL)
	v, carry, halfCarry := sub(emu.CPU.GetHalve(cpu.A), emu.RAM.GetByte(a))
	emu.CPU.SetHalve(cpu.A, v)

	emu.CPU.SetFlag(v == 0, cpu.FlagZ)
	emu.CPU.SetFlag(true, cpu.FlagN)
	emu.CPU.SetFlag(halfCarry, cpu.FlagH)
	emu.CPU.SetFlag(carry, cpu.FlagC)

	emu.CPU.PC++
}

// SUB n: Subtract (immediate)
//
// Loads into register A the value of A - the value in
// memory next to the instruction
func SUBn(emu emulator.Emulation) {
	emu.CPU.PC++
	a := emu.CPU.GetReg(cpu.PC)
	v, carry, halfCarry := sub(emu.CPU.GetHalve(cpu.A), emu.RAM.GetByte(a))
	emu.CPU.SetHalve(cpu.A, v)

	emu.CPU.SetFlag(v == 0, cpu.FlagZ)
	emu.CPU.SetFlag(true, cpu.FlagN)
	emu.CPU.SetFlag(halfCarry, cpu.FlagH)
	emu.CPU.SetFlag(carry, cpu.FlagC)

	emu.CPU.PC++
}

// SBC r: Subtract with carry (register)
//
// Loads into register A the value of A - the value in
// register r - flag C
func SBCr(r cpu.Halve, emu emulator.Emulation) {
	var f byte = 0
	if emu.CPU.IsFlag(cpu.FlagC) {
		f = 1
	}

	// First we compute the register sum, then we add 1 if the C flag
	// was set
	v1, hasC1, hasH1 := sub(emu.CPU.GetHalve(cpu.A), emu.CPU.GetHalve(r))
	v2, hasC2, hasH2 := sub(v1, f)

	emu.CPU.SetHalve(cpu.A, v2)

	emu.CPU.SetFlag(v2 == 0, cpu.FlagZ)
	emu.CPU.SetFlag(true, cpu.FlagN)
	// Since two operations were performed, we must set the carry
	// flags if a respective carry happened in any of them
	emu.CPU.SetFlag(hasC1 || hasC2, cpu.FlagC)
	emu.CPU.SetFlag(hasH1 || hasH2, cpu.FlagH)

	emu.CPU.PC++
}

// SBC (HL): Subtract with carry (indirect HL)
//
// Loads into register A the value of A - the memory value
// in address HL - flag C
func SBCHL(emu emulator.Emulation) {
	var f byte = 0
	if emu.CPU.IsFlag(cpu.FlagC) {
		f = 1
	}

	a := emu.CPU.GetReg(cpu.HL)

	// First we compute the register sum, then we add 1 if the C flag
	// was set
	v1, hasC1, hasH1 := sub(emu.CPU.GetHalve(cpu.A), emu.RAM.GetByte(a))
	v2, hasC2, hasH2 := sub(v1, f)

	emu.CPU.SetHalve(cpu.A, v2)

	emu.CPU.SetFlag(v2 == 0, cpu.FlagZ)
	emu.CPU.SetFlag(true, cpu.FlagN)
	// Since two operations were performed, we must set the carry
	// flags if a respective carry happened in any of them
	emu.CPU.SetFlag(hasC1 || hasC2, cpu.FlagC)
	emu.CPU.SetFlag(hasH1 || hasH2, cpu.FlagH)

	emu.CPU.PC++
}

// SBC n: Subtract with carry (immediate)
//
// Loads into register A the value of A - the memory value next to
// the instruction - flag C.
func SBCn(emu emulator.Emulation) {
	var f byte = 0
	if emu.CPU.IsFlag(cpu.FlagC) {
		f = 1
	}

	emu.CPU.PC++
	a := emu.CPU.GetReg(cpu.PC)

	// First we compute the register sum, then we add 1 if the C flag
	// was set
	v1, hasC1, hasH1 := sub(emu.CPU.GetHalve(cpu.A), emu.RAM.GetByte(a))
	v2, hasC2, hasH2 := sub(v1, f)

	emu.CPU.SetHalve(cpu.A, v2)

	emu.CPU.SetFlag(v2 == 0, cpu.FlagZ)
	emu.CPU.SetFlag(true, cpu.FlagN)
	// Since two operations were performed, we must set the carry
	// flags if a respective carry happened in any of them
	emu.CPU.SetFlag(hasC1 || hasC2, cpu.FlagC)
	emu.CPU.SetFlag(hasH1 || hasH2, cpu.FlagH)

	emu.CPU.PC++
}

// CP r: Compare (register)
//
// Compares A with the value in register R, and
// updates the flags accordingly.
//
// Identical to SUBr, bit without modifying A.
func CPr(r cpu.Halve, emu emulator.Emulation) {
	v, carry, halfCarry := sub(emu.CPU.GetHalve(cpu.A), emu.CPU.GetHalve(r))

	emu.CPU.SetFlag(v == 0, cpu.FlagZ)
	emu.CPU.SetFlag(true, cpu.FlagN)
	emu.CPU.SetFlag(halfCarry, cpu.FlagH)
	emu.CPU.SetFlag(carry, cpu.FlagC)

	emu.CPU.PC++
}

// CP (HL): Compare (indirect HL)
//
// Compares A with the value in memory in address HL, and
// updates the flags accordingly.
//
// Identical to SUBHL, bit without modifying A.
func CPHL(emu emulator.Emulation) {
	a := emu.CPU.GetReg(cpu.HL)
	v, carry, halfCarry := sub(emu.CPU.GetHalve(cpu.A), emu.RAM.GetByte(a))

	emu.CPU.SetFlag(v == 0, cpu.FlagZ)
	emu.CPU.SetFlag(true, cpu.FlagN)
	emu.CPU.SetFlag(halfCarry, cpu.FlagH)
	emu.CPU.SetFlag(carry, cpu.FlagC)

	emu.CPU.PC++
}

// CP n: Compare (immediate)
//
// Compares A with the value in memory next to
// the instruction, and updates the flags accordingly.
//
// Identical to SUBn, bit without modifying A.
func CPn(emu emulator.Emulation) {
	emu.CPU.PC++
	a := emu.CPU.GetReg(cpu.PC)
	v, carry, halfCarry := sub(emu.CPU.GetHalve(cpu.A), emu.RAM.GetByte(a))

	emu.CPU.SetFlag(v == 0, cpu.FlagZ)
	emu.CPU.SetFlag(true, cpu.FlagN)
	emu.CPU.SetFlag(halfCarry, cpu.FlagH)
	emu.CPU.SetFlag(carry, cpu.FlagC)

	emu.CPU.PC++
}

// INC r: Increment (register)
//
// Increments by 1 the value of register r.
func INCr(r cpu.Halve, emu emulator.Emulation) {
	v, _, halfCarry := sum(emu.CPU.GetHalve(r), 1)
	emu.CPU.SetHalve(r, v)

	emu.CPU.SetFlag(v == 0, cpu.FlagZ)
	emu.CPU.SetFlag(false, cpu.FlagN)
	emu.CPU.SetFlag(halfCarry, cpu.FlagH)
	emu.CPU.PC++
}

// INC (HL): Increment (indirect HL)
//
// Increments by 1 the value in memory in address HL.
func INCHL(emu emulator.Emulation) {
	a := emu.CPU.GetReg(cpu.HL)
	v, _, halfCarry := sum(emu.RAM.GetByte(a), 1)
	emu.RAM.SetByte(v, a)

	emu.CPU.SetFlag(v == 0, cpu.FlagZ)
	emu.CPU.SetFlag(false, cpu.FlagN)
	emu.CPU.SetFlag(halfCarry, cpu.FlagH)
	emu.CPU.PC++
}

// DEC r: Increment (register)
//
// Decrements by 1 the value of register r.
func DECr(r cpu.Halve, emu emulator.Emulation) {
	v, _, halfCarry := sub(emu.CPU.GetHalve(r), 1)
	emu.CPU.SetHalve(r, v)

	emu.CPU.SetFlag(v == 0, cpu.FlagZ)
	emu.CPU.SetFlag(true, cpu.FlagN)
	emu.CPU.SetFlag(halfCarry, cpu.FlagH)
	emu.CPU.PC++
}

// DEC (HL): Increment (indirect HL)
//
// Decrements by 1 the value in memory in address HL.
func DECHL(emu emulator.Emulation) {
	a := emu.CPU.GetReg(cpu.HL)
	v, _, halfCarry := sub(emu.RAM.GetByte(a), 1)
	emu.RAM.SetByte(v, a)

	emu.CPU.SetFlag(v == 0, cpu.FlagZ)
	emu.CPU.SetFlag(true, cpu.FlagN)
	emu.CPU.SetFlag(halfCarry, cpu.FlagH)
	emu.CPU.PC++
}

// AND r: Bitwise AND (register)
//
// Sets in register A the value of an AND operation
// between register A and r.
func ANDr(r cpu.Halve, emu emulator.Emulation) {
	// NOTE: According to Game Boy references,
	// the flag H is always set in AND operations.
	// I have not found the reason as to why this is done.

	v := emu.CPU.GetHalve(r) & emu.CPU.GetHalve(cpu.A)
	emu.CPU.SetHalve(cpu.A, v)

	emu.CPU.SetFlag(v == 0, cpu.FlagZ)
	emu.CPU.SetFlag(false, cpu.FlagN)
	emu.CPU.SetFlag(true, cpu.FlagH)
	emu.CPU.SetFlag(false, cpu.FlagC)
	emu.CPU.PC++
}
