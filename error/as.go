package main

import (
	"errors"
	"fmt"
)

type TypicalErr struct {
	e string
}

func (t TypicalErr) Error() string {
	return t.e
}

func (t TypicalErr) As(a interface{}) bool {
	return true
}

func main() {
	err := TypicalErr{"typical error"}
	err1 := fmt.Errorf("wrap err: %w", err)
	err2 := fmt.Errorf("wrap err1: %w", err1)
	var e TypicalErr
	if !errors.As(err2, &e) {
		panic("TypicalErr is not on the chain of err2")
	}
	println("TypicalErr is on the chain of err2")
	println(err == e)
}

/*
  打印结果：
  TypicalErr is on the chain of err2
  true
*/
