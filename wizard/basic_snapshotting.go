package wizard

var basic_snapshottingStep = &Step{
	name: "Basic Snapshotting",
	next: []*Step{},
	execute: func(s *Step) {
		println("my Basic Snapshotting function goes here")
	},
}
