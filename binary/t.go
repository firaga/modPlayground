package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type T2 struct {
	A int64
	S string
	B float64
}

func main() {
	// Create a struct and write it.
	t := T2{A: 0xEEFFEEFF, S: "hello world", B: 3.14}
	buf := &bytes.Buffer{}
	err := binary.Write(buf, binary.BigEndian, t)
	if err != nil {
		panic(err)
	}
	fmt.Println(buf.Bytes())

	// Read into an empty struct.
	t = T2{}
	err = binary.Read(buf, binary.BigEndian, &t)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%x %f", t.A, t.B)
}
