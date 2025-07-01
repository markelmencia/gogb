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
	if c.Get(cpu.A) != 0xFA || c.Get(cpu.F) != 0xCE {
		t.Fail()
	}

	if c.Get(cpu.B) != 0xBE || c.Get(cpu.C) != 0xEF {
		t.Fail()
	}

	if c.Get(cpu.D) != 0xFE || c.Get(cpu.E) != 0xED {
		t.Fail()
	}

	if c.Get(cpu.H) != 0xDE || c.Get(cpu.L) != 0xAD {
		t.Fail()
	}
}

func TestSetter(t *testing.T) {
	c := getExampleCPU()

	c.Set(cpu.A, 0x12)
	c.Set(cpu.F, 0x34)

	c.Set(cpu.B, 0x13)
	c.Set(cpu.C, 0x37)

	c.Set(cpu.D, 0x42)
	c.Set(cpu.E, 0x24)

	c.Set(cpu.H, 0x10)
	c.Set(cpu.L, 0x01)

	fmt.Printf("%X", c.AF)

	if c.Get(cpu.A) != 0x12 || c.Get(cpu.F) != 0x34 {
		t.Fail()
	}

	fmt.Printf("%X", c.AF)

	if c.Get(cpu.B) != 0x13 || c.Get(cpu.C) != 0x37 {
		t.Fail()
	}

	if c.Get(cpu.D) != 0x42 || c.Get(cpu.E) != 0x24 {
		t.Fail()
	}

	if c.Get(cpu.H) != 0x10 || c.Get(cpu.L) != 0x01 {
		t.Fail()
	}
}
