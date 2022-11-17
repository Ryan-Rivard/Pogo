package wizard

var plumbing_commandsStep = &Step{
	name: "Plumbing Commands",
	next: []*Step{},
	execute: func(s *Step) {
		println("my Plumbing Commands function goes here")
	},
}
