package cpu

// Masks that allow splitting combined registers
const (
	HIGH_MASK uint16 = 0xFF00
	LOW_MASK  uint16 = 0x00FF
)

type Halve byte

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

// Getters and Setters

// Getter note: In order to get the isolated value of a high 8-bit
// register an AND operation must be performed to unset the lower
// end of the 16-bit register via a mask (see HIGH_MASK and LOW_MASK).

// Then, as the value is on the higher end, it must be shifted to the
// lower end with a shift right operation in order to get the actual value.

// Setter Note: To set an 8-bit register, it is important to unset the register
// first, so the previous value does not overlap with the value to be set (v).
// That is why before the set a masking AND operation is performed in the 16-bit register,
// to only keep the halve of the register we do not want to change.

// High 8-bit register getters and setters

func (c CPU) GetA() byte {
	masked := c.AF & HIGH_MASK
	return byte(masked >> 8)
}

func (c CPU) GetB() byte {
	masked := c.BC & HIGH_MASK
	return byte(masked >> 8)
}

func (c CPU) GetD() byte {
	masked := c.DE & HIGH_MASK
	return byte(masked >> 8)
}

func (c CPU) GetH() byte {
	masked := c.HL & HIGH_MASK
	return byte(masked >> 8)
}

func (c *CPU) SetA(v byte) {
	v16 := uint16(v)
	c.AF = c.AF & LOW_MASK
	c.AF += v16 << 8
}

func (c *CPU) SetB(v byte) {
	v16 := uint16(v)
	c.BC = c.BC & LOW_MASK
	c.BC += v16 << 8
}

func (c *CPU) SetD(v byte) {
	v16 := uint16(v)
	c.DE = c.DE & LOW_MASK
	c.DE += v16 << 8
}

func (c *CPU) SetH(v byte) {
	v16 := uint16(v)
	c.HL = c.HL & LOW_MASK
	c.HL += v16 << 8
}

// Low 8-bit register getters and setters

func (c CPU) GetF() byte {
	return byte(c.AF & LOW_MASK)
}

func (c CPU) GetC() byte {
	return byte(c.BC & LOW_MASK)
}

func (c CPU) GetE() byte {
	return byte(c.DE & LOW_MASK)
}

func (c CPU) GetL() byte {
	return byte(c.HL & LOW_MASK)
}

func (c *CPU) SetF(v byte) {
	v16 := uint16(v)
	c.AF = c.AF & HIGH_MASK
	c.AF += v16
}

func (c *CPU) SetC(v byte) {
	v16 := uint16(v)
	c.BC = c.BC & HIGH_MASK
	c.BC += v16
}

func (c *CPU) SetE(v byte) {
	v16 := uint16(v)
	c.DE = c.DE & HIGH_MASK
	c.DE += v16
}

func (c *CPU) SetL(v byte) {
	v16 := uint16(v)
	c.HL = c.HL & HIGH_MASK
	c.HL += v16
}
