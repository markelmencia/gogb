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
		RAM: [65536]byte{0x00, 0x93},
		ROM: []byte{},
	}
}

func TestLDrr(t *testing.T) {
	emu := getExampleEmulation()
	instructions.LDrr(cpu.A, cpu.E, emu)
	if !(emu.CPU.Get(cpu.A) == emu.CPU.Get(cpu.E)) {
		t.Fatal("A does not match E")
	}

	if emu.CPU.PC != 1 {
		t.Fatal("Unexpected PC value")
	}
}

func TestLDra(t *testing.T) {
	emu := getExampleEmulation()
	instructions.LDra(cpu.F, emu)
	if !(emu.CPU.Get(cpu.F) == 0x93) {
		t.Fatal("A does not match expected memory value")
	}

	if emu.CPU.PC != 2 {
		t.Fatal("Unexpected PC value")
	}
}
