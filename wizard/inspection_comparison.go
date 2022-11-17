package wizard

var inspection_comparisonStep = &Step{
	name: "Inspection and Comparison",
	next: []*Step{},
	execute: func(s *Step) {
		println("my Inspection and Comparison function goes here")
	},
}
