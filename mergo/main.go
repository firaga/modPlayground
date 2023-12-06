package main

import (
	"dario.cat/mergo"
	"fmt"
	"log"
)

type redisConfig struct {
	Address string
	Port    int
	DB      int
}

type redisConfig2 struct {
	Address string
	Port    int
	DB      int
}

var defaultConfig = redisConfig2{
	Address: "127.0.0.1",
	Port:    6381,
	DB:      1,
}

func main() {
	var config redisConfig

	if err := mergo.Merge(&config, defaultConfig); err != nil {
		log.Fatal(err)
	}

	fmt.Println("redis address: ", config.Address)
	fmt.Println("redis port: ", config.Port)
	fmt.Println("redis db: ", config.DB)

	var m = make(map[string]interface{})
	if err := mergo.Map(&m, defaultConfig); err != nil {
		log.Fatal(err)
	}

	fmt.Println(m)
}
