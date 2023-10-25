package main

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"log"
	"time"
)

//https://github.com/kelseyhightower/envconfig

type Specification struct {
	Debug      bool
	Port       int
	User       string `json:"user" xml:"user" default:"Tom"`
	Users      []string
	Rate       float32
	Timeout    time.Duration
	ColorCodes map[string]int
}

func main() {
	var s Specification
	err := envconfig.Process("myapp", &s)
	if err != nil {
		log.Fatal(err.Error())
	}
	format := "Debug: %v\nPort: %d\nUser: %s\nRate: %f\nTimeout: %s\n"
	_, err = fmt.Printf(format, s.Debug, s.Port, s.User, s.Rate, s.Timeout)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Users:")
	for _, u := range s.Users {
		fmt.Printf("  %s\n", u)
	}

	fmt.Println("Color codes:")
	for k, v := range s.ColorCodes {
		fmt.Printf("  %s: %d\n", k, v)
	}
}
