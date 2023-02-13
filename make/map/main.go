package main

import (
	"fmt"
)

type H map[string]string

func main() {
	//a:=H{}
	//a["code"] = "100"

	//b:=map[string]string{}

	m := make(map[string]int64, 53)
	m["a"] = 111111111
	m["b"] = 111111111
	fmt.Println(m["a"], m["b"])
	delete(m, "a")
	fmt.Println(m["a"], m["b"])
}
