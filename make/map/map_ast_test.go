package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"testing"
)

func TestParser(t *testing.T) {
	src := []byte(`package main
    import "fmt"

    func main() {
        m := make(map[string]int64, 53)
        m["a"] = 111111111
        fmt.Println(m["a"])
    }
    `)
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		log.Fatal(err)
	}
	ast.Print(fset, file)
}
