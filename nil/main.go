package main

import (
	"errors"
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	// …
	return err
}
func Foo2() error {
	var er error
	er = &os.PathError{Op: "aa", Path: "path", Err: errors.New("err")}
	// …
	return er
}

func Bar() error {
	return nil
}

func main() {
	err := Foo()
	//fmt.Println(err)                         // <nil>
	fmt.Println(err == nil)                  // false
	fmt.Println(err == (*os.PathError)(nil)) // true

	err2 := Bar()
	fmt.Println(err2 == nil) // true

	err3 := Foo2()
	fmt.Println(err3) // <nil>
}
