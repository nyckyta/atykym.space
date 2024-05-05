package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
	"atykym.space/handlers"
)

const INTERNAL_ERROR_RESPONSE string = `<div class="terminal">Segmenation fault!!! Please, let me know when it happens.</div>`


func toCommand(input string) handlers.CommandRunner {
	tokens := strings.Fields(input)
	if (len(tokens) > 1) {
		return handlers.Cmd{
			Cmd: tokens[0],
			Params: tokens[1:],
		}
	}

	if (len(tokens) == 1) {
		return handlers.Cmd{
			Cmd: tokens[0],
			Params: []string{},
		}
	}

	return handlers.Cmd{
		Cmd: "",
		Params: []string{},
	}
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	t, e := template.ParseFiles("./templates/index-template.html")
	if e != nil {
		log.Printf("[WARN] Failed to handle req %e", e)
		w.WriteHeader(500)
		w.Write([]byte(INTERNAL_ERROR_RESPONSE))
		return
	}
	w.WriteHeader(200)
	t.Execute(w, "anon")
}

func handleEnterHit(w http.ResponseWriter, r *http.Request) {
	e := r.ParseForm()
	if e != nil {
		log.Printf("[WARN] Failed to handle req %e", e)
		w.WriteHeader(500)
		w.Write([]byte(INTERNAL_ERROR_RESPONSE))
		return
	}
	cmdVal := r.FormValue("cmd")
	cmd := toCommand(cmdVal)
	log.Printf("%v", cmd)
	output := cmd.Run()
	w.WriteHeader(200)
	w.Write([]byte(output))
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}

func handleFavicon(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
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

	http.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/enter", handleEnterHit)
	http.HandleFunc("/favicon.ico", handleFavicon)
	http.HandleFunc("/health", handleHealth)
	http.HandleFunc("/", handleIndex)
	log.Fatal(s.ListenAndServe())
}
