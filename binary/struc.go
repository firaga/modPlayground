package main

import (
	"bytes"
	"fmt"
	"github.com/lunixbochs/struc"
)

type Example struct {
	A int `struc:"big"`

	// B will be encoded/decoded as a 16-bit int (a "short")
	// but is stored as a native int in the struct
	B int `struc:"int16,little"`

	// the sizeof key links a buffer's size to any int field
	Size int `struc:"int8,little,sizeof=Str"`
	Str  string

	// you can get freaky if you want
	Str2 string `struc:"[5]int64"`
}

func main() {
	var buf bytes.Buffer
	t := &Example{1, 2, 0, "test", "test2"}
	err := struc.Pack(&buf, t)
	fmt.Println(err)
	fmt.Println(buf.Bytes())
	o := &Example{}
	err = struc.Unpack(&buf, o)
	fmt.Println(err)
	fmt.Println(o)
}
