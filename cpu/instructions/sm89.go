package instructions

import (
	"github.com/markelmencia/gogb/cpu"
	"github.com/markelmencia/gogb/emulator"
)

// LD r, r': Load register (register) (8-Bit)
//
// Loads the value of r' into r
func LDrr(dst, src cpu.Halve, emu emulator.Emulation) {
	v := emu.CPU.Get(src)
	emu.CPU.Set(dst, v)
	emu.CPU.PC++
}

// LD r, n: Load register (immediate)
//
// Loads n (the value in memory next to the instruction)
// into register r.
func LDra(dst cpu.Halve, emu emulator.Emulation) {
	emu.CPU.PC++
	v := emu.RAM.GetByte(emu.CPU.PC)
	emu.CPU.Set(dst, v)
	emu.CPU.PC++
}

// LD r, (HL): Load register (indirect HL)
//
// Loads the memory value in the index inside register
// HL (16 bits) into r.
func LDrHL(dst cpu.Halve, emu emulator.Emulation) {
	a := emu.CPU.HL
	v := emu.RAM.GetByte(a)
	emu.CPU.Set(dst, v)
	emu.CPU.PC++
}

// LD (HL), r: Load from register (indirect HL)
//
// Writes the value in register r into the memory
// address specified in HL.
func LDHLr(src cpu.Halve, emu emulator.Emulation) {
	a := emu.CPU.HL
	v := emu.CPU.Get(src)
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
	emu.CPU.Set(cpu.A, v)
	emu.CPU.PC++
}

// LD A, (DE): Load accumulator (indirect DE)
//
// Loads the memory value specified in DE into A.
func LDaDE(emu emulator.Emulation) {
	a := emu.CPU.DE
	v := emu.RAM.GetByte(a)
	emu.CPU.Set(cpu.A, v)
	emu.CPU.PC++
}

// LD (BC), A: Load accumulator (indirect BC)
//
// Writes the value of register A into the
// address specified in BC.
func LDBCa(emu emulator.Emulation) {
	a := emu.CPU.BC
	v := emu.CPU.Get(cpu.A)
	emu.RAM.SetByte(v, a)
	emu.CPU.PC++
}

// LD (DE), A: Load accumulator (indirect DE)
//
// Writes the value of register A into the
// address specified in BC.
func LDDEa(emu emulator.Emulation) {
	a := emu.CPU.DE
	v := emu.CPU.Get(cpu.A)
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

	emu.CPU.Set(cpu.A, v)
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
	v := emu.CPU.Get(cpu.A)

	emu.RAM.SetByte(v, a)
	emu.CPU.PC++
}

// LDH A, (C): Load accumulator (indirect 0xFF00+C)
//
// Loads the value in memory of the address 0xFF00 + C
// into A
func LDHaC(emu emulator.Emulation) {
	a := 0xFF00 | uint16(emu.CPU.Get(cpu.C))
	v := emu.RAM.GetByte(a)
	emu.CPU.Set(cpu.A, v)
	emu.CPU.PC++
}

// LDH (C), A: Load from accumulator (indirect 0xFF00+C)
//
// Loads the value of A into the memory address 0xFF + C
func LDHCa(emu emulator.Emulation) {
	a := 0xFF00 | uint16(emu.CPU.Get(cpu.C))
	v := emu.CPU.Get(cpu.A)
	emu.RAM.SetByte(v, a)
	emu.CPU.PC++
}
