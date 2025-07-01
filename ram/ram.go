package ram

type RAM [65536]byte

// Returns the byte stored in the
// specified memory address.
func (r *RAM) GetByte(a uint16) byte {
	return r[a]
}

// Sets v into the address specified in a
func (r *RAM) SetByte(v byte, a uint16) {
	r[a] = v
}

// Returns the 16-bit value stored in a
// and a + 1
func (r *RAM) Get16Bit(a uint16) uint16 {
	return uint16(r[a+1])<<8 | uint16(r[a])
}

// Sets the value v into a and a + 1
func (r *RAM) Set16Bit(v uint16, a uint16) {
	r[a] = byte(v)
	r[a+1] = byte(v >> 8)
}
