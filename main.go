package main

import (
	"net/http"
	"log"
	"time"
	"os"
)

func main() {
	port, isPresent := os.LookupEnv("PORT")
	if !isPresent {
		panic("PORT env must be present")
	}
	s := &http.Server{
		Addr:           ":" + port,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	http.Handle("/", http.FileServer(http.Dir("./static")))
	log.Fatal(s.ListenAndServe())
}