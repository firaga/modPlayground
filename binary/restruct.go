package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/go-restruct/restruct"
)

type Record struct {
	Message string `struct:"[128]byte"`
}

type Container struct {
	Version   int `struct:"int32"`
	NumRecord int `struct:"int32,sizeof=Records"`
	Records   []Record
}

type T3 struct {
	A int    `struct:"int32"`
	S string `struct:"[11]byte"`
	B int    `struct:"int32"`
}

func main() {
	var c T3

	//file, _ := os.Open("records")
	//defer file.Close()
	t := T3{A: 0xEEFFEEFF, S: "hello world", B: 10}
	buf := &bytes.Buffer{}
	err := binary.Write(buf, binary.BigEndian, t)
	if err != nil {
		panic(err)
	}
	fmt.Println(buf.Bytes())

	restruct.Unpack(buf.Bytes(), binary.LittleEndian, &c)
	fmt.Println(c)
}
