package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	e := errors.New("脑子进煎鱼了")
	w := fmt.Errorf("快抓住,%w", e)
	fmt.Println(w)
	fmt.Println(errors.Unwrap(w))

	if _, err := os.Open("non-existing"); err != nil {
		var pathError *os.PathError
		if errors.As(err, &pathError) {
			fmt.Println("Failed at path:", pathError.Path)
		} else {
			fmt.Println(err)
		}
	}

}
