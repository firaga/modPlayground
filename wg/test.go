package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Add(1)
	go func() {
		fmt.Println("b")
		wg.Wait()
		fmt.Println("n")
	}()
	time.Sleep(20 * time.Second)
	wg.Done()
	wg.Done()
	wg.Wait()
	time.Sleep(2 * time.Second)
}
