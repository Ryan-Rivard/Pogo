package wizard

var sharing_updatingStep = &Step{
	name: "Sharing and Updating Projects",
	next: []*Step{},
	execute: func(s *Step) {
		println("my Sharing and Updating Projects function goes here")
	},
}
