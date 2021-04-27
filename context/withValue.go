package main

import (
	"net/http"
	"context"
)

func main() {
	helloWorldHandler := http.HandlerFunc(HelloWorld)
	http.Handle("/hello", inejctRequestId(helloWorldHandler))
	http.ListenAndServe(":8080", nil)
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	requestId := ""
	if m := r.Context().Value("requestId"); m != nil {
		if value, ok := m.(string); ok {
			requestId = value
		}
	}
	w.Header().Add("requestId", requestId)
	w.Write([]byte("Hello, world"))
}

func inejctRequestId(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestId := uuid.New().String()
		ctx := context.WithValue(r.Context(), "requestId", requestId)
		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}
