package wizard

func init() {
	rootStep.addNext(statusStep)

	statusStep.execute = createStatusAction(statusStep)
}

var statusStep = &Step{
	name: "status",
}

func createStatusAction(s *Step) func() {
	return func() {
		println("You chose Status")
	}
}
