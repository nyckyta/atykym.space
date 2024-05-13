package app

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
	"atykym.space/app/cmd"
	"atykym.space/app/config"
)

const INTERNAL_ERROR_RESPONSE string = `<div class="terminal">Segmenation fault!!! Please, let me know when it happens.</div>`

// http handlers
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

func handleHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}

func handleFavicon(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/favicon.svg")
}

func handleEnterHit(w http.ResponseWriter, r *http.Request) {
	e := r.ParseForm()
	if e != nil {
		log.Printf("[WARN] Failed to handle req %e", e)
		w.WriteHeader(500)
		w.Write([]byte(INTERNAL_ERROR_RESPONSE))
		return
	}
	cmdStr := r.FormValue("cmd")
	cmd := cmd.ToCommand(cmdStr)
	output := cmd.Run()
	w.WriteHeader(200)
	w.Write([]byte(output))
}

type Application interface {
	Start()
}

type App struct {
	Config config.AppConfig
}

func (app *App) Start() {
	config := app.Config
	s := &http.Server{
		Addr:           config.Address,
		ReadTimeout:    time.Duration(config.ReadTimeoutSeconds * 1000_000),
		WriteTimeout:   time.Duration(config.WriteTimeoutSeconds * 1000_000),
		MaxHeaderBytes: int(config.MaxHeaderBytes),
	}

	http.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/enter", handleEnterHit)
	http.HandleFunc("/favicon.ico", handleFavicon)
	http.HandleFunc("/health", handleHealth)
	http.HandleFunc("/", handleIndex)
	log.Fatal(s.ListenAndServe())
}