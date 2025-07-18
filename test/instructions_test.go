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
	if !(emu.CPU.GetHalve(cpu.A) == emu.CPU.GetHalve(cpu.E)) {
		t.Fatal("A does not match E")
	}

	if emu.CPU.PC != 1 {
		t.Fatal("Unexpected PC value")
	}
}

func TestLDra(t *testing.T) {
	emu := getExampleEmulation()
	instructions.LDra(cpu.F, emu)
	if !(emu.CPU.GetHalve(cpu.F) == 0x93) {
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
	if !(emu.CPU.GetHalve(cpu.B) == 0xDD) {
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
	if !(emu.CPU.GetHalve(cpu.A) == 0x03) {
		t.Fatal("Unexpected register value in A")
	}

	if emu.CPU.PC != 1 {
		t.Fatal("Unexpected PC value")
	}
}

func TestLDaDE(t *testing.T) {
	emu := getExampleEmulation()
	emu.CPU.DE = 0x0007
	instructions.LDaDE(emu)
	if !(emu.CPU.GetHalve(cpu.A) == 0x03) {
		t.Fatal("Unexpected register value in A")
	}

	if emu.CPU.PC != 1 {
		t.Fatal("Unexpected PC value")
	}
}

func TestLDBCa(t *testing.T) {
	emu := getExampleEmulation()
	emu.CPU.SetHalve(cpu.A, 0x61)
	emu.CPU.BC = 0x0001
	instructions.LDBCa(emu)
	if !(emu.RAM.GetByte(0x0001) == 0x61) {
		t.Fatal("Unexpected memory value")
	}

	if emu.CPU.PC != 1 {
		t.Fatal("Unexpected PC value")
	}
}

func TestLDECa(t *testing.T) {
	emu := getExampleEmulation()
	emu.CPU.SetHalve(cpu.A, 0x61)
	emu.CPU.DE = 0x0001
	instructions.LDDEa(emu)
	if !(emu.RAM.GetByte(0x0001) == 0x61) {
		t.Fatal("Unexpected memory value")
	}

	if emu.CPU.PC != 1 {
		t.Fatal("Unexpected PC value")
	}
}

func TestLDAnn(t *testing.T) {
	emu := getExampleEmulation()
	emu.RAM.SetByte(0x32, 0xFF93)
	instructions.LDAnn(emu)
	if emu.CPU.GetHalve(cpu.A) != 0x32 {
		t.Fatal("Unexpected register value in A")
	}

	if emu.CPU.PC != 3 {
		t.Fatal("Unexpected PC value")
	}
}

func TestLDnnA(t *testing.T) {
	emu := getExampleEmulation()
	emu.CPU.SetHalve(cpu.A, 0x32)
	instructions.LDnnA(emu)
	if emu.RAM.GetByte(0xFF93) != 0x32 {
		t.Fatal("Unexpected register value in A")
	}

	if emu.CPU.PC != 3 {
		t.Fatal("Unexpected PC value")
	}
}

func TestLDHaC(t *testing.T) {
	emu := getExampleEmulation()
	emu.RAM.SetByte(0x47, 0xFFAF)
	emu.CPU.SetHalve(cpu.C, 0xAF)
	instructions.LDHaC(emu)

	if emu.CPU.GetHalve(cpu.A) != 0x47 {
		t.Fatal("Unexpected register value in A")
	}

	if emu.CPU.PC != 1 {
		t.Fatal("Unexpected PC value")
	}
}

func TestLDHCa(t *testing.T) {
	emu := getExampleEmulation()
	emu.CPU.SetHalve(cpu.A, 0x47)
	emu.CPU.SetHalve(cpu.C, 0xAF)
	instructions.LDHCa(emu)

	if emu.RAM.GetByte(0xFFAF) != 0x47 {
		t.Fatal("Unexpected memory value")
	}

	if emu.CPU.PC != 1 {
		t.Fatal("Unexpected PC value")
	}
}

func TestLDHAn(t *testing.T) {
	emu := getExampleEmulation()
	emu.RAM.SetByte(0x47, 0xFF93)
	instructions.LDHAn(emu)

	if emu.CPU.GetHalve(cpu.A) != 0x47 {
		t.Fatal("Unexpected register value in A")
	}

	if emu.CPU.PC != 2 {
		t.Fatal("Unexpected PC value")
	}
}

func TestLDHnA(t *testing.T) {
	emu := getExampleEmulation()
	emu.CPU.SetHalve(cpu.A, 0x47)
	instructions.LDHnA(emu)

	if emu.RAM.GetByte(0xFF93) != 0x47 {
		t.Fatal("Unexpected memory value")
	}

	if emu.CPU.PC != 2 {
		t.Fatal("Unexpected PC value")
	}
}

func TestLDaHLm(t *testing.T) {
	emu := getExampleEmulation()
	emu.CPU.HL = 0x0007
	instructions.LDaHLm(emu)

	if emu.CPU.GetHalve(cpu.A) != 0x03 {
		t.Fatal("Unexpected register value in A")
	}

	if emu.CPU.HL != 0x0006 {
		t.Fatal("HL did not decrement")
	}

	if emu.CPU.PC != 1 {
		t.Fatal("Unexpected PC value")
	}
}

func TestLDHLam(t *testing.T) {
	emu := getExampleEmulation()
	emu.CPU.SetHalve(cpu.A, 0x37)
	emu.CPU.HL = 0x0007
	instructions.LDHLam(emu)

	if emu.RAM.GetByte(0x007) != 0x37 {
		t.Fatal("Unexpected memory value")
	}

	if emu.CPU.HL != 0x0006 {
		t.Fatal("HL did not decrement")
	}

	if emu.CPU.PC != 1 {
		t.Fatal("Unexpected PC value")
	}
}

func TestLDaHLp(t *testing.T) {
	emu := getExampleEmulation()
	emu.CPU.HL = 0x0007
	instructions.LDaHLp(emu)

	if emu.CPU.GetHalve(cpu.A) != 0x03 {
		t.Fatal("Unexpected register value in A")
	}

	if emu.CPU.HL != 0x0008 {
		t.Fatal("HL did not increment")
	}

	if emu.CPU.PC != 1 {
		t.Fatal("Unexpected PC value")
	}
}

func TestLDHLap(t *testing.T) {
	emu := getExampleEmulation()
	emu.CPU.SetHalve(cpu.A, 0x37)
	emu.CPU.HL = 0x0007
	instructions.LDHLap(emu)

	if emu.RAM.GetByte(0x007) != 0x37 {
		t.Fatal("Unexpected memory value")
	}

	if emu.CPU.HL != 0x0008 {
		t.Fatal("HL did not increment")
	}

	if emu.CPU.PC != 1 {
		t.Fatal("Unexpected PC value")
	}
}

func TestLDrrnn(t *testing.T) {
	emu := getExampleEmulation()
	instructions.LDrrnn(cpu.HL, emu)
	if emu.CPU.GetReg(cpu.HL) != 0xFF93 {
		t.Fatal("Unexpected register value in HL")
	}

	if emu.CPU.PC != 3 {
		t.Fatal("Unexpected PC value")
	}
}

func TestLDnnSP(t *testing.T) {
	emu := getExampleEmulation()
	emu.CPU.SetReg(cpu.SP, 0x9876)
	instructions.LDnnSP(emu)
	if emu.RAM.Get16Bit(0xFF93) != 0x9876 {
		t.Fatal("Unexpected memory value")
	}

	if emu.CPU.PC != 3 {
		t.Fatal("Unexpected PC value")
	}
}

func TestLDSPHL(t *testing.T) {
	emu := getExampleEmulation()
	instructions.LDSPHL(emu)
	if emu.CPU.GetReg(cpu.SP) != emu.CPU.GetReg(cpu.HL) {
		t.Fatal("Unexpected register value in SP")
	}

	if emu.CPU.PC != 1 {
		t.Fatal("Unexpected PC value")
	}
}

func TestPUSHrr(t *testing.T) {
	emu := getExampleEmulation()
	emu.CPU.SetReg(cpu.DE, 0x1289)
	instructions.PUSHrr(cpu.DE, emu)
	if emu.RAM.Get16Bit(0xFADE) != 0x1289 {
		t.Fatal("Unexpected memory value")
	}

	if emu.CPU.GetReg(cpu.SP) != 0xFADC {
		t.Fatal("Unexpected register value in SP")
	}

	if emu.CPU.PC != 1 {
		t.Fatal("Unexpected PC value")
	}
}

func TestPOPrr(t *testing.T) {
	emu := getExampleEmulation()
	emu.RAM.Set16Bit(0x4532, emu.CPU.GetReg(cpu.SP))
	instructions.POPrr(cpu.HL, emu)
	if emu.CPU.GetReg(cpu.HL) != 0x4532 {
		t.Fatal("Unexpected register value in HL")
	}

	if emu.CPU.GetReg(cpu.SP) != 0xFAE0 {
		t.Fatal("Unexpected register value in SP")
	}

	if emu.CPU.PC != 1 {
		t.Fatal("Unexpected PC value")
	}
}

func TestLDHLSPpe(t *testing.T) {
	emu := getExampleEmulation()
	emu.RAM.SetByte(0xFE, 0x0001) // -2
	instructions.LDHLSPpe(emu)

	if emu.CPU.GetReg(cpu.HL) != 0xFADC {
		t.Fatal("Unexpected register value in HL")
	}

	if !emu.CPU.IsFlag(cpu.FlagH) {
		t.Fatal("Unexpected flag value in H")
	}

	if !emu.CPU.IsFlag(cpu.FlagC) {
		t.Fatal("Unexpected flag value in C")
	}

	if emu.CPU.IsFlag(cpu.FlagZ) || emu.CPU.IsFlag((cpu.FlagN)) {
		t.Fatal("Unexpected flag value in Z or N")
	}

	if emu.CPU.PC != 2 {
		t.Fatal("Unexpected PC value")
	}
}

func TestADDr(t *testing.T) {
	emu := getExampleEmulation()
	vA := emu.CPU.GetHalve(cpu.A)

	instructions.ADDr(cpu.L, emu)

	if emu.CPU.GetHalve(cpu.A) != vA+emu.CPU.GetHalve(cpu.L) {
		t.Fatal("Unexpected register value in A")
	}

	if !emu.CPU.IsFlag(cpu.FlagC) {
		t.Fatal("Unexpected flag value in C")
	}

	if !emu.CPU.IsFlag(cpu.FlagH) {
		t.Fatal("Unexpected flag value in H")
	}

	if emu.CPU.PC != 1 {
		t.Fatal("Unexpected PC value")
	}
}

func TestADDHL(t *testing.T) {
	emu := getExampleEmulation()
	emu.RAM[emu.CPU.GetReg(cpu.HL)] = 0xAD
	vA := emu.CPU.GetHalve(cpu.A)

	instructions.ADDHL(emu)

	if emu.CPU.GetHalve(cpu.A) != vA+0xAD {
		t.Fatal("Unexpected register value in A")
	}

	if !emu.CPU.IsFlag(cpu.FlagC) {
		t.Fatal("Unexpected flag value in C")
	}

	if !emu.CPU.IsFlag(cpu.FlagH) {
		t.Fatal("Unexpected flag value in H")
	}

	if emu.CPU.PC != 1 {
		t.Fatal("Unexpected PC value")
	}
}

func TestADDn(t *testing.T) {
	emu := getExampleEmulation()
	emu.RAM[1] = 0xAD
	vA := emu.CPU.GetHalve(cpu.A)

	instructions.ADDn(emu)

	if emu.CPU.GetHalve(cpu.A) != vA+0xAD {
		t.Fatal("Unexpected register value in A")
	}

	if !emu.CPU.IsFlag(cpu.FlagC) {
		t.Fatal("Unexpected flag value in C")
	}

	if !emu.CPU.IsFlag(cpu.FlagH) {
		t.Fatal("Unexpected flag value in H")
	}

	if emu.CPU.PC != 2 {
		t.Fatal("Unexpected PC value")
	}
}

func TestADCr(t *testing.T) {
	emu := getExampleEmulation()
	emu.CPU.SetFlag(true, cpu.FlagC)
	emu.CPU.SetHalve(cpu.B, 0xAD)
	vA := emu.CPU.GetHalve(cpu.A)

	instructions.ADCr(cpu.B, emu)

	if emu.CPU.GetHalve(cpu.A) != vA+0xAE {
		t.Fatal("Unexpected register value in A")
	}

	if !emu.CPU.IsFlag(cpu.FlagC) {
		t.Fatal("Unexpected flag value in C")
	}

	if !emu.CPU.IsFlag(cpu.FlagH) {
		t.Fatal("Unexpected flag value in H")
	}

	if emu.CPU.PC != 1 {
		t.Fatal("Unexpected PC value")
	}
}

func TestADCHL(t *testing.T) {
	emu := getExampleEmulation()
	emu.CPU.SetFlag(true, cpu.FlagC)
	emu.RAM[emu.CPU.GetReg(cpu.HL)] = 0xAD
	vA := emu.CPU.GetHalve(cpu.A)

	instructions.ADCHL(emu)

	if emu.CPU.GetHalve(cpu.A) != vA+0xAE {
		t.Fatal("Unexpected register value in A")
	}

	if !emu.CPU.IsFlag(cpu.FlagC) {
		t.Fatal("Unexpected flag value in C")
	}

	if !emu.CPU.IsFlag(cpu.FlagH) {
		t.Fatal("Unexpected flag value in H")
	}

	if emu.CPU.PC != 1 {
		t.Fatal("Unexpected PC value")
	}
}

func TestADCn(t *testing.T) {
	emu := getExampleEmulation()
	emu.CPU.SetFlag(true, cpu.FlagC)
	emu.RAM[0x01] = 0xAD
	vA := emu.CPU.GetHalve(cpu.A)

	instructions.ADCn(emu)

	if emu.CPU.GetHalve(cpu.A) != vA+0xAE {
		t.Fatal("Unexpected register value in A")
	}

	if !emu.CPU.IsFlag(cpu.FlagC) {
		t.Fatal("Unexpected flag value in C")
	}

	if !emu.CPU.IsFlag(cpu.FlagH) {
		t.Fatal("Unexpected flag value in H")
	}

	if emu.CPU.PC != 2 {
		t.Fatal("Unexpected PC value")
	}
}

func TestSUBr(t *testing.T) {
	emu := getExampleEmulation()
	a := emu.CPU.GetHalve(cpu.A)
	instructions.SUBr(cpu.Halve(cpu.D), emu)

	if emu.CPU.GetHalve(cpu.A) != a-emu.CPU.GetHalve(cpu.D) {
		t.Fatal("Unexpected register value in A")
	}

	if !emu.CPU.IsFlag(cpu.FlagC) {
		t.Fatal("Unexpected flag value in C")
	}

	if !emu.CPU.IsFlag(cpu.FlagH) {
		t.Fatal("Unexpected flag value in H")
	}

	if emu.CPU.PC != 1 {
		t.Fatal("Unexpected PC value")
	}
}

func TestSUBHL(t *testing.T) {
	emu := getExampleEmulation()
	emu.RAM.SetByte(0xFE, 0xDEAD)
	a := emu.CPU.GetHalve(cpu.A)
	instructions.SUBHL(emu)

	if emu.CPU.GetHalve(cpu.A) != a-0xFE {
		t.Fatal("Unexpected register value in A")
	}

	if !emu.CPU.IsFlag(cpu.FlagC) {
		t.Fatal("Unexpected flag value in C")
	}

	if !emu.CPU.IsFlag(cpu.FlagH) {
		t.Fatal("Unexpected flag value in H")
	}

	if emu.CPU.PC != 1 {
		t.Fatal("Unexpected PC value")
	}
}

func TestSUBn(t *testing.T) {
	emu := getExampleEmulation()
	emu.RAM.SetByte(0xFE, 0x0001)
	a := emu.CPU.GetHalve(cpu.A)
	instructions.SUBn(emu)

	if emu.CPU.GetHalve(cpu.A) != a-0xFE {
		t.Fatal("Unexpected register value in A")
	}

	if !emu.CPU.IsFlag(cpu.FlagC) {
		t.Fatal("Unexpected flag value in C")
	}

	if !emu.CPU.IsFlag(cpu.FlagH) {
		t.Fatal("Unexpected flag value in H")
	}

	if emu.CPU.PC != 2 {
		t.Fatal("Unexpected PC value")
	}
}

func TestSBCr(t *testing.T) {
	emu := getExampleEmulation()
	emu.CPU.SetFlag(true, cpu.FlagC)
	emu.CPU.SetHalve(cpu.E, 0xFE)
	a := emu.CPU.GetHalve(cpu.A)
	instructions.SBCr(cpu.E, emu)

	if emu.CPU.GetHalve(cpu.A) != a-0xFE-1 {
		t.Fatal("Unexpected register value in A")
	}

	if !emu.CPU.IsFlag(cpu.FlagC) {
		t.Fatal("Unexpected flag value in C")
	}

	if !emu.CPU.IsFlag(cpu.FlagH) {
		t.Fatal("Unexpected flag value in H")
	}

	if emu.CPU.PC != 1 {
		t.Fatal("Unexpected PC value")
	}
}

func TestSBCHL(t *testing.T) {
	emu := getExampleEmulation()
	emu.CPU.SetFlag(true, cpu.FlagC)
	emu.RAM.SetByte(0xFE, 0xDEAD)
	a := emu.CPU.GetHalve(cpu.A)
	instructions.SBCHL(emu)

	if emu.CPU.GetHalve(cpu.A) != a-0xFE-1 {
		t.Fatal("Unexpected register value in A")
	}

	if !emu.CPU.IsFlag(cpu.FlagC) {
		t.Fatal("Unexpected flag value in C")
	}

	if !emu.CPU.IsFlag(cpu.FlagH) {
		t.Fatal("Unexpected flag value in H")
	}

	if emu.CPU.PC != 1 {
		t.Fatal("Unexpected PC value")
	}
}

func TestSBCn(t *testing.T) {
	emu := getExampleEmulation()
	emu.CPU.SetFlag(true, cpu.FlagC)
	emu.RAM.SetByte(0xFE, 0x0001)
	a := emu.CPU.GetHalve(cpu.A)
	instructions.SBCn(emu)

	if emu.CPU.GetHalve(cpu.A) != a-0xFE-1 {
		t.Fatal("Unexpected register value in A")
	}

	if !emu.CPU.IsFlag(cpu.FlagC) {
		t.Fatal("Unexpected flag value in C")
	}

	if !emu.CPU.IsFlag(cpu.FlagH) {
		t.Fatal("Unexpected flag value in H")
	}

	if emu.CPU.PC != 2 {
		t.Fatal("Unexpected PC value")
	}
}

func TestCPr(t *testing.T) {
	emu := getExampleEmulation()
	instructions.CPr(cpu.Halve(cpu.D), emu)

	if !emu.CPU.IsFlag(cpu.FlagC) {
		t.Fatal("Unexpected flag value in C")
	}

	if !emu.CPU.IsFlag(cpu.FlagH) {
		t.Fatal("Unexpected flag value in H")
	}

	if emu.CPU.PC != 1 {
		t.Fatal("Unexpected PC value")
	}
}

func TestCPHL(t *testing.T) {
	emu := getExampleEmulation()
	emu.RAM.SetByte(0xFE, 0xDEAD)
	instructions.CPHL(emu)

	if !emu.CPU.IsFlag(cpu.FlagC) {
		t.Fatal("Unexpected flag value in C")
	}

	if !emu.CPU.IsFlag(cpu.FlagH) {
		t.Fatal("Unexpected flag value in H")
	}

	if emu.CPU.PC != 1 {
		t.Fatal("Unexpected PC value")
	}
}

func TestCPn(t *testing.T) {
	emu := getExampleEmulation()
	emu.RAM.SetByte(0xFE, 0x0001)

	instructions.CPn(emu)

	if !emu.CPU.IsFlag(cpu.FlagC) {
		t.Fatal("Unexpected flag value in C")
	}

	if !emu.CPU.IsFlag(cpu.FlagH) {
		t.Fatal("Unexpected flag value in H")
	}

	if emu.CPU.PC != 2 {
		t.Fatal("Unexpected PC value")
	}
}

func TestINCr(t *testing.T) {
	emu := getExampleEmulation()
	a := emu.CPU.GetHalve(cpu.A)
	instructions.INCr(cpu.A, emu)

	if emu.CPU.GetHalve(cpu.A) != a+1 {
		t.Fatal("Unexpected value in register a")
	}

	if emu.CPU.PC != 1 {
		t.Fatal("Unexpected PC value")
	}
}

func TestINCHL(t *testing.T) {
	emu := getExampleEmulation()
	a := emu.CPU.GetReg(cpu.HL)
	vA := emu.RAM.GetByte(a)
	instructions.INCHL(emu)

	if emu.RAM.GetByte(a) != vA+1 {
		t.Fatal("Unexpected value in memory")
	}

	if emu.CPU.PC != 1 {
		t.Fatal("Unexpected PC value")
	}
}

func TestDECr(t *testing.T) {
	emu := getExampleEmulation()
	a := emu.CPU.GetHalve(cpu.A)
	instructions.DECr(cpu.A, emu)

	if emu.CPU.GetHalve(cpu.A) != a-1 {
		t.Fatal("Unexpected value in register a")
	}

	if emu.CPU.PC != 1 {
		t.Fatal("Unexpected PC value")
	}
}

func TestDECHL(t *testing.T) {
	emu := getExampleEmulation()
	a := emu.CPU.GetReg(cpu.HL)
	vA := emu.RAM.GetByte(a)
	instructions.DECHL(emu)

	if emu.RAM.GetByte(a) != vA-1 {
		t.Fatal("Unexpected value in memory")
	}

	if emu.CPU.PC != 1 {
		t.Fatal("Unexpected PC value")
	}
}

func TestANDr(t *testing.T) {
	emu := getExampleEmulation()
	a := emu.CPU.GetHalve(cpu.A)
	instructions.ANDr(cpu.B, emu)

	if emu.CPU.GetHalve(cpu.A) != a&emu.CPU.GetHalve(cpu.B) {
		t.Fatal("Unexpected value in register A")
	}

	if emu.CPU.IsFlag(cpu.FlagZ) {
		t.Fatal("Unexpected value in flag Z")
	}

	if emu.CPU.IsFlag(cpu.FlagN) {
		t.Fatal("Unexpected value in flag N")
	}

	if !emu.CPU.IsFlag(cpu.FlagH) {
		t.Fatal("Unexpected value in flag H")
	}

	if emu.CPU.IsFlag(cpu.FlagC) {
		t.Fatal("Unexpected value in flag C")
	}

	if emu.CPU.PC != 1 {
		t.Fatal("Unexpected PC value")
	}
}

func TestANDHL(t *testing.T) {
	emu := getExampleEmulation()
	a := emu.CPU.GetReg(cpu.HL)
	aV := emu.CPU.GetHalve(cpu.A)
	instructions.ANDHL(emu)

	if emu.CPU.GetHalve(cpu.A) != aV&emu.RAM.GetByte(a) {
		t.Fatal("Unexpected value in register A")
	}

	if !emu.CPU.IsFlag(cpu.FlagZ) {
		t.Fatal("Unexpected value in flag Z")
	}

	if emu.CPU.IsFlag(cpu.FlagN) {
		t.Fatal("Unexpected value in flag N")
	}

	if !emu.CPU.IsFlag(cpu.FlagH) {
		t.Fatal("Unexpected value in flag H")
	}

	if emu.CPU.IsFlag(cpu.FlagC) {
		t.Fatal("Unexpected value in flag C")
	}

	if emu.CPU.PC != 1 {
		t.Fatal("Unexpected PC value")
	}
}

func TestANDn(t *testing.T) {
	emu := getExampleEmulation()
	a := emu.CPU.GetHalve(cpu.A)
	instructions.ANDn(emu)

	if emu.CPU.GetHalve(cpu.A) != a&0x93 {
		t.Fatal("Unexpected value in register A")
	}

	if emu.CPU.IsFlag(cpu.FlagZ) {
		t.Fatal("Unexpected value in flag Z")
	}

	if emu.CPU.IsFlag(cpu.FlagN) {
		t.Fatal("Unexpected value in flag N")
	}

	if !emu.CPU.IsFlag(cpu.FlagH) {
		t.Fatal("Unexpected value in flag H")
	}

	if emu.CPU.IsFlag(cpu.FlagC) {
		t.Fatal("Unexpected value in flag C")
	}

	if emu.CPU.PC != 2 {
		t.Fatal("Unexpected PC value")
	}
}

func TestORr(t *testing.T) {
	emu := getExampleEmulation()
	a := emu.CPU.GetHalve(cpu.A)
	instructions.ORr(cpu.B, emu)

	if emu.CPU.GetHalve(cpu.A) != a|emu.CPU.GetHalve(cpu.B) {
		t.Fatal("Unexpected value in register A")
	}

	if emu.CPU.IsFlag(cpu.FlagZ) {
		t.Fatal("Unexpected value in flag Z")
	}

	if emu.CPU.IsFlag(cpu.FlagN) {
		t.Fatal("Unexpected value in flag N")
	}

	if emu.CPU.IsFlag(cpu.FlagH) {
		t.Fatal("Unexpected value in flag H")
	}

	if emu.CPU.IsFlag(cpu.FlagC) {
		t.Fatal("Unexpected value in flag C")
	}

	if emu.CPU.PC != 1 {
		t.Fatal("Unexpected PC value")
	}
}

func TestORHL(t *testing.T) {
	emu := getExampleEmulation()
	a := emu.CPU.GetReg(cpu.HL)
	aV := emu.CPU.GetHalve(cpu.A)
	instructions.ORHL(emu)

	if emu.CPU.GetHalve(cpu.A) != aV|emu.RAM.GetByte(a) {
		t.Fatal("Unexpected value in register A")
	}

	if emu.CPU.IsFlag(cpu.FlagZ) {
		t.Fatal("Unexpected value in flag Z")
	}

	if emu.CPU.IsFlag(cpu.FlagN) {
		t.Fatal("Unexpected value in flag N")
	}

	if emu.CPU.IsFlag(cpu.FlagH) {
		t.Fatal("Unexpected value in flag H")
	}

	if emu.CPU.IsFlag(cpu.FlagC) {
		t.Fatal("Unexpected value in flag C")
	}

	if emu.CPU.PC != 1 {
		t.Fatal("Unexpected PC value")
	}
}

func TestXORNn(t *testing.T) {
	emu := getExampleEmulation()
	a := emu.CPU.GetHalve(cpu.A)
	instructions.XORn(emu)

	if emu.CPU.GetHalve(cpu.A) != a^0x93 {
		t.Fatal("Unexpected value in register A")
	}

	if emu.CPU.IsFlag(cpu.FlagZ) {
		t.Fatal("Unexpected value in flag Z")
	}

	if emu.CPU.IsFlag(cpu.FlagN) {
		t.Fatal("Unexpected value in flag N")
	}

	if emu.CPU.IsFlag(cpu.FlagH) {
		t.Fatal("Unexpected value in flag H")
	}

	if emu.CPU.IsFlag(cpu.FlagC) {
		t.Fatal("Unexpected value in flag C")
	}

	if emu.CPU.PC != 2 {
		t.Fatal("Unexpected PC value")
	}
}

func TestXORr(t *testing.T) {
	emu := getExampleEmulation()
	a := emu.CPU.GetHalve(cpu.A)
	instructions.XORr(cpu.B, emu)

	if emu.CPU.GetHalve(cpu.A) != a^emu.CPU.GetHalve(cpu.B) {
		t.Fatal("Unexpected value in register A")
	}

	if emu.CPU.IsFlag(cpu.FlagZ) {
		t.Fatal("Unexpected value in flag Z")
	}

	if emu.CPU.IsFlag(cpu.FlagN) {
		t.Fatal("Unexpected value in flag N")
	}

	if emu.CPU.IsFlag(cpu.FlagH) {
		t.Fatal("Unexpected value in flag H")
	}

	if emu.CPU.IsFlag(cpu.FlagC) {
		t.Fatal("Unexpected value in flag C")
	}

	if emu.CPU.PC != 1 {
		t.Fatal("Unexpected PC value")
	}
}

func TestXORHL(t *testing.T) {
	emu := getExampleEmulation()
	a := emu.CPU.GetReg(cpu.HL)
	aV := emu.CPU.GetHalve(cpu.A)
	instructions.XORHL(emu)

	if emu.CPU.GetHalve(cpu.A) != aV^emu.RAM.GetByte(a) {
		t.Fatal("Unexpected value in register A")
	}

	if emu.CPU.IsFlag(cpu.FlagZ) {
		t.Fatal("Unexpected value in flag Z")
	}

	if emu.CPU.IsFlag(cpu.FlagN) {
		t.Fatal("Unexpected value in flag N")
	}

	if emu.CPU.IsFlag(cpu.FlagH) {
		t.Fatal("Unexpected value in flag H")
	}

	if emu.CPU.IsFlag(cpu.FlagC) {
		t.Fatal("Unexpected value in flag C")
	}

	if emu.CPU.PC != 1 {
		t.Fatal("Unexpected PC value")
	}
}

func TestXORn(t *testing.T) {
	emu := getExampleEmulation()
	a := emu.CPU.GetHalve(cpu.A)
	instructions.XORn(emu)

	if emu.CPU.GetHalve(cpu.A) != a^0x93 {
		t.Fatal("Unexpected value in register A")
	}

	if emu.CPU.IsFlag(cpu.FlagZ) {
		t.Fatal("Unexpected value in flag Z")
	}

	if emu.CPU.IsFlag(cpu.FlagN) {
		t.Fatal("Unexpected value in flag N")
	}

	if emu.CPU.IsFlag(cpu.FlagH) {
		t.Fatal("Unexpected value in flag H")
	}

	if emu.CPU.IsFlag(cpu.FlagC) {
		t.Fatal("Unexpected value in flag C")
	}

	if emu.CPU.PC != 2 {
		t.Fatal("Unexpected PC value")
	}
}

func TestCCF(t *testing.T) {
	emu := getExampleEmulation()
	c := emu.CPU.IsFlag(cpu.FlagC)
	instructions.CCF(emu)

	if emu.CPU.IsFlag(cpu.FlagC) == !c {
		t.Fatal("Unexpected value in flag C")
	}

	if emu.CPU.IsFlag(cpu.FlagN) {
		t.Fatal("Unexpected value in flag Z")
	}

	if emu.CPU.IsFlag(cpu.FlagH) {
		t.Fatal("Unexpected value in flag N")
	}

	if emu.CPU.PC != 1 {
		t.Fatal("Unexpected PC value")
	}
}

func TestSCF(t *testing.T) {
	emu := getExampleEmulation()
	instructions.SCF(emu)

	if emu.CPU.IsFlag(cpu.FlagC) != true {
		t.Fatal("Unexpected value in flag C")
	}

	if emu.CPU.IsFlag(cpu.FlagN) {
		t.Fatal("Unexpected value in flag Z")
	}

	if emu.CPU.IsFlag(cpu.FlagH) {
		t.Fatal("Unexpected value in flag N")
	}

	if emu.CPU.PC != 1 {
		t.Fatal("Unexpected PC value")
	}
}

func TestDAA(t *testing.T) {
	emu := getExampleEmulation()

	emu.CPU.SetHalve(cpu.A, 0x07)
	emu.CPU.SetHalve(cpu.B, 0x08)

	instructions.ADDr(cpu.B, emu)
	instructions.DAA(emu)

	if emu.CPU.GetHalve(cpu.A) != 0x15 {
		t.Fatal("Unexpected value in register A")
	}

	if emu.CPU.PC != 2 {
		t.Fatal("Unexpected PC value")
	}

	emu.CPU.SetHalve(cpu.A, 0x54)
	emu.CPU.SetHalve(cpu.B, 0x70)

	instructions.ADDr(cpu.B, emu)
	instructions.DAA(emu)

	if emu.CPU.GetHalve(cpu.A) != 0x24 {
		t.Fatal("Unexpected value in register A")
	}

	if !emu.CPU.IsFlag(cpu.FlagC) {
		t.Fatal("Unexpected C flag value")
	}

	if emu.CPU.PC != 4 {
		t.Fatal("Unexpected PC value")
	}

	emu.CPU.SetHalve(cpu.A, 0x05)
	emu.CPU.SetHalve(cpu.B, 0x21)

	instructions.SUBr(cpu.B, emu)
	instructions.DAA(emu)

	if emu.CPU.GetHalve(cpu.A) != 0x84 {
		t.Fatal("Unexpected value in register A")
	}

	if emu.CPU.IsFlag(cpu.FlagH) {
		t.Fatal("Unexpected H flag value")
	}

	if emu.CPU.PC != 6 {
		t.Fatal("Unexpected PC value")
	}
}

func TestCPL(t *testing.T) {
	emu := getExampleEmulation()
	a := emu.CPU.GetHalve(cpu.A)

	instructions.CPL(emu)

	if emu.CPU.GetHalve(cpu.A) != ^a {
		t.Fatal("Unexpected value in register A")
	}

	if emu.CPU.PC != 1 {
		t.Fatal("Unexpected PC value")
	}
}

func TestINCrr(t *testing.T) {
	emu := getExampleEmulation()
	bc := emu.CPU.GetReg(cpu.BC)

	instructions.INCrr(cpu.BC, emu)

	if emu.CPU.GetReg(cpu.BC) != bc+1 {
		t.Fatal("Unexpected value in register BC")
	}

	if emu.CPU.PC != 1 {
		t.Fatal("Unexpected PC value")
	}
}

func TestDECrr(t *testing.T) {
	emu := getExampleEmulation()
	bc := emu.CPU.GetReg(cpu.BC)

	instructions.DECrr(cpu.BC, emu)

	if emu.CPU.GetReg(cpu.BC) != bc-1 {
		t.Fatal("Unexpected value in register BC")
	}

	if emu.CPU.PC != 1 {
		t.Fatal("Unexpected PC value")
	}
}

func TestADDHLrr(t *testing.T) {
	emu := getExampleEmulation()
	emu.CPU.SetReg(cpu.DE, 0xEAEA)
	emu.CPU.SetReg(cpu.HL, 0x2601)

	instructions.ADDHLrr(cpu.DE, emu)

	if emu.CPU.GetReg(cpu.HL) != 0x10EB {
		t.Fatal("Unexpected value in register HL")
	}

	if !emu.CPU.IsFlag(cpu.FlagC) {
		t.Fatal("Unexpected flag C value")
	}

	if !emu.CPU.IsFlag(cpu.FlagH) {
		t.Fatal("Unexpected flag H value")
	}

	if emu.CPU.PC != 1 {
		t.Fatal("Unexpected PC value")
	}
}

func TestADDSPpe(t *testing.T) {
	emu := getExampleEmulation()
	sp := emu.CPU.GetReg(cpu.SP)
	emu.RAM.SetByte(0x32, 0x0001)

	instructions.ADDSPpe(emu)

	if emu.CPU.GetReg(cpu.SP) != sp+0x32 {
		t.Fatal("Unexpected value in register SP")
	}

	if emu.CPU.IsFlag(cpu.FlagZ) {
		t.Fatal("Unexpected flag Z value")
	}

	if emu.CPU.IsFlag(cpu.FlagN) {
		t.Fatal("Unexpected flag N value")
	}

	if !emu.CPU.IsFlag(cpu.FlagC) {
		t.Fatal("Unexpected flag C value")
	}

	if !emu.CPU.IsFlag(cpu.FlagH) {
		t.Fatal("Unexpected flag H value")
	}

	if emu.CPU.PC != 2 {
		t.Fatal("Unexpected PC value")
	}
}

func TestRLCA(t *testing.T) {
	emu := getExampleEmulation()
	emu.CPU.SetHalve(cpu.A, 0x7F)
	instructions.RLCA(emu)

	if emu.CPU.GetHalve(cpu.A) != 0xFE {
		t.Fatal("Unexpected value in register A")
	}

	if emu.CPU.IsFlag(cpu.FlagC) {
		t.Fatal("Unexpected flag C value")
	}

	// A: 0xFE
	instructions.RLCA(emu)

	if emu.CPU.GetHalve(cpu.A) != 0xFD {
		t.Fatal("Unexpected value in register A")
	}

	if !emu.CPU.IsFlag(cpu.FlagC) {
		t.Fatal("Unexpected flag C value")
	}

	if emu.CPU.PC != 2 {
		t.Fatal("Unexpected PC value")
	}
}

func TestRRCA(t *testing.T) {
	emu := getExampleEmulation()
	emu.CPU.SetHalve(cpu.A, 0xFE)
	instructions.RRCA(emu)

	if emu.CPU.GetHalve(cpu.A) != 0x7F {
		t.Fatal("Unexpected value in register A")
	}

	if emu.CPU.IsFlag(cpu.FlagC) {
		t.Fatal("Unexpected flag C value")
	}

	// A: 0x7F
	instructions.RRCA(emu)

	if emu.CPU.GetHalve(cpu.A) != 0xBF {
		t.Fatal("Unexpected value in register A")
	}

	if !emu.CPU.IsFlag(cpu.FlagC) {
		t.Fatal("Unexpected flag C value")
	}

	if emu.CPU.PC != 2 {
		t.Fatal("Unexpected PC value")
	}
}

func TestRLA(t *testing.T) {
	emu := getExampleEmulation()
	emu.CPU.SetHalve(cpu.A, 0x7F)
	instructions.RLA(emu)

	if emu.CPU.GetHalve(cpu.A) != 0xFE {
		t.Fatal("Unexpected value in register A")
	}

	if emu.CPU.IsFlag(cpu.FlagC) {
		t.Fatal("Unexpected flag C value")
	}

	// A: 0xFE
	emu.CPU.SetFlag(true, cpu.FlagC)
	instructions.RLA(emu)

	if emu.CPU.GetHalve(cpu.A) != 0xFD {
		t.Fatal("Unexpected value in register A")
	}

	if !emu.CPU.IsFlag(cpu.FlagC) {
		t.Fatal("Unexpected flag C value")
	}

	if emu.CPU.PC != 2 {
		t.Fatal("Unexpected PC value")
	}
}

func TestRRA(t *testing.T) {
	emu := getExampleEmulation()
	emu.CPU.SetHalve(cpu.A, 0xFE)
	instructions.RRA(emu)

	if emu.CPU.GetHalve(cpu.A) != 0x7F {
		t.Fatal("Unexpected value in register A")
	}

	if emu.CPU.IsFlag(cpu.FlagC) {
		t.Fatal("Unexpected flag C value")
	}

	// A: 0x7F
	instructions.RRA(emu)

	if emu.CPU.GetHalve(cpu.A) != 0x3F {
		t.Fatal("Unexpected value in register A")
	}

	if !emu.CPU.IsFlag(cpu.FlagC) {
		t.Fatal("Unexpected flag C value")
	}

	if emu.CPU.PC != 2 {
		t.Fatal("Unexpected PC value")
	}
}

func TestRLCr(t *testing.T) {
	emu := getExampleEmulation()
	emu.CPU.SetHalve(cpu.B, 0x7F)
	instructions.RLCr(cpu.B, emu)

	if emu.CPU.GetHalve(cpu.B) != 0xFE {
		t.Fatal("Unexpected value in register B")
	}

	if emu.CPU.IsFlag(cpu.FlagC) {
		t.Fatal("Unexpected flag C value")
	}

	// B: 0xFE
	instructions.RLCr(cpu.B, emu)

	if emu.CPU.GetHalve(cpu.B) != 0xFD {
		t.Fatal("Unexpected value in register B")
	}

	if !emu.CPU.IsFlag(cpu.FlagC) {
		t.Fatal("Unexpected flag C value")
	}

	if emu.CPU.PC != 2 {
		t.Fatal("Unexpected PC value")
	}
}

func TestRLCHL(t *testing.T) {
	emu := getExampleEmulation()
	emu.RAM.SetByte(0x7F, 0xDEAD)
	instructions.RLCHL(emu)

	if emu.RAM.GetByte(0xDEAD) != 0xFE {
		t.Fatal("Unexpected value in memory")
	}

	if emu.CPU.IsFlag(cpu.FlagC) {
		t.Fatal("Unexpected flag C value")
	}

	// (HL): 0xFE
	instructions.RLCHL(emu)

	if emu.RAM.GetByte(0xDEAD) != 0xFD {
		t.Fatal("Unexpected value in memory")
	}

	if !emu.CPU.IsFlag(cpu.FlagC) {
		t.Fatal("Unexpected flag C value")
	}

	if emu.CPU.PC != 2 {
		t.Fatal("Unexpected PC value")
	}
}

func TestRRCr(t *testing.T) {
	emu := getExampleEmulation()
	emu.CPU.SetHalve(cpu.B, 0xFE)
	instructions.RRCr(cpu.B, emu)

	if emu.CPU.GetHalve(cpu.B) != 0x7F {
		t.Fatal("Unexpected value in register B")
	}

	if emu.CPU.IsFlag(cpu.FlagC) {
		t.Fatal("Unexpected flag C value")
	}

	// B: 0x7F
	instructions.RRCr(cpu.B, emu)

	if emu.CPU.GetHalve(cpu.B) != 0xBF {
		t.Fatal("Unexpected value in register B")
	}

	if !emu.CPU.IsFlag(cpu.FlagC) {
		t.Fatal("Unexpected flag C value")
	}

	if emu.CPU.PC != 2 {
		t.Fatal("Unexpected PC value")
	}
}

func TestRLr(t *testing.T) {
	emu := getExampleEmulation()
	emu.CPU.SetHalve(cpu.B, 0x7F)
	instructions.RLr(cpu.B, emu)

	if emu.CPU.GetHalve(cpu.B) != 0xFE {
		t.Fatal("Unexpected value in register B")
	}

	if emu.CPU.IsFlag(cpu.FlagC) {
		t.Fatal("Unexpected flag C value")
	}

	// B: 0xFE
	emu.CPU.SetFlag(true, cpu.FlagC)
	instructions.RLr(cpu.B, emu)

	if emu.CPU.GetHalve(cpu.B) != 0xFD {
		t.Fatal("Unexpected value in register B")
	}

	if !emu.CPU.IsFlag(cpu.FlagC) {
		t.Fatal("Unexpected flag C value")
	}

	if emu.CPU.PC != 2 {
		t.Fatal("Unexpected PC value")
	}
}

func TestRLHL(t *testing.T) {
	emu := getExampleEmulation()
	emu.RAM.SetByte(0x7F, 0xDEAD)
	instructions.RLHL(emu)

	if emu.RAM.GetByte(0xDEAD) != 0xFE {
		t.Fatal("Unexpected value in memory")
	}

	if emu.CPU.IsFlag(cpu.FlagC) {
		t.Fatal("Unexpected flag C value")
	}

	// (HL): 0xFE
	emu.CPU.SetFlag(true, cpu.FlagC)
	instructions.RLHL(emu)

	if emu.RAM.GetByte(0xDEAD) != 0xFD {
		t.Fatal("Unexpected value in memory")
	}

	if !emu.CPU.IsFlag(cpu.FlagC) {
		t.Fatal("Unexpected flag C value")
	}

	if emu.CPU.PC != 2 {
		t.Fatal("Unexpected PC value")
	}
}

func TestRRr(t *testing.T) {
	emu := getExampleEmulation()
	emu.CPU.SetHalve(cpu.B, 0xFE)
	instructions.RRr(cpu.B, emu)

	if emu.CPU.GetHalve(cpu.B) != 0x7F {
		t.Fatal("Unexpected value in register B")
	}

	if emu.CPU.IsFlag(cpu.FlagC) {
		t.Fatal("Unexpected flag C value")
	}

	// B: 0x7F
	instructions.RRr(cpu.B, emu)

	if emu.CPU.GetHalve(cpu.B) != 0x3F {
		t.Fatal("Unexpected value in register B")
	}

	if !emu.CPU.IsFlag(cpu.FlagC) {
		t.Fatal("Unexpected flag C value")
	}

	if emu.CPU.PC != 2 {
		t.Fatal("Unexpected PC value")
	}
}

func TestRRHL(t *testing.T) {
	emu := getExampleEmulation()
	emu.RAM.SetByte(0xFE, 0xDEAD)
	instructions.RRHL(emu)

	if emu.RAM.GetByte(0xDEAD) != 0x7F {
		t.Fatal("Unexpected value in memory")
	}

	if emu.CPU.IsFlag(cpu.FlagC) {
		t.Fatal("Unexpected flag C value")
	}

	// (HL): 0x7F
	instructions.RRHL(emu)

	if emu.RAM.GetByte(0xDEAD) != 0x3F {
		t.Fatal("Unexpected value in memory")
	}

	if !emu.CPU.IsFlag(cpu.FlagC) {
		t.Fatal("Unexpected flag C value")
	}

	if emu.CPU.PC != 2 {
		t.Fatal("Unexpected PC value")
	}
}

func TestSLAr(t *testing.T) {
	emu := getExampleEmulation()
	emu.CPU.SetHalve(cpu.B, 0x7F)
	instructions.SLAr(cpu.B, emu)

	if emu.CPU.GetHalve(cpu.B) != 0xFE {
		t.Fatal("Unexpected value in register B")
	}

	if emu.CPU.IsFlag(cpu.FlagC) {
		t.Fatal("Unexpected flag C value")
	}

	// B: 0xFE
	emu.CPU.SetFlag(true, cpu.FlagC)
	instructions.SLAr(cpu.B, emu)

	if emu.CPU.GetHalve(cpu.B) != 0xFC {
		t.Fatal("Unexpected value in register B")
	}

	if !emu.CPU.IsFlag(cpu.FlagC) {
		t.Fatal("Unexpected flag C value")
	}

	if emu.CPU.PC != 2 {
		t.Fatal("Unexpected PC value")
	}
}

func TestSLAHL(t *testing.T) {
	emu := getExampleEmulation()
	emu.RAM.SetByte(0x7F, 0xDEAD)
	instructions.SLAHL(emu)

	if emu.RAM.GetByte(0xDEAD) != 0xFE {
		t.Fatal("Unexpected value in memory")
	}

	if emu.CPU.IsFlag(cpu.FlagC) {
		t.Fatal("Unexpected flag C value")
	}

	// B: 0xFE
	emu.CPU.SetFlag(true, cpu.FlagC)
	instructions.SLAHL(emu)

	if emu.RAM.GetByte(0xDEAD) != 0xFC {
		t.Fatal("Unexpected value in memory")
	}

	if !emu.CPU.IsFlag(cpu.FlagC) {
		t.Fatal("Unexpected flag C value")
	}

	if emu.CPU.PC != 2 {
		t.Fatal("Unexpected PC value")
	}
}

func TestSRAr(t *testing.T) {
	emu := getExampleEmulation()
	emu.CPU.SetHalve(cpu.B, 0xFE)
	instructions.SRAr(cpu.B, emu)

	if emu.CPU.GetHalve(cpu.B) != 0xFF {
		t.Fatal("Unexpected value in register B")
	}

	if emu.CPU.IsFlag(cpu.FlagC) {
		t.Fatal("Unexpected flag C value")
	}

	emu.CPU.SetHalve(cpu.B, 0x7F)
	instructions.SRAr(cpu.B, emu)

	if emu.CPU.GetHalve(cpu.B) != 0x3F {
		t.Fatal("Unexpected value in register B")
	}

	if !emu.CPU.IsFlag(cpu.FlagC) {
		t.Fatal("Unexpected flag C value")
	}

	if emu.CPU.PC != 2 {
		t.Fatal("Unexpected PC value")
	}
}

func TestSRAHL(t *testing.T) {
	emu := getExampleEmulation()
	emu.RAM.SetByte(0xFE, 0xDEAD)
	instructions.SRAHL(emu)

	if emu.RAM.GetByte(0xDEAD) != 0xFF {
		t.Fatal("Unexpected value in memory")
	}

	if emu.CPU.IsFlag(cpu.FlagC) {
		t.Fatal("Unexpected flag C value")
	}

	emu.RAM.SetByte(0x7F, 0xDEAD)
	instructions.SRAHL(emu)

	if emu.RAM.GetByte(0xDEAD) != 0x3F {
		t.Fatal("Unexpected value in memory")
	}

	if !emu.CPU.IsFlag(cpu.FlagC) {
		t.Fatal("Unexpected flag C value")
	}

	if emu.CPU.PC != 2 {
		t.Fatal("Unexpected PC value")
	}
}

func TestSWAPr(t *testing.T) {
	emu := getExampleEmulation()
	instructions.SWAPr(cpu.D, emu) // 0xFE

	if emu.CPU.GetHalve(cpu.D) != 0xEF {
		t.Fatal("Unexpected value in register D")
	}

	if emu.CPU.IsFlag(cpu.FlagZ) {
		t.Fatal("Unexpected flag Z value")
	}

	if emu.CPU.IsFlag(cpu.FlagN) {
		t.Fatal("Unexpected flag N value")
	}

	if emu.CPU.IsFlag(cpu.FlagH) {
		t.Fatal("Unexpected flag H value")
	}

	if emu.CPU.IsFlag(cpu.FlagC) {
		t.Fatal("Unexpected flag C value")
	}

	if emu.CPU.PC != 1 {
		t.Fatal("Unexpected PC value")
	}
}

func TestSWAPHL(t *testing.T) {
	emu := getExampleEmulation()
	emu.CPU.SetReg(cpu.HL, 0x0003)
	instructions.SWAPHL(emu) // 0x54

	if emu.RAM.GetByte(0x0003) != 0x45 {
		t.Fatal("Unexpected value in memory")
	}

	if emu.CPU.IsFlag(cpu.FlagZ) {
		t.Fatal("Unexpected flag Z value")
	}

	if emu.CPU.IsFlag(cpu.FlagN) {
		t.Fatal("Unexpected flag N value")
	}

	if emu.CPU.IsFlag(cpu.FlagH) {
		t.Fatal("Unexpected flag H value")
	}

	if emu.CPU.IsFlag(cpu.FlagC) {
		t.Fatal("Unexpected flag C value")
	}

	if emu.CPU.PC != 1 {
		t.Fatal("Unexpected PC value")
	}
}

func TestSRLr(t *testing.T) {
	emu := getExampleEmulation()
	emu.CPU.SetHalve(cpu.B, 0xFE)
	instructions.SRLr(cpu.B, emu)

	if emu.CPU.GetHalve(cpu.B) != 0x7F {
		t.Fatal("Unexpected value in register B")
	}

	if emu.CPU.IsFlag(cpu.FlagC) {
		t.Fatal("Unexpected flag C value")
	}

	emu.CPU.SetHalve(cpu.B, 0x7F)
	instructions.SRLr(cpu.B, emu)

	if emu.CPU.GetHalve(cpu.B) != 0x3F {
		t.Fatal("Unexpected value in register B")
	}

	if !emu.CPU.IsFlag(cpu.FlagC) {
		t.Fatal("Unexpected flag C value")
	}

	if emu.CPU.PC != 2 {
		t.Fatal("Unexpected PC value")
	}
}

func TestSRLHL(t *testing.T) {
	emu := getExampleEmulation()
	emu.RAM.SetByte(0xFE, 0xDEAD)
	instructions.SRLHL(emu)

	if emu.RAM.GetByte(0xDEAD) != 0x7F {
		t.Fatal("Unexpected value in memory")
	}

	if emu.CPU.IsFlag(cpu.FlagC) {
		t.Fatal("Unexpected flag C value")
	}

	emu.RAM.SetByte(0x7F, 0xDEAD)
	instructions.SRLHL(emu)

	if emu.RAM.GetByte(0xDEAD) != 0x3F {
		t.Fatal("Unexpected value in memory")
	}

	if !emu.CPU.IsFlag(cpu.FlagC) {
		t.Fatal("Unexpected flag C value")
	}

	if emu.CPU.PC != 2 {
		t.Fatal("Unexpected PC value")
	}
}
