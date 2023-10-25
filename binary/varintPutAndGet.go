package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	num := int64(1734)
	bs := make([]byte, binary.MaxVarintLen64)
	binary.PutVarint(bs, num)
	fmt.Printf("%#b\n", bs)
	i, n := binary.Varint(bs)
	fmt.Printf("value = %v,length=%v\n", i, n)
}
