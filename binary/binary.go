package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	bs := make([]byte, 8)
	binary.LittleEndian.String()
	binary.LittleEndian.PutUint64(bs, uint64(64))
	fmt.Printf("%#v\n", bs)

	i := binary.LittleEndian.Uint64(bs)
	fmt.Println(i)
}
