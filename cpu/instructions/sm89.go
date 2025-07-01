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
