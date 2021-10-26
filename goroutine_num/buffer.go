package main

//
//import (
//	"fmt"
//	"math"
//	"runtime"
//)
//
//// 模拟执行业务的 goroutine
//func doBusiness(ch chan bool, i int) {
//	fmt.Println("go func:", i, "goroutine count:", runtime.NumGoroutine())
//	<-ch
//}
//
//func main() {
//	max_cnt := math.MaxInt64
//	// max_cnt := 10
//	fmt.Println(max_cnt)
//
//	ch := make(chan bool, 3)
//
//	for i := 0; i < max_cnt; i++ {
//		ch <- true
//		go doBusiness(ch, i)
//	}
//}
