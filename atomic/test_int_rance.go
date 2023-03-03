package main

import (
	"fmt"
	"strconv"
	"sync"
)

var a int64
var b string

func main() {
	var wg sync.WaitGroup
	c := 1000
	wg.Add(c)
	for i := 0; i < c; i++ {
		go func() {
			//a = 6
			b = "hello" + strconv.Itoa(i)
			//fmt.Println(b)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(b)
}
