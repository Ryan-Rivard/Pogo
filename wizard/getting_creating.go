package wizard

import (
	inquier "github.com/Ryan-Rivard/Pogo/inquire"
)

func init() {
	rootStep.addNext(getting_creatingStep)

	getting_creatingStep.execute = createGetting_CreatingAction(getting_creatingStep)

	getting_creatingStep.prev = rootStep
}

var getting_creatingStep = &Step{
	name: "Getting and Creating Projects",
}

func createGetting_CreatingAction(s *Step) func() {
	return func() {
		options := []string{}

		for _, step := range s.next {
			options = append(options, step.name)
		}

		options = append(options, "Back")

		search := inquier.AskWithOptions("Would you like to Get or Create a Project?", options)

		if *search == "Back" {
			s.prev.execute()
			return
		}

		for _, step := range s.next {
			if step.name == *search {
				step.execute()
				return
			}
		}
	}
}
