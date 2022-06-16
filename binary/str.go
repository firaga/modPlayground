package main

import (
	"bytes"
	"fmt"
	"github.com/lunixbochs/struc"
)

//IQ45sx9s
type msgStruct struct {
	A int `struc:"big"`
	//b uint64 `struc:"int64,little"`
	//s string
	//x  int8 `struc:"int8,little"`
	//s2 string
}

func main() {
	c := msgStruct{}
	msg_data := "\x00\x00\x00\x40"
	byt := []byte(msg_data)
	fmt.Println(byt)
	buf := bytes.NewBuffer(byt)
	err := struc.Unpack(buf, c)
	fmt.Println(err)
	fmt.Println(c)
}
