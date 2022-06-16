package main

import (
	"encoding/binary"
	"fmt"
)

//https://zhuanlan.zhihu.com/p/35326716
func main() {
	inputs := [][]byte{
		[]byte{0x81, 0x01},
		[]byte{0x7f},
		[]byte{0x03},
		[]byte{0x01},
		[]byte{0x00},
		[]byte{0x02},
		[]byte{0x04},
		[]byte{0x7e},       //63 1110100 01110100
		[]byte{0x80, 0x01}, //64 10000000 00000001 00000010000000  1000000
		{0x8d, 0x44},
	}
	for _, b := range inputs {
		x, n := binary.Uvarint(b)
		if n != len(b) {
			fmt.Println("Varint did not consume all of in")
		}
		fmt.Println(x) // -65,-64,-2,-1,0,1,2,63,64,
	}
}
