package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"time"
)

//https://github.com/go-ini/ini
type Note struct {
	Content string
	Cities  []string
}

type Person struct {
	Name string
	Age  int `ini:"age"`
	Male bool
	Born time.Time
	Note
	Created time.Time `ini:"-"`
}

func main() {
	cfg, err := ini.Load("e.ini")
	if err != nil {
		fmt.Println(err)
	}
	// ...
	p := new(Person)
	var a map[string]interface{}
	err = cfg.MapTo(&a)
	fmt.Println(a)
	// ...

	// 一切竟可以如此的简单。
	err = ini.MapTo(p, "path/to/ini")
	// ...

	// 嗯哼？只需要映射一个分区吗？
	n := new(Note)
	err = cfg.Section("Note").MapTo(n)
	// ...
}
