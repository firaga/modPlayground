package main

import "fmt"

func main() {
	fmt.Printf("with named param, x: %d\n", namedParam())
	fmt.Printf("without named param, x: %d\n", notNamedParam())
	fmt.Printf("with named param, x: %d\n", paramParam(3))
	//a()
	fmt.Printf("c func %d\n", c())
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

func c() (i int) {
	defer func() {
		//fmt.Println(i)
		//2a i++
		i++
		//2b i重新赋值
		//i=100
	}()
	//1. 赋值 i=10
	//3. 返回i
	return 10
}

//函数可能是将当前参数拷贝到defer结构中,函数是需要时去取
func a() {
	i := 0
	defer func() {
		fmt.Println("lest 2", i)
	}()
	defer fmt.Println(i)
	defer func() {
		fmt.Println("lest 1", i)
	}()
	i++
	defer fmt.Println(i)
	return
}
