package wizard

func init() {
	getting_creatingStep.addNext(initStep)

	initStep.execute = createInitAction(initStep)
}

var initStep = &Step{
	name: "Init",
}

func createInitAction(s *Step) func() {
	return func() {
		println("my init function goes here")
	}
}
