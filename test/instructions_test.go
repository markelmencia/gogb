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
		RAM: getExampleRAM(),
		ROM: &[]byte{},
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
		t.Fatal("F does not match expected memory value")
	}

	if emu.CPU.PC != 2 {
		t.Fatal("Unexpected PC value")
	}
}

func TestLDrHL(t *testing.T) {
	emu := getExampleEmulation()
	emu.CPU.HL = 0x0005
	instructions.LDrHL(cpu.B, emu)
	if !(emu.CPU.Get(cpu.B) == 0xDD) {
		t.Fatal("B does not match expected memory value")
	}

	if emu.CPU.PC != 1 {
		t.Fatal("Unexpected PC value")
	}
}

func TestLDHLr(t *testing.T) {
	emu := getExampleEmulation()
	emu.CPU.HL = 0x0002
	instructions.LDHLr(cpu.E, emu)
	if !(emu.RAM.GetByte(0x0002) == 0xED) {
		t.Fatal("E does not match expected memory value")
	}

	if emu.CPU.PC != 1 {
		t.Fatal("Unexpected PC value")
	}
}

func TestLDHLn(t *testing.T) {
	emu := getExampleEmulation()
	emu.CPU.HL = 0x0006
	instructions.LDHLn(emu)
	if !(emu.RAM.GetByte(0x0006) == 0x93) {
		t.Fatal("Unexpected memory value")
	}

	if emu.CPU.PC != 2 {
		t.Fatal("Unexpected PC value")
	}
}

func TestLDaBC(t *testing.T) {
	emu := getExampleEmulation()
	emu.CPU.BC = 0x0007
	instructions.LDaBC(emu)
	if !(emu.CPU.Get(cpu.A) == 0x03) {
		t.Fatal("Unexpected memory value")
	}

	if emu.CPU.PC != 1 {
		t.Fatal("Unexpected PC value")
	}
}
