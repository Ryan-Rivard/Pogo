package wizard

func init() {
	branching_mergingStep.execute = createBranching_MergingsAction(branching_mergingStep)
}

var branching_mergingStep = &Step{
	name: "Branching and Merging",
}

func createBranching_MergingsAction(s *Step) func() {
	return func() {
		println("my Branching and Merging function goes here")
	}
}
