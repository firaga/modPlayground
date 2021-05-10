package main

import (
	"fmt"
	"os"
)

func main() {
	err := getErr()
	if err != nil {
		fmt.Printf("FATAL %+v\n", err)
		fmt.Printf("FATAL %v\n", err)
	}

}

func getErr() error {
	_, err := os.Open("/he")
	if err != nil {
		//return errors.Wrapf()
		return err
	}
	return nil
}
