package wizard

var guidesStep = &Step{
	name: "Guides",
	next: []*Step{},
	execute: func(s *Step) {
		println("my Guides function goes here")
	},
}
