package wizard

import (
	git "github.com/Ryan-Rivard/Pogo/git"
)

var pullStep = &Step{
	name: "Pull",
	next: []*Step{},
	execute: func(s *Step) {
		println("executing git pull")
		git.ExecuteGitCommand("pull")
	},
}
