package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Name string
	Age  int
}

type Department struct {
	Name string
	Age  int
}

func main() {
	var p Person = Person{Name: "John", Age: 30}
	var e Employee = Employee{Name: "Babala", Age: 18}
	var d Department
	p = Person(e) // 把 e 转成Person
	fmt.Println(p)
	d = Department(e)
	fmt.Println(d)
}

//输出 {Babala 18}
