package wizard

import (
	inquier "github.com/Ryan-Rivard/Pogo/inquire"
)

var cloneStep = &Step{
	name: "Clone",
	next: []*Step{},
	execute: func(s *Step) {
		// get url
		inquier.AskPrompt("What is the url of your repo?")

		// git.ExecuteGitCommand("clone", *repoUrl)
		println("my Clone function goes here")
	},
}
