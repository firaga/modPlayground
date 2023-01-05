package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 5}
	fmt.Println(a)
	b := a[:0]
	fmt.Println(b)
}
