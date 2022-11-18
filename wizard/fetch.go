package wizard

import (
	git "github.com/Ryan-Rivard/Pogo/git"
)

var fetchStep = &Step{
	name: "fetch",
	next: []*Step{},
	execute: func(s *Step) {
		println("executing git fetch")
		git.ExecuteGitCommand("fetch", "-v")
	},
}
