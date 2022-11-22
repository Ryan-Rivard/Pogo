package wizard

import (
	git "github.com/Ryan-Rivard/Pogo/git"
)

var pushStep = &Step{
	name: "Push",
	next: []*Step{},
	execute: func(s *Step) {
		println("executing git push")
		git.ExecuteGitCommand("push")
	},
}
