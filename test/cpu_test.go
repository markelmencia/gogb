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
	}
}

func TestGetter(t *testing.T) {
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

func TestSetter(t *testing.T) {
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
