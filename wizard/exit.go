package wizard

func init() {
	exitStep.execute = createExitAction(exitStep)
}

var exitStep = &Step{
	name: "Exit",
}

func createExitAction(s *Step) func() {
	return func() {
		println("exitting")
	}
}
