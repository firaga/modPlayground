package main

import (
	"log"
	"net"
	"os"
)

type App struct {
	File *os.File
	Conn net.Conn
}

func provideFile() (*os.File, func(), error) {
	f, err := os.Open("foo.txt")
	if err != nil {
		return nil, nil, err
	}
	cleanup := func() {
		if err := f.Close(); err != nil {
			log.Println(err)
		}
	}
	return f, cleanup, nil
}

func provideNetConn() (net.Conn, func(), error) {
	conn, err := net.Dial("tcp", "foo.com:80")
	if err != nil {
		return nil, nil, err
	}
	cleanup := func() {
		if err := conn.Close(); err != nil {
			log.Println(err)
		}
	}
	return conn, cleanup, nil
}

//https://medium.com/@dche423/master-wire-cn-d57de86caa1b
//带清理函数的wire
func main() {
	_, cleanup, err := NewApp()
	if err != nil {
		log.Fatal(err)
	}
	defer cleanup()
}
