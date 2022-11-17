package wizard

var server_adminStep = &Step{
	name: "Server Admin",
	next: []*Step{},
	execute: func(s *Step) {
		println("my Server Admin function goes here")
	},
}
