package wizard

import (
	inquier "github.com/Ryan-Rivard/Pogo/inquire"
)

func init() {
	rootStep.execute = createRootAction(rootStep)
}

var rootStep = &Step{
	name: "root",
}

func createRootAction(s *Step) func() {
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

// func askRoot() *string {
// 	options := []string{
// 		"status",
// 		"version",
// 	}
// 	return inquier.AskWithOptions("What git command would you like to execute?", options)
// }