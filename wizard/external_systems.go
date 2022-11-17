package wizard

var external_systemsStep = &Step{
	name: "External Systems",
	next: []*Step{},
	execute: func(s *Step) {
		println("my External Systems function goes here")
	},
}
