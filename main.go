package main

import (
	"fmt"
	"log"
	"io"
	"net/http"
	"os"
	"time"
)

func logReq(w http.ResponseWriter, r *http.Request) {
	b, e := io.ReadAll(r.Body)
	if e != nil {
		log.Printf("[WARN] Failed to handle req %e", e)
		w.WriteHeader(400)
		// TODO: inject html
		w.Write([]byte("Segmentation fault!!!"))
		return
	}

	log.Printf("Body %s", string(b))
	w.WriteHeader(200)
	w.Write([]byte(`<div class="terminal">Unknown command</div>`))
}

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
	http.Handle("/enter",  http.HandlerFunc(logReq))
	http.HandleFunc("/health", healthHandler)
	log.Fatal(s.ListenAndServe())
}
