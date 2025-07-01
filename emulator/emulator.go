package emulator

import (
	"github.com/markelmencia/gogb/cpu"
	"github.com/markelmencia/gogb/ram"
)

// Represents an instance of an emulation
type Emulation struct {
	CPU *cpu.CPU
	RAM ram.RAM
	ROM []byte
}
