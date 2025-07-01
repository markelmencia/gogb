package test

import (
	"testing"

	"github.com/markelmencia/gogb/cpu"
	"github.com/markelmencia/gogb/cpu/instructions"
	"github.com/markelmencia/gogb/emulator"
)

func getExampleEmulation() emulator.Emulation {
	return emulator.Emulation{
		CPU: getExampleCPU(),
		RAM: [65536]byte{},
		ROM: []byte{},
	}
}

func TestLDrr(t *testing.T) {
	emu := getExampleEmulation()
	instructions.LDrr(cpu.A, cpu.E, emu)
	emu.CPU.PrintStatus()
	if !(emu.CPU.Get(cpu.A) == emu.CPU.Get(cpu.E)) {
		t.Fatal("A does not match E")
	}

	if emu.CPU.PC != 1 {
		t.Fatal("PC has not incremented")
	}
	instructions.LDrr(cpu.A, cpu.D, emu)
}
