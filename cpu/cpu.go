package cpu

import (
	"fmt"
)

// Masks that allow splitting combined registers.
const (
	HIGH_MASK uint16 = 0xFF00
	LOW_MASK  uint16 = 0x00FF
)

type Halve byte
type Register byte
type Flag byte

// Defines an enum for each halve register
// in the CPU.
const (
	A Halve = iota
	F
	B
	C
	D
	E
	H
	L
)

// Defines an enum for each register
// in the CPU.
const (
	AF Register = iota
	BC
	DE
	HL
	IR
	IE
	SP
	PC
)

// Defines an enum with each flag
const (
	FlagZ Flag = iota
	FlagN
	FlagH
	FlagC
)

// Represents a GB CPU.
// TODO: Perhaps remove exportation
type CPU struct {
	// Registers
	// Combined registers (two 8-bit registers combined)
	// NOTE: the second letter represents the lower register
	AF uint16 // Acummulator / Flag Register
	BC uint16
	DE uint16
	HL uint16
	// 16-bit registers
	IR uint16 // Instruction Register
	IE uint16 // Interrupt Enable
	SP uint16 // Stack Pointer
	PC uint16 // Program Counter
}

// Getters / Setters

// Interface that stores both
// the getter and the setter
// of a halve.
type HalveAccessor struct {
	Get func(*CPU) byte
	Set func(*CPU, byte)
}

// Interface that stores both
// the getter and the setter
// of a register.
type RegisterAccessor struct {
	Get func(*CPU) uint16
	Set func(*CPU, uint16)
}

// Getter note: In order to get the isolated value of a high 8-bit
// register an AND operation must be performed to unset the lower
// end of the 16-bit register via a mask (see HIGH_MASK and LOW_MASK).
// Then, as the value is on the higher end, it must be shifted to the
// lower end with a shift right operation in order to get the actual value.

// Setter Note: To set an 8-bit register, it is important to unset the register
// first, so the previous value does not overlap with the value to be set (v).
// That is why before the set a masking AND operation is performed in the 16-bit register,
// to only keep the halve of the register we do not want to change.

// Map that contains a getter/setter pair
// for each Halve register.
var halveToAccessor = map[Halve]HalveAccessor{
	A: {
		Get: func(c *CPU) byte {
			return byte(c.AF >> 8)
		},
		Set: func(c *CPU, v byte) {
			c.AF = (c.AF & LOW_MASK) | (uint16(v) << 8)
		},
	},
	F: {
		Get: func(c *CPU) byte {
			return byte(c.AF)
		},
		Set: func(c *CPU, v byte) {
			v16 := uint16(v)
			c.AF = c.AF & HIGH_MASK
			c.AF |= v16
		},
	},
	B: {
		Get: func(c *CPU) byte {
			return byte(c.BC >> 8)
		},
		Set: func(c *CPU, v byte) {
			c.BC = (c.BC & LOW_MASK) | (uint16(v) << 8)
		},
	},
	C: {
		Get: func(c *CPU) byte {
			return byte(c.BC)
		},
		Set: func(c *CPU, v byte) {
			v16 := uint16(v)
			c.BC = c.BC & HIGH_MASK
			c.BC |= v16
		},
	},
	D: {
		Get: func(c *CPU) byte {
			return byte(c.DE >> 8)
		},
		Set: func(c *CPU, v byte) {
			c.DE = (c.DE & LOW_MASK) | (uint16(v) << 8)
		},
	},
	E: {
		Get: func(c *CPU) byte {
			return byte(c.DE)
		},
		Set: func(c *CPU, v byte) {
			v16 := uint16(v)
			c.DE = c.DE & HIGH_MASK
			c.DE |= v16
		},
	},
	H: {
		Get: func(c *CPU) byte {
			return byte(c.HL >> 8)
		},
		Set: func(c *CPU, v byte) {
			c.HL = (c.HL & LOW_MASK) | (uint16(v) << 8)
		},
	},
	L: {
		Get: func(c *CPU) byte {
			return byte(c.HL)
		},
		Set: func(c *CPU, v byte) {
			v16 := uint16(v)
			c.HL = c.HL & HIGH_MASK
			c.HL |= v16
		},
	},
}

// Gets the appropiate register getter
// and returns the value of the register.
func (c *CPU) GetHalve(h Halve) byte {
	return halveToAccessor[h].Get(c)
}

