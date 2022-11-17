package wizard

var debuggingStep = &Step{
	name: "Debugging",
	next: []*Step{},
	execute: func(s *Step) {
		println("my Debugging function goes here")
	},
}
