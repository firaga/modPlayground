package main

import "fmt"

func main() {
	a := []string{"a"}
	fmt.Println(a)
	i := 0
	b := append(a[:i], a[i+1:]...)
	fmt.Println(b)
}
