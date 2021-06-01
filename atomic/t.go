package main

import "fmt"

func main() {
	a := uintptr(0)
	fmt.Println(a)
	fmt.Println(^a)
	fmt.Printf("%b\n", a)
	fmt.Printf("%b\n", ~5)
}
