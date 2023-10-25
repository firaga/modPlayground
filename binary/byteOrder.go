package main

import (
	"encoding/binary"
	"fmt"
)

//https://pkg.go.dev/encoding/binary#ByteOrder
//byteOrder 是 binary.LittleEndian binary.BigEndian实现的接口

func main() {
	b := []byte{0xe8, 0x03, 0xd0, 0x07}
	x1 := binary.LittleEndian.Uint16(b[0:])
	x2 := binary.LittleEndian.Uint16(b[2:])
	x3 := binary.LittleEndian.Uint32(b[0:])
	x4 := binary.BigEndian.Uint32(b[0:])
	fmt.Printf("%#08x %#08x\n", b[0:], b[2:])
	fmt.Printf("%#04x %#04x\n", x1, x2)
	fmt.Printf("%#08x %#08x\n", x3, x4)
}
