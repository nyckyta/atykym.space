package handlers

import (
	"bufio"
	"html/template"
	"strings"
)

type Cmd struct {
	Cmd string
	Params[] string
	User string 
}

type CommandRunner interface {
	Run() string;
}

func (cmd Cmd) Run() string {
	switch(cmd.Cmd) {
	case "": return "" 
	default: return handleUnknown(cmd)
	}
}

func handleUnknown(cmd Cmd) string {
	t, e := template.ParseFiles("templates/unknown-command-template.html")
	if e != nil {
		panic("Failed to parse unknown command template")
	}
	b := strings.Builder{}
	w := bufio.NewWriter(&b)
	e = t.Execute(w, cmd)
	w.Flush()
	if e != nil {
		panic("Failed to execute template")
	}
	s := b.String()
	return s
}
