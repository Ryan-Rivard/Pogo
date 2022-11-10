package wizard

import (
	inquier "github.com/Ryan-Rivard/Pogo/inquire"
)

func init() {
	rootStep.addNext(setup_configStep)

	setup_configStep.execute = createSetup_ConfigAction(setup_configStep)
}

var setup_configStep = &Step{
	name: "Setup and Config",
}

func createSetup_ConfigAction(s *Step) func() {
	return func() {
		options := []string{}

		for _, step := range s.next {
			options = append(options, step.name)
		}

		search := inquier.AskWithOptions("What git category would you like to explore?", options)

		for _, step := range s.next {
			if step.name == *search {
				step.execute()
				break
			}
		}
	}
}
