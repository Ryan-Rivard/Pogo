package wizard

import (
	git "github.com/Ryan-Rivard/Pogo/git"
)

var fetchStep = &Step2{
	name: "Fetch",
	next: []*Step2{},
	execute: func(s *Step2) {
		println("executing git fetch")
		git.ExecuteGitCommand("fetch", "-v")
	},
}
