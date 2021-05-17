package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("defer1")
		defer func() {
			fmt.Println("defer1-1")
		}()
	}()
	defer func() {
		fmt.Println("defer2")
		defer func() {
			fmt.Println("defer2-1")
		}()
		defer func() {
			fmt.Println("defer2-2")
		}()
	}()
}
