package cmd

import (
	"bufio"
	"html/template"
	"io"
	"io/fs"
	"log"
	"os"
	"path"
	"strings"
)

const ERROR_TEMPLATE_PATH = "templates/error.html"
const CAT_OUTPUT_PATH = "templates/cat-output.html"
const HELP_OUTPUT_PATH = "templates/help.html"
const LS_OUTPUT_PATH = "templates/ls-output.html"

type Cmd struct {
	Cmd    string
	Params []string
}

type CmdErr struct {
	Cmd    string
	ErrMsg string
}

type CommandRunner interface {
	Run() string
}

func (cmd Cmd) Run() string {
	switch cmd.Cmd {
	case "cat":
		return handleCat(cmd)
	case "ls":
		return handleLs(cmd)
	case "help":
		return handleHelp()
	case "":
		return ""
	default:
		return handleUnknown(cmd)
	}
}

func handleLs(cmd Cmd) string {
	var directoryToSearch string
	if len(cmd.Params) < 1 {
		directoryToSearch = "."
	}
	files, err := lsGetFiles(directoryToSearch)

	if err != nil {
		log.Printf("[ERROR] Failed to get list of files")
		return executeTemplateAgainstAny(ERROR_TEMPLATE_PATH, "Failed to read directory")
	}
	return executeTemplateAgainstAny(LS_OUTPUT_PATH, files)

}

func lsGetFiles(dir string) ([]string, error) {
	_fs := os.DirFS("anon/home").(fs.ReadDirFS)
	path := path.Clean(dir)
	files, err := _fs.ReadDir(path)

	if err != nil {
		return nil, err
	}

	filesInDirectory := make([]string, len(files))
	for i, v := range files {
		filesInDirectory[i] = v.Name()
	}

	return filesInDirectory, nil
}

func handleUnknown(cmd Cmd) string {
	err := CmdErr{cmd.Cmd, "unknown command."}
	return executeTemplateAgainstAny(ERROR_TEMPLATE_PATH, err)
}

func handleHelp() string {
	return executeTemplateAgainstAny(HELP_OUTPUT_PATH, nil)
}

func handleCat(cmd Cmd) string {
	params := cmd.Params
	if len(params) < 1 {
		err := CmdErr{cmd.Cmd, "no file provided."}
		return executeTemplateAgainstAny(ERROR_TEMPLATE_PATH, err)
	}

	fs := os.DirFS("anon/home")

	path := path.Clean(params[0])

	f, err := fs.Open(path)
	if err != nil {
		errMsg := err.Error()
		if strings.HasSuffix(errMsg, "invalid argument") {
			// invalid argument commonly happens when there is an attempt to get access to the directory outside of file system
			errMsg = "permission denied: " + path
		} else if strings.HasSuffix(errMsg, "no such file or directory") {
			// strip the system call eg. 'open <file>: no such file or directory'
			errMsg = path + ": no such file or directory"
		}
		cmdErr := CmdErr{cmd.Cmd, errMsg}
		return executeTemplateAgainstAny(ERROR_TEMPLATE_PATH, cmdErr)
	}

	fileContent, err := io.ReadAll(f)
	if err != nil {
		cmdErr := CmdErr{cmd.Cmd, err.Error()}
		return executeTemplateAgainstAny(ERROR_TEMPLATE_PATH, cmdErr)
	}

	return executeTemplateAgainstAny(CAT_OUTPUT_PATH, string(fileContent))
}

func executeTemplateAgainstAny(pathToTemplate string, vars any) string {
	t, e := template.ParseFiles(pathToTemplate)
	if e != nil {
		panic("Failed to parse unknown command template, " + e.Error())
	}
	b := strings.Builder{}
	w := bufio.NewWriter(&b)
	e = t.Execute(w, vars)
	w.Flush()
	if e != nil {
		panic("Failed to execute template")
	}
	return b.String()
}

/*
Converts string to Command that is able to run.
*/
func ToCommand(input string) CommandRunner {
	tokens := strings.Fields(input)
	if len(tokens) > 1 {
		return Cmd{
			Cmd:    tokens[0],
			Params: tokens[1:],
		}
	}

	if len(tokens) == 1 {
		return Cmd{
			Cmd:    tokens[0],
			Params: []string{},
		}
	}

	return Cmd{
		Cmd:    "",
		Params: []string{},
	}
}
