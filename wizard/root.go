package wizard

import (
	inquier "github.com/Ryan-Rivard/Pogo/inquire"
)

var rootStep = &Step{
	name: "root",
	next: []*Step{
		setup_configStep,
		getting_creatingStep,
		basic_snapshottingStep,
		branching_mergingStep,
		sharing_updatingStep,
		inspection_comparisonStep,
		patchingStep,
		debuggingStep,
		guidesStep,
		emailStep,
		external_systemsStep,
		administrationStep,
		server_adminStep,
		plumbing_commandsStep,
	},
	execute: func(s *Step) {
		options := []string{}

		for _, step := range s.next {
			options = append(options, step.name)
		}

		options = append(options, s.prev.name)

		cat := inquier.AskWithOptions("What git category would you like to explore?", options)

		for _, step := range s.next {
			if step.name == *cat {
				step.execute(step)
				break
			}
		}
	},
}
