package main

import "fmt"

type inter interface {
	Print()
}
type impl struct {
	a string
}

func NewImpl() inter {
	return &impl{a: "hello"}
}

func (i *impl) Print() {
	fmt.Println(i.a)
}

func main() {
	var a inter
	a = NewImpl()
	v := a.(*impl)
	fmt.Println(v)
	a.Print()
}
