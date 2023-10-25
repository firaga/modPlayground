package main

import (
	"fmt"
	"log"

	"github.com/kelseyhightower/memkv"
)

func main() {
	s := memkv.New()
	s.Set("/myapp/database/username", "admin")
	s.Set("/myapp/database/password", "123456789")
	s.Set("/myapp/port", "80")
	kv, err := s.Get("/myapp/database/username")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Key: %s, Value: %s\n", kv.Key, kv.Value)
	fmt.Println("----GetAll----")
	ks, err := s.GetAll("/myapp/*/*")
	if err == nil {
		for _, kv := range ks {
			fmt.Printf("Key: %s, Value: %s\n", kv.Key, kv.Value)
		}
	}

	//test List
	fmt.Println("----list----")
	res := s.List("/myapp")
	for i, re := range res {
		fmt.Printf("sub path %d:%s\n", i, re)
	}

	//test ListDir
	fmt.Println("----listDir----")
	res2 := s.ListDir("/myapp")
	fmt.Println(res2)
}
