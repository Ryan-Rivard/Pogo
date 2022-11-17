package wizard

var exitStep = &Step{
	name: "Exit",
	execute: func(s *Step) {
		println("my Exit goes here")
	},
}
