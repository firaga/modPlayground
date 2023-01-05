package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int64, 53)
	m["a"] = 111111111
	m["b"] = 111111111
	fmt.Println(m["a"], m["b"])
	delete(m, "a")
	fmt.Println(m["a"], m["b"])
}
