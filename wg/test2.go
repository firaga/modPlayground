package main

import (
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Add(1)
	wg.Done()
	wg.Done()
	wg.Wait()
	time.Sleep(2 * time.Second)
}
