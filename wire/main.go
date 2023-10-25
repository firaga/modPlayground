package main

import "fmt"

type Message struct {
	msg string
}
type Greeter struct {
	Message Message
}
type Greeter2 struct {
	Message Message
}

type Event struct {
	Greeter  Greeter
	Greeter2 Greeter2
}

// NewMessage Message的构造函数
func NewMessage(msg string) Message {
	return Message{
		msg: msg,
	}
}

// NewGreeter Greeter构造函数
func NewGreeter(m Message) Greeter {
	return Greeter{Message: m}
}

func NewGreeter2(m Message) Greeter2 {
	return Greeter2{Message: m}
}

// NewEvent Event构造函数
func NewEvent(g Greeter, g2 Greeter2) Event {
	return Event{Greeter: g, Greeter2: g2}
}
func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
	fmt.Println(e.Greeter2.Greet())
}
func (g Greeter) Greet() Message {
	return g.Message
}

func (g Greeter2) Greet() Message {
	return g.Message
}

// 使用wire前
func main() {
	message := NewMessage("hello world")
	message2 := NewMessage("hello world2")
	greeter := NewGreeter(message)
	greeter2 := NewGreeter2(message2)
	event := NewEvent(greeter, greeter2)

	event.Start()
}

/*
// 使用wire后
func main() {
	event := InitializeEvent("hello_world")

	event.Start()
}*/
