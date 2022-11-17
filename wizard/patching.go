package wizard

var patchingStep = &Step{
	name: "Patching",
	next: []*Step{},
	execute: func(s *Step) {
		println("my Patching function goes here")
	},
}
