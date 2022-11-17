package wizard

var initStep = &Step{
	name: "Init",
	next: []*Step{},
	execute: func(s *Step) {
		println("my Init function goes here")
	},
}
