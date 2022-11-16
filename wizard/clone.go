package wizard

import (
	git "github.com/Ryan-Rivard/Pogo/git"
	inquier "github.com/Ryan-Rivard/Pogo/inquire"
)

func init() {
	cloneStep.execute = createCloneAction(cloneStep)
}

var cloneStep = &Step{
	name: "Clone",
}

func createCloneAction(s *Step) func() {
	return func() {
		repoUrl := inquier.AskPrompt("What is the url of your repo?")

		git.ExecuteGitCommand("clone", *repoUrl)
	}
}
