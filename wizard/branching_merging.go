package wizard

var branching_mergingStep = &Step{
	name: "Branching and Merging",
	next: []*Step{},
	execute: func(s *Step) {
		println("my Branching and Merging function goes here")
	},
}
