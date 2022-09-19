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

func main() {
	var cs []Component
	cs = append(cs, ComA{})
	cs = append(cs, &ComB{})
	cs = append(cs, &ComC{})
	for _, c := range cs {
		c.hello()
	}
}
