package main

import (
	"encoding/json"
	"fmt"
	"github.com/icza/dyno"
)

//https://github.com/icza/dyno
func main() {
	src := `{"login":{"password":"secret","user":"bob"},"name":"cmpA"}`
	var v interface{}
	if err := json.Unmarshal([]byte(src), &v); err != nil {
		panic(err)
	}
	// Edit (mask out) password:
	if err := dyno.Set(v, "xxx", "login", "password"); err != nil {
		fmt.Printf("Failed to set password: %v\n", err)
	}
	edited, err := json.Marshal(v)
	fmt.Printf("Edited JSON: %s, error: %v\n", edited, err)
}
