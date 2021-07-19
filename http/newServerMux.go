package main

import (
	"fmt"
	"net/http"
	"time"
)

type greeting string

func (g greeting) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintln(w, g)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	//mux.Handle("/greeting", PanicRecover(WithLogger(Metric(greeting("Welcome to go web frameworks")))))
	//middlewares := []Middleware{
	//	PanicRecover,
	//	WithLogger,
	//	Metric,
	//}
	mux.Handle("/greeting", applyMiddlewares(greeting("hello world"), PanicRecover, WithLogger, Metric))

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 20 * time.Second,
	}
	_ = server.ListenAndServe()
}
