package wizard

func init() {
	getting_creatingStep.addNext(cloneStep)

	cloneStep.execute = createCloneAction(cloneStep)
}

var cloneStep = &Step{
	name: "Clone",
}

func createCloneAction(s *Step) func() {
	return func() {
		println("my clone function goes here")
	}
}
