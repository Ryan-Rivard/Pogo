package wizard

var emailStep = &Step{
	name: "Email",
	next: []*Step{},
	execute: func(s *Step) {
		println("my Email function goes here")
	},
}