// Gets the appropiate register setter
// and sets v into the register.
func (c *CPU) SetHalve(h Halve, v byte) {
	halveToAccessor[h].Set(c, v)
}

// Map that contains a getter/setter pair
// for each register.
var registerToAccessor = map[Register]RegisterAccessor{
	AF: {
		Get: func(c *CPU) uint16 {
			return c.AF
		},
		Set: func(c *CPU, v uint16) {
			c.AF = v
		},
	},
	BC: {
		Get: func(c *CPU) uint16 {
			return c.BC
		},
		Set: func(c *CPU, v uint16) {
			c.BC = v
		},
	},
	DE: {
		Get: func(c *CPU) uint16 {
			return c.DE
		},
		Set: func(c *CPU, v uint16) {
			c.DE = v
		},
	},
	HL: {
		Get: func(c *CPU) uint16 {
			return c.HL
		},
		Set: func(c *CPU, v uint16) {
			c.HL = v
		},
	},
	IR: {
		Get: func(c *CPU) uint16 {
			return c.IR
		},
		Set: func(c *CPU, v uint16) {
			c.IR = v
		},
	},
	IE: {
		Get: func(c *CPU) uint16 {
			return c.IE
		},
		Set: func(c *CPU, v uint16) {
			c.IE = v
		},
	},
	SP: {
		Get: func(c *CPU) uint16 {
			return c.SP
		},
		Set: func(c *CPU, v uint16) {
			c.SP = v
		},
	},
	PC: {
		Get: func(c *CPU) uint16 {
			return c.PC
		},
		Set: func(c *CPU, v uint16) {
			c.PC = v
		},
	},
}

// Returns the appropiate mask
// to obtain the corresponding
// byte to the desired flag.
var flagToMask = map[Flag]byte{
	FlagZ: 0b10000000,
	FlagN: 0b01000000,
	FlagH: 0b00100000,
	FlagC: 0b00010000,
}

// Flags

// Obtains the bit that corresponds to flag f, and
// returns true if it was set. If it wasn't, it returns
// false.
func (c *CPU) IsFlag(f Flag) bool {
	mask := flagToMask[f]

	// We mask the value of F so that all the bits in the register
	// are set to 0 except the bit that corresponds to the flag
	masked := c.GetHalve(F) & mask

	// If masked is 0, it means that the flag bit was also 0
	if masked == 0 {
		return false
	} else { // If masked is non-0, it means that the flag bit was 1
		return true
	}
}

// Sets the desired flag f to 1 if s is true
// or 0 if s is false.
func (c *CPU) SetFlag(s bool, f Flag) {
	mask := flagToMask[f]
	var v byte
	if s {
		// Sets the flag bit to 1 and leaves the rest untouched
		v = c.GetHalve(F) | mask
	} else {
		// Sets the flag bit to 0 and leaves the rest untouched
		v = c.GetHalve(F) & ^mask
	}
	c.SetHalve(F, v)
}

// Gets the appropiate register getter
// and returns the value of the register.
func (c *CPU) GetReg(r Register) uint16 {
	return registerToAccessor[r].Get(c)
}

// Gets the appropiate register setter
// and sets v into the register.
func (c *CPU) SetReg(r Register, v uint16) {
	registerToAccessor[r].Set(c, v)
}

// Prints the values of all registers
// at the current CPU state.
func (c CPU) PrintStatus() {
	fmt.Printf("CPU Status:\n\n")

	fmt.Printf("Register values:\n")
	fmt.Printf("AF: 0x%X\n", c.AF)
	fmt.Printf("BC: 0x%X\n", c.BC)
	fmt.Printf("DE: 0x%X\n", c.DE)
	fmt.Printf("HL: 0x%X\n", c.HL)
	fmt.Printf("IR: 0x%X\n", c.IR)
	fmt.Printf("IE: 0x%X\n", c.IE)
	fmt.Printf("SP: 0x%X\n", c.SP)
	fmt.Printf("PC: 0x%X\n\n", c.PC)

	fmt.Printf("Flags:\n")
	fmt.Printf("Z: %t\n", c.IsFlag(FlagZ))
	fmt.Printf("N: %t\n", c.IsFlag(FlagN))
	fmt.Printf("H: %t\n", c.IsFlag(FlagH))
	fmt.Printf("C: %t\n", c.IsFlag(FlagC))
}
