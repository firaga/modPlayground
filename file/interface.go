package main

import "fmt"

type Component interface {
	hello()
}

func NewComA() Component {
	return ComA{}
}

type ComA struct {
	a int
}

func (ComA) hello() {
	fmt.Println("A")
}

func NewComB() Component {
	return ComB{}
}

type ComB struct {
	a int
}

func (c ComB) hello() {
	fmt.Println("B")
}

func NewComC() Component {
	return &ComC{}
}

type ComC struct {
	a int
}

func (c *ComC) hello() {
	fmt.Println("C")
}

func (c ComC) hello2() {
	fmt.Println("C2")
}

func main() {
	c := &ComC{}
	c.hello()
	c.hello2()
	(*c).hello()
	(*c).hello2()

	c2 := ComC{}
	c2.hello()
	c2.hello2()
	(&c2).hello()
	(&c2).hello2()
	//结论: 无论是否引用,都能调用所有方法
	//return

	var cs []Component
	cs = append(cs, ComA{})
	cs = append(cs, ComB{})
	cs = append(cs, &ComC{}) //接口中,实现方法和变量的指针需要对应上 //比如方法是指针实现,那么实现接口变量必须取地址
	for _, c := range cs {
		c.hello()
	}
}
