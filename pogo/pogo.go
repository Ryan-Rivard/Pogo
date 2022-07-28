package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/AlecAivazis/survey/v2"
)

// the questions to ask
var qs = []*survey.Question{
	{
		Name: "gitcommand",
		Prompt: &survey.Select{
			Message: "What would you like to do?",
			Options: []string{"init", "clone", "log", "switch/branch", "add", "reset", "diff/status", "commit", "restore", "merge", "rebase", "tag"},
			Default: "status",
		},
	},
}

// git add .
// git reset

func main() {
	fmt.Println("Hello World")

	answers := struct {
		GitCommand string `survey:"gitcommand"` // or you can tag fields to match a specific name
	}{}

	err := survey.Ask(qs, &answers)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("executing command: %s.\n", answers.GitCommand)

	ExecGitCmd(answers.GitCommand)

	// ExecGitCmd("add", ".")
}

func ExecGitCmd(gitCmd ...string) {
	cmd := exec.Command("git", gitCmd...)

	var stdBuffer bytes.Buffer
	mw := io.MultiWriter(os.Stdout, &stdBuffer)

	cmd.Stdout = mw
	cmd.Stderr = mw

	// Execute the command
	if err := cmd.Run(); err != nil {
		log.Panic(err)
	}

	// automatically see it on the terminal
	// log.Println(stdBuffer.String())
}
