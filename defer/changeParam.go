package main

import "fmt"

func main() {
	fmt.Printf("with named param, x: %d\n", namedParam())
	fmt.Printf("without named param, x: %d\n", notNamedParam())
	fmt.Printf("with named param param, x: %d\n", paramParam(3))
}
func namedParam() (x int) {
	x = 1
	defer func() { x = 2 }()
	return x
}

func notNamedParam() int {
	x := 1
	defer func() { x = 2 }()
	return x
}

func paramParam(x int) int {
	defer func() { x = 2 }()
	return x
}
