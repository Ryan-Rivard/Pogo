package main

import (
	"fmt"
)

type pogoLog struct {
	// create your services here
	inquier iInquier
	git     iGit
}

func (l *pogoLog) process() {
	// l.inquier.askWithPresetAnswers("Answer", )

	fmt.Println("outputting commit log. press q to quit")

	// call your services here
	l.git.executeGitCommand("log")

}
