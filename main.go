package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
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

	healthHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "OK")
	}

	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/health", healthHandler)
	log.Fatal(s.ListenAndServe())
}
