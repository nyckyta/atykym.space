package main

import (
	"net/http"
	"log"
	"time"
)

func main() {
	s := &http.Server{
		Addr:           "127.0.0.1:8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	http.Handle("/", http.FileServer(http.Dir("./static")))
	log.Fatal(s.ListenAndServe())
}