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
func add8(a, b byte) (byte, bool, bool) {
	// Calculates the result of the sum
	// The casting to uint16 is necessary because if the sum of two
	// bytes overflows it warps around to zero.
	result := uint16(a) + uint16(b)
	// To check the 7th bit  we can simply check if the result
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
func sub8(a, b byte) (byte, bool, bool) {
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

// 16-bit variant of add8
//
// Calculates a + b and returns it, along
// with two boolean values that are true
// if the 15th and 11th bits have a carry
// respectively, meaning:
//
// Returns:
// (result, hasCarry, hasHalfCarry)
func add16(a, b uint16) (uint16, bool, bool) {
	// Calculates the result of the sum
	// The casting to uint32 is necessary because if the sum of two
	// uint16s overflows it warps around to zero.
	result := uint32(a) + uint32(b)
	// To check the 15th bit carry we can simply check if the result
	// of the sum of these two values exceeds 0xFFFF (16-bit overflow).
	carry := result > 0xFFFF

	// To check the 11th bit carry we have to mask the lower 12 bits of
	// the operands and check if its sum exceeds 0xFFF (12-bit overflow).
	lo12a := a & 0x0FFF
	lo12b := b & 0x0FFF
	halfCarry := lo12a+lo12b > 0x0FFF

	return uint16(result), carry, halfCarry // result is casted back to uint16
}

// 16-bit variant of sub8
//
// Returns the subtraction of a - b along
// with its half and full carries (borrows).
// Meaning:
//
// Returns: (result, hasCarry, hasHalfCarry)
func sub16(a, b uint16) (uint16, bool, bool) {
	// Calculates the result
	result := a - b
	// If a < b, a carry will happen in bit 16
	carry := a < b

	// Masks the lowest 4 bits in order to
	// calculate the carry in bit 3
	lo12a := a & 0x0FFF
	lo12b := b & 0x0FFF

	// If lo4a < lo4b, a carry will happen in bit 3
	halfCarry := lo12a < lo12b

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
	_, carry, hCarry := add8(loSP, loe)

	// Sets the flag if it has a carry
	emu.CPU.SetFlag(carry, cpu.FlagC)
	// Sets the flag if it has a half carry
	emu.CPU.SetFlag(hCarry, cpu.FlagH)
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
	v, carry, hCarry := add8(emu.CPU.GetHalve(cpu.A), emu.CPU.GetHalve(r))
	emu.CPU.SetHalve(cpu.A, v)

	emu.CPU.SetFlag(v == 0, cpu.FlagZ)
	emu.CPU.SetFlag(false, cpu.FlagN)
	emu.CPU.SetFlag(carry, cpu.FlagC)
	emu.CPU.SetFlag(hCarry, cpu.FlagH)

	emu.CPU.PC++
}

// ADD (HL): Add (indirect HL)
//
// Loads into register A the value of A + the value of
// in memory in address HL.
func ADDHL(emu emulator.Emulation) {
	a := emu.CPU.GetReg(cpu.HL)
	v, carry, hCarry := add8(emu.CPU.GetHalve(cpu.A), emu.RAM.GetByte(a))
	emu.CPU.SetHalve(cpu.A, v)

	emu.CPU.SetFlag(v == 0, cpu.FlagZ)
	emu.CPU.SetFlag(false, cpu.FlagN)
	emu.CPU.SetFlag(carry, cpu.FlagC)
	emu.CPU.SetFlag(hCarry, cpu.FlagH)

	emu.CPU.PC++
}

// ADD n: Add (immediate)
//
// Loads into register A the value of A + the value of
// in memory next to the instruction.
func ADDn(emu emulator.Emulation) {
	emu.CPU.PC++
	a := emu.CPU.PC
	v, carry, hCarry := add8(emu.CPU.GetHalve(cpu.A), emu.RAM.GetByte(a))
	emu.CPU.SetHalve(cpu.A, v)

	emu.CPU.SetFlag(v == 0, cpu.FlagZ)
	emu.CPU.SetFlag(false, cpu.FlagN)
	emu.CPU.SetFlag(carry, cpu.FlagC)
	emu.CPU.SetFlag(hCarry, cpu.FlagH)

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
	v1, hasC1, hasH1 := add8(emu.CPU.GetHalve(cpu.A), emu.CPU.GetHalve(r))
	v2, hasC2, hasH2 := add8(v1, f)

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
	v1, hasC1, hasH1 := add8(emu.CPU.GetHalve(cpu.A), emu.RAM.GetByte(a))
	v2, hasC2, hasH2 := add8(v1, f)

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
	v1, hasC1, hasH1 := add8(emu.CPU.GetHalve(cpu.A), emu.RAM.GetByte(a))
	v2, hasC2, hasH2 := add8(v1, f)

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
	v, carry, hCarry := sub8(emu.CPU.GetHalve(cpu.A), emu.CPU.GetHalve(r))
	emu.CPU.SetHalve(cpu.A, v)

	emu.CPU.SetFlag(v == 0, cpu.FlagZ)
	emu.CPU.SetFlag(true, cpu.FlagN)
	emu.CPU.SetFlag(hCarry, cpu.FlagH)
	emu.CPU.SetFlag(carry, cpu.FlagC)

	emu.CPU.PC++
}

// SUB (HL): Subtract (indirect HL)
//
// Loads into register A the value of A - the value in
// memory in address HL
func SUBHL(emu emulator.Emulation) {
	a := emu.CPU.GetReg(cpu.HL)
	v, carry, hCarry := sub8(emu.CPU.GetHalve(cpu.A), emu.RAM.GetByte(a))
	emu.CPU.SetHalve(cpu.A, v)

	emu.CPU.SetFlag(v == 0, cpu.FlagZ)
	emu.CPU.SetFlag(true, cpu.FlagN)
	emu.CPU.SetFlag(hCarry, cpu.FlagH)
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
	v, carry, hCarry := sub8(emu.CPU.GetHalve(cpu.A), emu.RAM.GetByte(a))
	emu.CPU.SetHalve(cpu.A, v)

	emu.CPU.SetFlag(v == 0, cpu.FlagZ)
	emu.CPU.SetFlag(true, cpu.FlagN)
	emu.CPU.SetFlag(hCarry, cpu.FlagH)
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
	v1, hasC1, hasH1 := sub8(emu.CPU.GetHalve(cpu.A), emu.CPU.GetHalve(r))
	v2, hasC2, hasH2 := sub8(v1, f)

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
	v1, hasC1, hasH1 := sub8(emu.CPU.GetHalve(cpu.A), emu.RAM.GetByte(a))
	v2, hasC2, hasH2 := sub8(v1, f)

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
	v1, hasC1, hasH1 := sub8(emu.CPU.GetHalve(cpu.A), emu.RAM.GetByte(a))
	v2, hasC2, hasH2 := sub8(v1, f)

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
	v, carry, hCarry := sub8(emu.CPU.GetHalve(cpu.A), emu.CPU.GetHalve(r))

	emu.CPU.SetFlag(v == 0, cpu.FlagZ)
	emu.CPU.SetFlag(true, cpu.FlagN)
	emu.CPU.SetFlag(hCarry, cpu.FlagH)
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
	v, carry, hCarry := sub8(emu.CPU.GetHalve(cpu.A), emu.RAM.GetByte(a))

	emu.CPU.SetFlag(v == 0, cpu.FlagZ)
	emu.CPU.SetFlag(true, cpu.FlagN)
	emu.CPU.SetFlag(hCarry, cpu.FlagH)
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
	v, carry, hCarry := sub8(emu.CPU.GetHalve(cpu.A), emu.RAM.GetByte(a))

	emu.CPU.SetFlag(v == 0, cpu.FlagZ)
	emu.CPU.SetFlag(true, cpu.FlagN)
	emu.CPU.SetFlag(hCarry, cpu.FlagH)
	emu.CPU.SetFlag(carry, cpu.FlagC)

	emu.CPU.PC++
}

// INC r: Increment (register)
//
// Increments by 1 the value of register r.
func INCr(r cpu.Halve, emu emulator.Emulation) {
	v, _, hCarry := add8(emu.CPU.GetHalve(r), 1)
	emu.CPU.SetHalve(r, v)

	emu.CPU.SetFlag(v == 0, cpu.FlagZ)
	emu.CPU.SetFlag(false, cpu.FlagN)
	emu.CPU.SetFlag(hCarry, cpu.FlagH)
	emu.CPU.PC++
}

// INC (HL): Increment (indirect HL)
//
// Increments by 1 the value in memory in address HL.
func INCHL(emu emulator.Emulation) {
	a := emu.CPU.GetReg(cpu.HL)
	v, _, hCarry := add8(emu.RAM.GetByte(a), 1)
	emu.RAM.SetByte(v, a)

	emu.CPU.SetFlag(v == 0, cpu.FlagZ)
	emu.CPU.SetFlag(false, cpu.FlagN)
	emu.CPU.SetFlag(hCarry, cpu.FlagH)
	emu.CPU.PC++
}

// DEC r: Increment (register)
//
// Decrements by 1 the value of register r.
func DECr(r cpu.Halve, emu emulator.Emulation) {
	v, _, hCarry := sub8(emu.CPU.GetHalve(r), 1)
	emu.CPU.SetHalve(r, v)

	emu.CPU.SetFlag(v == 0, cpu.FlagZ)
	emu.CPU.SetFlag(true, cpu.FlagN)
	emu.CPU.SetFlag(hCarry, cpu.FlagH)
	emu.CPU.PC++
}

// DEC (HL): Increment (indirect HL)
//
// Decrements by 1 the value in memory in address HL.
func DECHL(emu emulator.Emulation) {
	a := emu.CPU.GetReg(cpu.HL)
	v, _, hCarry := sub8(emu.RAM.GetByte(a), 1)
	emu.RAM.SetByte(v, a)

	emu.CPU.SetFlag(v == 0, cpu.FlagZ)
	emu.CPU.SetFlag(true, cpu.FlagN)
	emu.CPU.SetFlag(hCarry, cpu.FlagH)
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

// AND (HL): Bitwise AND (indirect HL)
//
// Sets in register A the value of an AND operation
// between register A and the value in memory in
// address HL.
func ANDHL(emu emulator.Emulation) {
	// NOTE: According to Game Boy references,
	// the flag H is always set in AND operations.
	// I have not found the reason as to why this is done.
	a := emu.CPU.GetReg(cpu.HL)

	v := emu.RAM.GetByte(a) & emu.CPU.GetHalve(cpu.A)
	emu.CPU.SetHalve(cpu.A, v)

	emu.CPU.SetFlag(v == 0, cpu.FlagZ)
	emu.CPU.SetFlag(false, cpu.FlagN)
	emu.CPU.SetFlag(true, cpu.FlagH)
	emu.CPU.SetFlag(false, cpu.FlagC)
	emu.CPU.PC++
}

// AND n: Bitwise AND (immediate)
//
// Sets in register A the value of an AND operation
// between register A and the value in memory in
// the address next to the instruction..
func ANDn(emu emulator.Emulation) {
	// NOTE: According to Game Boy references,
	// the flag H is always set in AND operations.
	// I have not found the reason as to why this is done.

	emu.CPU.PC++
	a := emu.CPU.GetReg(cpu.PC)

	v := emu.RAM.GetByte(a) & emu.CPU.GetHalve(cpu.A)
	emu.CPU.SetHalve(cpu.A, v)

	emu.CPU.SetFlag(v == 0, cpu.FlagZ)
	emu.CPU.SetFlag(false, cpu.FlagN)
	emu.CPU.SetFlag(true, cpu.FlagH)
	emu.CPU.SetFlag(false, cpu.FlagC)
	emu.CPU.PC++
}

// OR r: Bitwise OR (register)
//
// Sets in register A the value of an OR operation
// between register A and r.
func ORr(r cpu.Halve, emu emulator.Emulation) {
	v := emu.CPU.GetHalve(r) | emu.CPU.GetHalve(cpu.A)
	emu.CPU.SetHalve(cpu.A, v)

	emu.CPU.SetFlag(v == 0, cpu.FlagZ)
	emu.CPU.SetFlag(false, cpu.FlagN)
	emu.CPU.SetFlag(false, cpu.FlagH)
	emu.CPU.SetFlag(false, cpu.FlagC)
	emu.CPU.PC++
}

// OR (HL): Bitwise OR (indirect HL)
//
// Sets in register A the value of an OR operation
// between register A and the value in memory in
// address HL.
func ORHL(emu emulator.Emulation) {
	a := emu.CPU.GetReg(cpu.HL)

	v := emu.RAM.GetByte(a) | emu.CPU.GetHalve(cpu.A)
	emu.CPU.SetHalve(cpu.A, v)

	emu.CPU.SetFlag(v == 0, cpu.FlagZ)
	emu.CPU.SetFlag(false, cpu.FlagN)
	emu.CPU.SetFlag(false, cpu.FlagH)
	emu.CPU.SetFlag(false, cpu.FlagC)
	emu.CPU.PC++
}

// OR n: Bitwise OR (immediate)
//
// Sets in register A the value of an OR operation
// between register A and the value in memory in
// the address next to the instruction..
func ORn(emu emulator.Emulation) {
	emu.CPU.PC++
	a := emu.CPU.GetReg(cpu.PC)

	v := emu.RAM.GetByte(a) | emu.CPU.GetHalve(cpu.A)
	emu.CPU.SetHalve(cpu.A, v)

	emu.CPU.SetFlag(v == 0, cpu.FlagZ)
	emu.CPU.SetFlag(false, cpu.FlagN)
	emu.CPU.SetFlag(false, cpu.FlagH)
	emu.CPU.SetFlag(false, cpu.FlagC)
	emu.CPU.PC++
}

// XOR r: Bitwise XOR (register)
//
// Sets in register A the value of an XOR operation
// between register A and r.
func XORr(r cpu.Halve, emu emulator.Emulation) {
	v := emu.CPU.GetHalve(r) ^ emu.CPU.GetHalve(cpu.A)
	emu.CPU.SetHalve(cpu.A, v)

	emu.CPU.SetFlag(v == 0, cpu.FlagZ)
	emu.CPU.SetFlag(false, cpu.FlagN)
	emu.CPU.SetFlag(false, cpu.FlagH)
	emu.CPU.SetFlag(false, cpu.FlagC)
	emu.CPU.PC++
}

// XOR (HL): Bitwise XOR (indirect HL)
//
// Sets in register A the value of an XOR operation
// between register A and the value in memory in
// address HL.
func XORHL(emu emulator.Emulation) {
	a := emu.CPU.GetReg(cpu.HL)

	v := emu.RAM.GetByte(a) ^ emu.CPU.GetHalve(cpu.A)
	emu.CPU.SetHalve(cpu.A, v)

	emu.CPU.SetFlag(v == 0, cpu.FlagZ)
	emu.CPU.SetFlag(false, cpu.FlagN)
	emu.CPU.SetFlag(false, cpu.FlagH)
	emu.CPU.SetFlag(false, cpu.FlagC)
	emu.CPU.PC++
}

// XOR n: Bitwise XOR (immediate)
//
// Sets in register A the value of an XOR operation
// between register A and the value in memory in
// the address next to the instruction.
func XORn(emu emulator.Emulation) {
	emu.CPU.PC++
	a := emu.CPU.GetReg(cpu.PC)

	v := emu.RAM.GetByte(a) ^ emu.CPU.GetHalve(cpu.A)
	emu.CPU.SetHalve(cpu.A, v)

	emu.CPU.SetFlag(v == 0, cpu.FlagZ)
	emu.CPU.SetFlag(false, cpu.FlagN)
	emu.CPU.SetFlag(false, cpu.FlagH)
	emu.CPU.SetFlag(false, cpu.FlagC)
	emu.CPU.PC++
}

// CCF: Complement carry flag
//
// Flips the value of the carry flag
// and clears N and H.
func CCF(emu emulator.Emulation) {
	emu.CPU.SetFlag(emu.CPU.IsFlag(cpu.FlagC), cpu.FlagC)
	emu.CPU.SetFlag(false, cpu.FlagN)
	emu.CPU.SetFlag(false, cpu.FlagH)
	emu.CPU.PC++
}

// SCF: Set carry flag
//
//	Set the carry flag and clears N and H.
func SCF(emu emulator.Emulation) {
	emu.CPU.SetFlag(true, cpu.FlagC)
	emu.CPU.SetFlag(false, cpu.FlagN)
	emu.CPU.SetFlag(false, cpu.FlagH)
	emu.CPU.PC++
}

// DAA: Decimal adjust accumulator
//
// Adjusts the value of A to turn it into a Binary Coded
// Decimal (BCD) according to the previously performed
// arithmetic operation.
func DAA(emu emulator.Emulation) {
	// NOTE: This instruction is rather unintuitive
	// if the context of BCD is unknown. For more
	// information about the workings of this
	// instruction, this article is quite helpful:
	// https://blog.ollien.com/posts/gb-daa/

	offset := byte(0)
	resultCarry := false

	// Obtains relevant flags about the previous arithmetic operation
	isSub := emu.CPU.IsFlag(cpu.FlagN)
	carry := emu.CPU.IsFlag(cpu.FlagC)
	hCarry := emu.CPU.IsFlag(cpu.FlagH)

	v := emu.CPU.GetHalve(cpu.A)
	loV := v & 0x0F

	if (!isSub && loV > 0x09) || hCarry {
		offset |= 0x06
	}

	if (!isSub && v > 0x99) || carry {
		offset |= 0x60
		resultCarry = true
	}

	var result byte
	if isSub {
		result = v - offset
	} else {
		result = v + offset
	}

	emu.CPU.SetHalve(cpu.A, result)

	emu.CPU.SetFlag(result == 0, cpu.FlagZ)
	emu.CPU.SetFlag(false, cpu.FlagH)
	emu.CPU.SetFlag(resultCarry, cpu.FlagC)
	emu.CPU.PC++
}

// CPL: Complement accumulator
//
// Complements register A (flips its bits)
// and sets flags N and H
func CPL(emu emulator.Emulation) {
	v := ^emu.CPU.GetHalve(cpu.A)

	emu.CPU.SetHalve(cpu.A, v)
	emu.CPU.SetFlag(true, cpu.FlagN)
	emu.CPU.SetFlag(true, cpu.FlagH)
	emu.CPU.PC++
}

// INC rr: Increment 16-bit register
//
// Increments by 1 the value of 16-bit register rr
func INCrr(rr cpu.Register, emu emulator.Emulation) {
	v := emu.CPU.GetReg(rr) + 1

	emu.CPU.SetReg(rr, v)
	emu.CPU.PC++
}

// DEC rr: Decrement 16-bit register
//
// Decrements by 1 the value of 16-bit register rr
func DECrr(rr cpu.Register, emu emulator.Emulation) {
	v := emu.CPU.GetReg(rr) - 1

	emu.CPU.SetReg(rr, v)
	emu.CPU.PC++
}

// ADD HL, rr: Add (16-bit register)
//
// Sets in register HL the value of HL + rr
func ADDHLrr(rr cpu.Register, emu emulator.Emulation) {
	v, carry, hCarry := add16(emu.CPU.GetReg(cpu.HL), emu.CPU.GetReg(rr))
	emu.CPU.SetReg(cpu.HL, v)

	emu.CPU.SetFlag(v == 0, cpu.FlagZ)
	emu.CPU.SetFlag(false, cpu.FlagN)
	emu.CPU.SetFlag(carry, cpu.FlagC)
	emu.CPU.SetFlag(hCarry, cpu.FlagH)

	emu.CPU.PC++
}

// ADD SP, e: Add to stack pointer (relative)
//
// Sets in register SP the value of SP + e (8 bit)
func ADDSPpe(emu emulator.Emulation) {
	emu.CPU.PC++
	e := int8(emu.RAM.GetByte(emu.CPU.PC)) // casted so its signed
	sp := emu.CPU.GetReg(cpu.SP)

	// Here we cast into int 32 to respect e's signed value
	v := int32(emu.CPU.GetReg(cpu.SP)) + int32(e)

	emu.CPU.SetReg(cpu.SP, uint16(v)) // And now we recast it to uint16

	// Gets the lower byte of SP
	loSP := byte(sp & cpu.LOW_MASK)
	// Gets the unsigned value of e
	loe := byte(e)

	// Calculates if there was a carry in the 7th and 3rd bit
	_, carry, hCarry := add8(loSP, loe)

	// Sets the flag if it has a carry
	emu.CPU.SetFlag(carry, cpu.FlagC)
	// Sets the flag if it has a half carry
	emu.CPU.SetFlag(hCarry, cpu.FlagH)
	// The rest of the flags will never be set
	emu.CPU.SetFlag(false, cpu.FlagZ)
	emu.CPU.SetFlag(false, cpu.FlagN)

	emu.CPU.PC++
}

// RLCA: Rotate left circular (accumulator)
//
// Shifts register A to the left once, and
// bit 7 is copied into the C flag and bit 0.
func RLCA(emu emulator.Emulation) {
	a := emu.CPU.GetHalve(cpu.A)
	// Moves bit 7 to the lowest position,
	// esentially rotating it
	rot := a >> 7
	v := a<<1 | rot

	emu.CPU.SetHalve(cpu.A, v)
	// If bit 7 was 1, we set flag C
	emu.CPU.SetFlag(rot > 0, cpu.FlagC)
	emu.CPU.PC++
}

// RRCA: Rotate right circular (accumulator)
//
// Shifts register A to the right once, and
// bit 0 is copied into the C flag and bit 7.
func RRCA(emu emulator.Emulation) {
	a := emu.CPU.GetHalve(cpu.A)
	// Moves bit 0 to the highest position,
	// esentially rotating it
	rot := a << 7
	v := a>>1 | rot

	emu.CPU.SetHalve(cpu.A, v)
	// If bit 7 was 1, we set flag C
	emu.CPU.SetFlag(rot > 0, cpu.FlagC)
	emu.CPU.PC++
}

// RLA: Rotate left (accumulator)
//
// Shifts register A to the left once, bit 7
// is copied into flag C, and flag C is copied
// into bit 0.
func RLA(emu emulator.Emulation) {
	a := emu.CPU.GetHalve(cpu.A)
	var rot byte
	if emu.CPU.IsFlag(cpu.FlagC) {
		rot = 1
	}

	v := a<<1 | rot

	emu.CPU.SetHalve(cpu.A, v)
	// If bit 7 was 1, we set flag C
	emu.CPU.SetFlag(a>>7 > 0, cpu.FlagC)
	emu.CPU.PC++
}

// RRA: Rotate right (accumulator)
//
// Shifts register A to the right once, bit 0
// is copied into flag C, and flag C is copied
// into bit 7.
func RRA(emu emulator.Emulation) {
	a := emu.CPU.GetHalve(cpu.A)
	var rot byte
	if emu.CPU.IsFlag(cpu.FlagC) {
		rot = 0b10000000 // 0x80
	}

	v := a>>1 | rot

	emu.CPU.SetHalve(cpu.A, v)
	// If bit 7 was 1, we set flag C
	emu.CPU.SetFlag(a<<7 > 0, cpu.FlagC)
	emu.CPU.PC++
}

// RLC r: Rotate left circular (reigster)
//
// Shifts register r to the left once, and
// bit 7 is copied into the C flag and bit 0.
func RLCr(r cpu.Halve, emu emulator.Emulation) {
	v := emu.CPU.GetHalve(r)
	// Moves bit 7 to the lowest position,
	// esentially rotating it
	rot := v >> 7
	result := v<<1 | rot

	emu.CPU.SetHalve(r, result)
	// If bit 7 was 1, we set flag C
	emu.CPU.SetFlag(rot > 0, cpu.FlagC)
	emu.CPU.PC++
}

// RLC (HL): Rotate left circular (indirect HL)
//
// Shifts the value in address HL to the left once, and
// bit 7 is copied into the C flag and bit 0.
func RLCHL(emu emulator.Emulation) {
	a := emu.CPU.GetReg(cpu.HL)
	v := emu.RAM.GetByte(a)
	// Moves bit 7 to the lowest position,
	// esentially rotating it
	rot := v >> 7
	result := v<<1 | rot

	emu.RAM.SetByte(result, a)
	// If bit 7 was 1, we set flag C
	emu.CPU.SetFlag(rot > 0, cpu.FlagC)
	emu.CPU.PC++
}

// RRC r: Rotate right circular (register)
//
// Shifts register r to the right once, and
// bit 0 is copied into the C flag and bit 7.
func RRCr(r cpu.Halve, emu emulator.Emulation) {
	v := emu.CPU.GetHalve(r)
	// Moves bit 0 to the highest position,
	// esentially rotating it
	rot := v << 7
	result := v>>1 | rot

	emu.CPU.SetHalve(r, result)
	// If bit 7 was 1, we set flag C
	emu.CPU.SetFlag(rot > 0, cpu.FlagC)
	emu.CPU.PC++
}

// RRC (HL): Rotate right circular (indirect HL)
//
// Shifts the value in address HL to the right once, and
// bit 0 is copied into the C flag and bit 7.
func RRCHL(emu emulator.Emulation) {
	a := emu.CPU.GetReg(cpu.HL)
	v := emu.RAM.GetByte(a)
	// Moves bit 0 to the highest position,
	// esentially rotating it
	rot := v << 7
	result := v>>1 | rot

	emu.RAM.SetByte(result, a)
	// If bit 7 was 1, we set flag C
	emu.CPU.SetFlag(rot > 0, cpu.FlagC)
	emu.CPU.PC++
}

// RL r: Rotate left (register)
//
// Shifts register r to the left once, bit 7
// is copied into flag C, and flag C is copied
// into bit 0.
func RLr(r cpu.Halve, emu emulator.Emulation) {
	v := emu.CPU.GetHalve(r)
	var rot byte
	if emu.CPU.IsFlag(cpu.FlagC) {
		rot = 1
	}

	result := v<<1 | rot

	emu.CPU.SetHalve(r, result)
	// If bit 7 was 1, we set flag C
	emu.CPU.SetFlag(v>>7 > 0, cpu.FlagC)
	emu.CPU.PC++
}

// RL (HL): Rotate left (indirect HL)
//
// Shifts the memory value in HL to the left once, bit 7
// is copied into flag C, and flag C is copied
// into bit 0.
func RLHL(emu emulator.Emulation) {
	a := emu.CPU.GetReg(cpu.HL)
	v := emu.RAM.GetByte(a)

	var rot byte
	if emu.CPU.IsFlag(cpu.FlagC) {
		rot = 1
	}

	result := v<<1 | rot

	emu.RAM.SetByte(result, a)
	// If bit 7 was 1, we set flag C
	emu.CPU.SetFlag(v>>7 > 0, cpu.FlagC)
	emu.CPU.PC++
}

// RR r: Rotate right (register)
//
// Shifts register r to the right once, bit 0
// is copied into flag C, and flag C is copied
// into bit 7.
func RRr(r cpu.Halve, emu emulator.Emulation) {
	v := emu.CPU.GetHalve(r)
	var rot byte
	if emu.CPU.IsFlag(cpu.FlagC) {
		rot = 0b10000000 // 0x80
	}

	result := v>>1 | rot

	emu.CPU.SetHalve(r, result)
	// If bit 7 was 1, we set flag C
	emu.CPU.SetFlag(v<<7 > 0, cpu.FlagC)
	emu.CPU.PC++
}

// RR (HL): Rotate right (indirect HL)
//
// Shifts the memory value in HL to the right once, bit 0
// is copied into flag C, and flag C is copied
// into bit 7.
func RRHL(emu emulator.Emulation) {
	a := emu.CPU.GetReg(cpu.HL)
	v := emu.RAM.GetByte(a)
	var rot byte
	if emu.CPU.IsFlag(cpu.FlagC) {
		rot = 0b10000000 // 0x80
	}

	result := v>>1 | rot

	emu.RAM.SetByte(result, a)
	// If bit 7 was 1, we set flag C
	emu.CPU.SetFlag(v<<7 > 0, cpu.FlagC)
	emu.CPU.PC++
}

// SLA r: Shift left arithmetic (register)
//
// Shifts register r to the left once and bit 7
// is copied into flag C
func SLAr(r cpu.Halve, emu emulator.Emulation) {
	v := emu.CPU.GetHalve(r)
	result := v << 1

	emu.CPU.SetHalve(r, result)
	// If bit 7 was 1, we set flag C
	emu.CPU.SetFlag(v>>7 > 0, cpu.FlagC)
	emu.CPU.PC++
}

// SLA (HL): Shift left arithmetic (indirect HL)
//
// Shifts the memory value in address HL
// to the left once and bit 7 is copied into flag C
func SLAHL(emu emulator.Emulation) {
	a := emu.CPU.GetReg(cpu.HL)
	v := emu.RAM.GetByte(a)
	result := v << 1

	emu.RAM.SetByte(result, a)
	// If bit 7 was 1, we set flag C
	emu.CPU.SetFlag(v>>7 > 0, cpu.FlagC)
	emu.CPU.PC++
}

// SRA r: Shift right arithmetic (register)
//
// Shifts register r to the right once and bit 7
// is copied into flag C.
//
// NOTE: Bit 7 is kept as-is. Its value *will* be
// rotated right, but bit 7 will remain unchanged.
func SRAr(r cpu.Halve, emu emulator.Emulation) {
	v := emu.CPU.GetHalve(r)
	bit7 := (v & 0b10000000) // Masks v to clear everything but bit 7

	// Bit 7 is added to the result so it remains unchanged
	result := v>>1 | bit7

	emu.CPU.SetHalve(r, result)
	// If bit 0 was 1, we set flag C
	emu.CPU.SetFlag(v<<7 > 0, cpu.FlagC)
	emu.CPU.PC++
}

// SRA (HL): Shift right arithmetic (indirect HL)
//
// Shifts memory value in address HL
// to the right once and bit 7 is copied into flag C.
//
// NOTE: Bit 7 is kept as-is. Its value *will* be
// rotated right, but bit 7 will remain unchanged.
func SRAHL(emu emulator.Emulation) {
	a := emu.CPU.GetReg(cpu.HL)
	v := emu.RAM.GetByte(a)
	bit7 := (v & 0b10000000) // Masks v to clear everything but bit 7

	// Bit 7 is added to the result so it remains unchanged
	result := v>>1 | bit7

	emu.RAM.SetByte(result, a)
	// If bit 0 was 1, we set flag C
	emu.CPU.SetFlag(v<<7 > 0, cpu.FlagC)
	emu.CPU.PC++
}

// SWAP r: Swap nibbles (register)
//
// Swaps the high nibble of a register
// with its low nibble.
func SWAPr(r cpu.Halve, emu emulator.Emulation) {
	v := emu.CPU.GetHalve(r)
	// Stores the high nibble
	hNib := v & 0xF0

	// Shifts the register value left to move
	// the low nibble value to the high nibble.
	// Then adds the high nibble shifted to the
	// right to perform the swap
	result := v<<4 | (hNib >> 4)

	emu.CPU.SetHalve(r, result)

	emu.CPU.SetFlag(result == 0, cpu.FlagZ)
	emu.CPU.SetFlag(false, cpu.FlagN)
	emu.CPU.SetFlag(false, cpu.FlagH)
	emu.CPU.SetFlag(false, cpu.FlagC)
	emu.CPU.PC++
}

// SWAP HL: Swap nibbles (indirect HL)
//
// Swaps the high nibble of the memory
// value in HL with its low nibble.
func SWAPHL(emu emulator.Emulation) {
	a := emu.CPU.GetReg(cpu.HL)
	v := emu.RAM.GetByte(a)
	// Stores the high nibble
	hNib := v & 0xF0

	// Shifts the register value left to move
	// the low nibble value to the high nibble.
	// Then adds the high nibble shifted to the
	// right to perform the swap
	result := v<<4 | (hNib >> 4)

	emu.RAM.SetByte(result, a)

	emu.CPU.SetFlag(result == 0, cpu.FlagZ)
	emu.CPU.SetFlag(false, cpu.FlagN)
	emu.CPU.SetFlag(false, cpu.FlagH)
	emu.CPU.SetFlag(false, cpu.FlagC)
	emu.CPU.PC++
}

// SRL r: Shift right logical (register)
//
// Shifts register r to the right logically once and bit 7
// is copied into flag C.
//
// NOTE: Bit 7 is kept as-is. Its value *will* be
// rotated right, but bit 7 will remain unchanged.
func SRLr(r cpu.Halve, emu emulator.Emulation) {
	v := emu.CPU.GetHalve(r)
	result := v >> 1

	emu.CPU.SetHalve(r, result)
	// If bit 0 was 1, we set flag C
	emu.CPU.SetFlag(v<<7 > 0, cpu.FlagC)
	emu.CPU.PC++
}

// SRL (HL): Shift right logical (indirect HL)
//
// Shifts memory value in address HL
// to the right logically once and bit 7 is copied into flag C.
//
// NOTE: Bit 7 is kept as-is. Its value *will* be
// rotated right, but bit 7 will remain unchanged.
func SRLHL(emu emulator.Emulation) {
	a := emu.CPU.GetReg(cpu.HL)
	v := emu.RAM.GetByte(a)

	result := v >> 1

	emu.RAM.SetByte(result, a)
	// If bit 0 was 1, we set flag C
	emu.CPU.SetFlag(v<<7 > 0, cpu.FlagC)
	emu.CPU.PC++
}

// BIT b, r: Test bit (register)
//
// Sets flag Z if the bit in position b of register r is zero.
func BITbr(b byte, r cpu.Halve, emu emulator.Emulation) {
	v := emu.CPU.GetHalve(r)
	// Filters out everything but bit b
	bit := v & cpu.GetBitMask(b)

	emu.CPU.SetFlag(bit > 0, cpu.FlagZ)
	emu.CPU.SetFlag(false, cpu.FlagN)
	emu.CPU.SetFlag(true, cpu.FlagH)
	emu.CPU.PC++
}

// BIT b, (HL): Test bit (indirect HL)
//
// Sets flag Z if the bit in position b of the
// memory value in address HL is zero.
func BITbHL(b byte, emu emulator.Emulation) {
	a := emu.CPU.GetReg(cpu.HL)
	v := emu.RAM.GetByte(a)
	// Filters out everything but bit b
	bit := v & cpu.GetBitMask(b)

	emu.CPU.SetFlag(bit > 0, cpu.FlagZ)
	emu.CPU.SetFlag(false, cpu.FlagN)
	emu.CPU.SetFlag(true, cpu.FlagH)
	emu.CPU.PC++
}
