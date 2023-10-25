package main

import "fmt"

type T struct{}

func (T) M1()  { fmt.Println("m1") }
func (*T) M2() { fmt.Println("m2") }

type T1 T

func (T1) M3()  { fmt.Println("m3") }
func (*T1) M4() { fmt.Println("m4") }

func main() {
	var t T
	var pt *T
	pt = &t
	//var t1 T1
	//var pt1 *T1
	t.M1()
	t.M2()
	pt.M1()
	pt.M2()
}
