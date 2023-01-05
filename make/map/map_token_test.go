package main

import (
	"fmt"
	"go/scanner"
	"go/token"
	"testing"
)

func TestToken(t *testing.T) {
	src := []byte(`package main
    import "fmt"

    func main() {
        m := make(map[string]int64, 53)
        m["a"] = 111111111
        fmt.Println(m["a"])
    }
    `)
	var s scanner.Scanner
	fset := token.NewFileSet()
	file := fset.AddFile("", fset.Base(), len(src))
	s.Init(file, src, nil, 0)
	for {
		pos, tok, lit := s.Scan()
		if tok == token.EOF {
			break
		}
		fmt.Printf("%s\t%s\t%q\n", fset.Position(pos), tok, lit)
	}
}
