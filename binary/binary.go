package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	bs := make([]byte, 8)
	binary.LittleEndian.PutUint64(bs, uint64(1732))
	fmt.Printf("%#b\n", bs)

	i := binary.LittleEndian.Uint64(bs)
	fmt.Println(i)
}
