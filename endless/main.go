package main

import (
	"github.com/fvbock/endless"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	duration, err := time.ParseDuration(r.FormValue("duration"))
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	time.Sleep(duration)
	w.Write([]byte("Hello World333"))
}

func main() {
	mux1 := mux.NewRouter()
	mux1.HandleFunc("/sleep", handler)

	w := sync.WaitGroup{}
	w.Add(1)
	go func() {
		err := endless.ListenAndServe("127.0.0.1:5003", mux1)
		if err != nil {
			log.Println(err)
		}
		log.Println("Server on 5003 stopped")
		w.Done()
	}()
	w.Wait()
	log.Println("All servers stopped. Exiting.")

	os.Exit(0)
}
