package cartridge

import (
	"encoding/binary"
	"fmt"
	"os"
	"reflect"
)

// Returns a byte slice containing all the data
// of the cartridge defined by its path
func GetCartridgeData(dir string) ([]byte, error) {
	cartridge, rdErr := os.ReadFile(dir)
	if rdErr != nil {
		return []byte{}, rdErr
	}
	return cartridge, nil
}

// Splits the header into different fields
// to print out specific information
// about the cartridge header.
func PrintHeaderData(cart []byte) error {
	// Checks cartridge size
	if len(cart) < 0x150 {
		return fmt.Errorf("Cartridge is too small (%d bytes - Min. size: 336 bytes)",
			len(cart),
		)
	}

	// Calculates the cartridge checksums
	actualHDChecksum := GetCartHDChecksum(cart)
	actualGlobalChecksum := GetCartGlobalChecksum(cart)

	// Splits header data
	entryPoint := cart[0x100:0x104]
	logo := cart[0x104:0x134]
	title := cart[0x134:0x144]
	manufacturer := cart[0x013F:0x0143]
	cgbFlag := cart[0x143]
	licensee := cart[0x0144:0x0146]
	sgbFlag := cart[0x146]
	cartType := cart[0x0147]
	romSizeCode := cart[0x0148]
	ramSizeCode := cart[0x149]
	destination := cart[0x14A]
	licenseeOld := cart[0x14B]
	romVersion := cart[0x14C]
	hdChecksum := cart[0x14D]
	globalChecksum := cart[0x14E:0x150]

	fmt.Printf("Cartridge header data:\n\n")

	// Title
	fmt.Printf("- ROM title: %s\n", title)

	// ROM version number
	fmt.Printf("- ROM version number: 0x%X\n", romVersion)

	// Entry point instructions
	fmt.Print("- Entry point instructions:")
	// TODO: Print disassembled instructions
	fmt.Printf("0x%X ", entryPoint[0])
	fmt.Printf("0x%X ", entryPoint[1])
	fmt.Printf("0x%X ", entryPoint[2])
	fmt.Printf("0x%X\n", entryPoint[3])

	// Logo dump check
	matches := "Logo does not match expected dump"
	if reflect.DeepEqual(logo, logoBitmap) {
		matches = "Logo matches expected dump"
	}
	fmt.Printf("- ROM logo: %s\n", matches)

	// Manufacturer code
	if manufacturer[0] == 0x00 {
		manufacturer = []byte("(Old cartridge)")
	}
	fmt.Printf("- Manufacturer code: %s\n", manufacturer)

	// CGB flag
	cgbFlagInfo := ""
	switch cgbFlag {
	case 0x00:
		cgbFlagInfo = "0x00 (Old Cartridge)"
	case 0x80:
		cgbFlagInfo = "0x80 (GBC Enhanced)"
	case 0xC0:
		cgbFlagInfo = "0xC0 (GBC Only)"
	default:
		cgbFlagInfo = fmt.Sprintf("0x%X (Unknown value)", cgbFlag)
	}
	fmt.Printf("- CGB flag value: %s\n", cgbFlagInfo)

	// New licensee code
	fmt.Printf("- New licensee code: %s (%s)\n", licensee, GetNewLicenseePublisher(string(licensee)))

	// SGB flag
	sgbFlagInfo := ""
	if sgbFlag == 0x03 {
		sgbFlagInfo = "0x03 (supports GBS)"
	} else {
		sgbFlagInfo = fmt.Sprintf("0x%X (does not support GBS functions)\n", sgbFlag)
	}
	fmt.Printf("- SGB Flag Value: %s\n", sgbFlagInfo)

	// Cartridge type
	fmt.Printf("- Cartridge Type: 0x%X (%s)\n", cartType, romTypeToString[cartType])

	// ROM size
	romSize, ok := GetRomSize(romSizeCode)
	romSizeInfo := "Unknown size"
	if ok {
		romSizeInfo = fmt.Sprintf("%d KiB", romSize)
	}
	fmt.Printf("- ROM size code: 0x%X (%s)\n", romSizeCode, romSizeInfo)

	// RAM size
	ramSize, ok := GetRamSize(ramSizeCode)
	ramSizeInfo := "Unknown size"
	if ok {
		romSizeInfo = fmt.Sprintf("%d KiB", ramSize)
	}
	fmt.Printf("- RAM size code: 0x%X (%s)\n", ramSize, ramSizeInfo)

	// Destination code
	destinationInfo := ""
	switch destination {
	case 0x00:
		destinationInfo = "0x00 (Japan and possibly overseas)"
	case 0x01:
		destinationInfo = "0x01 (Overseas only)"
	default:
		destinationInfo = fmt.Sprintf("0x%X (Unknown code)", destination)
	}
	fmt.Printf("- Destination code: %s\n", destinationInfo)

	// Old licensee code
	fmt.Printf("- Old License code: 0x%X (%s)\n", licenseeOld, GetOldLicenseePublisher(licenseeOld))

	// Header checksum check
	hdChecksumInfo := ""
	if hdChecksum == actualHDChecksum {
		hdChecksumInfo = fmt.Sprintf("0x%X (Matches actual checksum)", hdChecksum)
	} else {
		hdChecksumInfo = fmt.Sprintf("0x%X (Does not match actual checksum)", hdChecksum)
	}
	fmt.Printf("- Header checksum: %s\n", hdChecksumInfo)

	// Global checksum check
	globalChecksumInfo := ""
	if binary.BigEndian.Uint16(globalChecksum) == actualGlobalChecksum {
		globalChecksumInfo = fmt.Sprintf("0x%X (Matches actual checksum)", globalChecksum)
	} else {
		globalChecksumInfo = fmt.Sprintf("0x%X (Does not match actual checksum)", globalChecksum)
	}
	fmt.Printf("- Global checksum: %s\n", globalChecksumInfo)

	return nil
}

// Calculates the cartridge's header checksum
// using its specific algoritm.
// (read https://gbdev.io/pandocs/The_Cartridge_Header.html)
// for more information.
func GetCartHDChecksum(cart []byte) byte {
	var checksum byte = 0
	for a := 0x0134; a <= 0x014C; a++ {
		checksum = checksum - cart[a] - 1
	}
	return checksum
}

// Calculates the cartridge's global checksum
// using its specific algoritm.
// (read https://gbdev.io/pandocs/The_Cartridge_Header.html)
// for more information.
func GetCartGlobalChecksum(cart []byte) uint16 {
	var checksum uint16 = 0
	for a := 0x0; a <= 0x014D; a++ {
		checksum += uint16(cart[a])
	}
	for a := 0x0150; a < len(cart); a++ {
		checksum += uint16(cart[a])
	}
	return checksum
}
