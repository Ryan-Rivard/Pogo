package wizard

var exitStep = &Step2{
	name: "Exit",
	execute: func(s *Step2) {
		println("my Exit goes here")
	},
}
