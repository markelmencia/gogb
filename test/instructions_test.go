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

	emu.CPU.PrintStatus()

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

	emu.CPU.PrintStatus()

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
