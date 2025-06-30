package ram

type RAM [65536]byte

// Returns the byte stored in the
// specified memory address.
func (r RAM) GetByte(a uint16) byte {
	return r[a]
}
