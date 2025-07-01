package main

import (
	"flag"
	"log"
	"os"

	"github.com/markelmencia/gogb/cartridge"
)

func main() {
	var header bool
	flag.BoolVar(&header, "header", false, "Prints information about the specified ROM file")
	flag.Parse()
	l := log.New(os.Stderr, "gogb: ", 0)

	if len(os.Args) < 2 {
		l.Fatal("not enough arguments: please specify the ROM path")
	}

	if header {
		if len(os.Args) < 3 {
			l.Fatal("not enough arguments: please specify the ROM path to read the cartridge header")
		}
		romPath := os.Args[2]
		cart, err := cartridge.GetCartridgeData(romPath)
		if err != nil {
			log.Fatal(err)
		}
		hdErr := cartridge.PrintHeaderData(cart)
		if hdErr != nil {
			l.Fatal(hdErr)
		}
		return // Execution ends
	}
}
