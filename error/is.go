package main

import (
	"errors"
	"fmt"
)

var BaseErr = errors.New("the underlying base error")

func main() {
	err1 := fmt.Errorf("wrap base: %w", BaseErr)
	err2 := fmt.Errorf("wrap err1: %w", err1)
	println(err2 == BaseErr) // false
	if !errors.Is(err2, BaseErr) {
		panic("err2 is not BaseErr")
	}
	println("err2 is BaseErr")
}

/*
  打印结果：
  false
  err2 is BaseErr
*/
