package main

import (
	"fmt"
	"unsafe"
)

// 248 bit
type Memory struct {
	Age     uint8  // 8 bit
	Power   int64  // 64 bit
	Feeling bool   // 8 bit
	Health  uint32 // 32 bit
}

// 184 bit
type UpdateMemory struct {
	Power   int64  // 64 bit
	Health  uint32 // 32 bit
	Age     uint8  // 8 bit
	Feeling bool   // 8 bit
}

func main() {

	memory := &Memory{
		Age:     4,
		Power:   20,
		Feeling: true,
		Health:  30,
	}

	updateMemory := &UpdateMemory{
		Power:   20,
		Health:  30,
		Age:     4,
		Feeling: true,
	}

	fmt.Printf("Size: %d\n", unsafe.Sizeof(*memory))       // 24 bytes
	fmt.Printf("Size: %d\n", unsafe.Sizeof(*updateMemory)) // 16 bytes

}
