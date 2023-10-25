package main

import "fmt"

func main() {
	m1 := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}

	//note: 不可以在range的同时赋值,不可预期
	for k, v := range m1 {
		m1[k+"_new"] = v + 100
	}

	fmt.Printf("m1 len: %d", len(m1))
}
