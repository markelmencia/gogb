package test

import (
	"testing"

	"github.com/markelmencia/gogb/ram"
)

func getExampleRAM() *ram.RAM {
	return &ram.RAM{0x00, 0x93, 0xFF, 0x54, 0xAE, 0xDD, 0x12, 0x03, 0x41}
}

func TestGetterSetter(t *testing.T) {
	r := getExampleRAM()
	if r.GetByte(2) != 0xFF {
		t.Fatal("Unexpected value in getter")
	}

	r.SetByte(0xFE, 3)
	if r.GetByte(3) != 0xFE {
		t.Fatal("Unexpected value after set")
	}

	if r.Get16Bit(1) != 0xFF93 {
		t.Fatal("Unexpected value in getter")
	}

	r.Set16Bit(0x3412, 3)
	if r.Get16Bit(3) != 0x3412 {
		t.Fatal("Unexpected value in getter")
	}

}
