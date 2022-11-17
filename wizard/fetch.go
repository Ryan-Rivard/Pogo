package wizard

import (
	git "github.com/Ryan-Rivard/Pogo/git"
)

var fetchStep = &Step{
	name: "root",
	next: []*Step{},
	prev: sharing_updatingStep,
	execute: func(s *Step) {
		println("executing git fetch")
		git.ExecuteGitCommand("fetch", "-v")
	},
}
