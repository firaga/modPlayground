package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Printf("error from err1: %v\n", err1())
	fmt.Printf("error from err2: %v\n", err2())
}

func err1() error {
	var err error

	defer func() {
		if r := recover(); r != nil {
			err = errors.New("recovered err")
		}
	}()
	panic(`foo`)

	return err
}

func err2() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("recovered err2")
		}
	}()
	panic(`foo`)

	return err
}
