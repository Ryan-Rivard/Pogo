package wizard

var setup_configStep = &Step{
	name: "Setup and Config",
	next: []*Step{},
	execute: func(s *Step) {
		println("my Setup and Config function goes here")
	},
}
