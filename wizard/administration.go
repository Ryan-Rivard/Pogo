package wizard

var administrationStep = &Step{
	name: "Administration",
	next: []*Step{},
	execute: func(s *Step) {
		println("my Administration function goes here")
	},
}
