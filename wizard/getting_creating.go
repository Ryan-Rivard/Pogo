package wizard

import (
	inquier "github.com/Ryan-Rivard/Pogo/inquire"
)

var getting_creatingStep = &Step{
	name: "Getting and Creating Projects",
	next: []*Step{},
	execute: func(s *Step) {
		options := []string{}

		for _, step := range s.next {
			options = append(options, step.name)
		}

		options = append(options, "Back")

		search := inquier.AskWithOptions("Would you like to Get or Create a Project?", options)

		for _, step := range s.next {
			if step.name == *search {
				step.execute(step)
				return
			}
		}
	},
}
