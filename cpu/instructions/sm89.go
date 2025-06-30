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
