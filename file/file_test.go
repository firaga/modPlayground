package main

import (
	"fmt"
	"os"
	"testing"
)

func TestMkdir(t *testing.T) {
	dir := "/tmp/a/b/"
	err := os.Mkdir(dir, 0731)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("done")
}
