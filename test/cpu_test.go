package test

import (
	"testing"

	"github.com/markelmencia/gogb/cpu"
)

func getExampleCPU() cpu.CPU {
	return cpu.CPU{
		AF: 0xFACE,
		BC: 0xBEEF,
		DE: 0xFEED,
		HL: 0xDEAD,
	}
}

func TestGetter(t *testing.T) {
	cpu := getExampleCPU()
	if cpu.GetA() != 0xFA || cpu.GetF() != 0xCE {
		t.Fail()
	}

	if cpu.GetB() != 0xBE || cpu.GetC() != 0xEF {
		t.Fail()
	}

	if cpu.GetD() != 0xFE || cpu.GetE() != 0xED {
		t.Fail()
	}

	if cpu.GetH() != 0xDE || cpu.GetL() != 0xAD {
		t.Fail()
	}
}

func TestSetter(t *testing.T) {
	cpu := getExampleCPU()

	cpu.SetA(0x12)
	cpu.SetF(0x34)

	cpu.SetB(0x13)
	cpu.SetC(0x37)

	cpu.SetD(0x42)
	cpu.SetE(0x24)

	cpu.SetH(0x10)
	cpu.SetL(0x01)

	if cpu.GetA() != 0x12 || cpu.GetF() != 0x34 {
		t.Fail()
	}

	if cpu.GetB() != 0x13 || cpu.GetC() != 0x37 {
		t.Fail()
	}

	if cpu.GetD() != 0x42 || cpu.GetE() != 0x24 {
		t.Fail()
	}

	if cpu.GetH() != 0x10 || cpu.GetL() != 0x01 {
		t.Fail()
	}
}
