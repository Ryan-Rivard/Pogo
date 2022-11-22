package wizard

import (
	inquier "github.com/Ryan-Rivard/Pogo/inquire"
)

var sharing_updatingStep = &Step{
	name: "Sharing and Updating Projects",
	next: []*Step{
		fetchStep,
		pullStep,
		pushStep,
	},
	execute: func(s *Step) {
		options := []string{}

		for _, step := range s.next {
			options = append(options, step.name)
		}

		if s.prev.name == "Exit" {
			options = append(options, s.prev.name)
		} else {
			options = append(options, "Back")
		}

		cat := inquier.AskWithOptions("Sharing and Updating:", options)

		if *cat == "Back" {
			s.prev.execute(s.prev)
			return
		}

		for _, step := range s.next {
			if step.name == *cat {
				step.execute(step)
				break
			}
		}
	},
}
