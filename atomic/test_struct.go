package main

import (
	"fmt"
	"sync"
)

type Test struct {
	X int
	Y int
}

var aa int32

func main() {
	//var g Test

	for i := 0; i < 1000000; i++ {
		var wg sync.WaitGroup
		// 协程 1
		wg.Add(1)
		go func() {
			defer wg.Done()
			//g = Test{1, 2}
			aa = 65535
		}()

		// 协程 2
		wg.Add(1)
		go func() {
			defer wg.Done()
			//g = Test{3, 4}
			aa = 0
		}()
		wg.Wait()

		// 赋值异常判断
		if aa != 0 && aa != 65535 {
			fmt.Printf("concurrent assignment error,i=%v a=%v", i, aa)
			break
		}
		//if !((g.X == 1 && g.Y == 2) || (g.X == 3 && g.Y == 4)) {
		//	fmt.Printf("concurrent assignment error, i=%v g=%+v", i, g)
		//	break
		//}
	}
}
