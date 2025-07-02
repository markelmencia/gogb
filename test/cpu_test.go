package test

import (
	"fmt"
	"testing"

	"github.com/markelmencia/gogb/cpu"
)

func getExampleCPU() *cpu.CPU {
	return &cpu.CPU{
		AF: 0xFACE,
		BC: 0xBEEF,
		DE: 0xFEED,
		HL: 0xDEAD,
		IR: 0xACED,
		IE: 0xDACE,
		SP: 0xFADE,
		PC: 0x0000,
	}
}

func TestHalveGetter(t *testing.T) {
	c := getExampleCPU()
	if c.GetHalve(cpu.A) != 0xFA || c.GetHalve(cpu.F) != 0xCE {
		t.Fail()
	}

	if c.GetHalve(cpu.B) != 0xBE || c.GetHalve(cpu.C) != 0xEF {
		t.Fail()
	}

	if c.GetHalve(cpu.D) != 0xFE || c.GetHalve(cpu.E) != 0xED {
		t.Fail()
	}

	if c.GetHalve(cpu.H) != 0xDE || c.GetHalve(cpu.L) != 0xAD {
		t.Fail()
	}
}

func TestHalveSetter(t *testing.T) {
	c := getExampleCPU()

	c.SetHalve(cpu.A, 0x12)
	c.SetHalve(cpu.F, 0x34)

	c.SetHalve(cpu.B, 0x13)
	c.SetHalve(cpu.C, 0x37)

	c.SetHalve(cpu.D, 0x42)
	c.SetHalve(cpu.E, 0x24)

	c.SetHalve(cpu.H, 0x10)
	c.SetHalve(cpu.L, 0x01)

	if c.GetHalve(cpu.A) != 0x12 || c.GetHalve(cpu.F) != 0x34 {
		t.Fail()
	}

	fmt.Printf("%X", c.AF)

	if c.GetHalve(cpu.B) != 0x13 || c.GetHalve(cpu.C) != 0x37 {
		t.Fail()
	}

	if c.GetHalve(cpu.D) != 0x42 || c.GetHalve(cpu.E) != 0x24 {
		t.Fail()
	}

	if c.GetHalve(cpu.H) != 0x10 || c.GetHalve(cpu.L) != 0x01 {
		t.Fail()
	}
}

func TestRegGetter(t *testing.T) {
	c := getExampleCPU()
	if c.GetReg(cpu.AF) != 0xFACE {
		t.Fail()
	}

	if c.GetReg(cpu.BC) != 0xBEEF {
		t.Fail()
	}

	if c.GetReg(cpu.DE) != 0xFEED {
		t.Fail()
	}

	if c.GetReg(cpu.HL) != 0xDEAD {
		t.Fail()
	}

	if c.GetReg(cpu.IE) != 0xDACE {
		t.Fail()
	}

	if c.GetReg(cpu.IR) != 0xACED {
		t.Fail()
	}

	if c.GetReg(cpu.SP) != 0xFADE {
		t.Fail()
	}

	if c.GetReg(cpu.PC) != 0xDEED {
		t.Fail()
	}
	if c.GetReg(cpu.AF) != 0xFACE {
		t.Fail()
	}

	if c.GetReg(cpu.BC) != 0xBEEF {
		t.Fail()
	}

	if c.GetReg(cpu.DE) != 0xFEED {
		t.Fail()
	}

	if c.GetReg(cpu.HL) != 0xDEAD {
		t.Fail()
	}

	if c.GetReg(cpu.IE) != 0xDACE {
		t.Fail()
	}

	if c.GetReg(cpu.IR) != 0xACED {
		t.Fail()
	}

	if c.GetReg(cpu.SP) != 0xFADE {
		t.Fail()
	}

	if c.GetReg(cpu.PC) != 0xDEED {
		t.Fail()
	}
}

func TestRegSetter(t *testing.T) {
	c := getExampleCPU()

	c.SetReg(cpu.AF, 0x1234)
	c.SetReg(cpu.BC, 0x1337)
	c.SetReg(cpu.DE, 0x4224)
	c.SetReg(cpu.HL, 0x1001)
	c.SetReg(cpu.IE, 0x4321)
	c.SetReg(cpu.IR, 0x5678)
	c.SetReg(cpu.SP, 0x9037)
	c.SetReg(cpu.PC, 0x7132)

	if c.GetReg(cpu.AF) != 0x1234 {
		t.Fail()
	}

	if c.GetReg(cpu.BC) != 0x1337 {
		t.Fail()
	}

	if c.GetReg(cpu.DE) != 0x4224 {
		t.Fail()
	}

	if c.GetReg(cpu.HL) != 0x1001 {
		t.Fail()
	}

	if c.GetReg(cpu.IE) != 0x4321 {
		t.Fail()
	}

	if c.GetReg(cpu.IR) != 0x5678 {
		t.Fail()
	}

	if c.GetReg(cpu.SP) != 0x9037 {
		t.Fail()
	}

	if c.GetReg(cpu.PC) != 0x7132 {
		t.Fail()
	}
}

func TestFlags(t *testing.T) {
	c := getExampleCPU()

	if !c.IsFlag(cpu.FlagZ) {
		t.Fatal("Unexpected value in Z")
	}

	if !c.IsFlag(cpu.FlagN) {
		t.Fatal("Unexpected value in N")
	}

	if c.IsFlag(cpu.FlagH) {
		t.Fatal("Unexpected value in H")
	}

	if c.IsFlag(cpu.FlagC) {
		t.Fatal("Unexpected value in C")
	}

	c.SetFlag(true, cpu.FlagZ)
	c.SetFlag(false, cpu.FlagN)
	c.SetFlag(true, cpu.FlagH)
	c.SetFlag(false, cpu.FlagC)

	if c.GetHalve(cpu.F) != 0xAE {
		t.Fatal("Unexpected value in F")
	}
}
