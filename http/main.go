package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintln(w, "Hello World")
}

func main() {
	http.HandleFunc("/", index)
	_ = http.ListenAndServe(":8080", nil)
}
