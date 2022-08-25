package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	s := "\x00\x00\x00\x40\x00\x00\x00\x00\x00\x00\x00\x00\x7B\x22\x63\x6F\x6E\x74\x65\x6E\x74\x22\x3A\x22\x68\x65\x6C\x6C\x6F\x22\x2C\x22\x73\x74\x72\x54\x61\x6C\x6B\x65\x72\x22\x3A\x22\x31\x32\x31\x30\x30\x30\x30\x30\x30\x37\x31\x22\x7D\x00\x00\x00\xE6\x88\x98\xE5\x8F\x8B\x00"
	msgData := []byte(s)
	fmt.Println(msgData)

	c := msgData[12:]
	pos, err := findPos(c, "\x00")
	if err != nil {
		os.Exit(1)
	}
	//pos := strings.Index(c, "\x00")
	//fmt.Println(pos)
	c = c[:pos]
	chatMessage := string(c)
	fmt.Println(chatMessage, len(chatMessage))
	//var v map[string]interface{}
	//json.Unmarshal([]byte(c), &v)
	//fmt.Println(v)
}

func findPos(byes []byte, f string) (int, error) {
	for pos, byt := range byes {
		if string(byt) == f {
			return pos, nil
		}
	}
	return 0, errors.New("pos not find")
}
