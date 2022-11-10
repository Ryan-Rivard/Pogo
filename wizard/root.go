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

		options = append(options, "Exit")

		cat := inquier.AskWithOptions("What git category would you like to explore?", options)

		if *cat == "Exit" {
			println("Thank you for using pogo")
			return
		}

		for _, step := range s.next {
			if step.name == *cat {
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
