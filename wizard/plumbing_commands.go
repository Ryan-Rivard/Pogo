package wizard

func init() {
	plumbing_commandsStep.execute = createplumbing_commandsAction(plumbing_commandsStep)
}

var plumbing_commandsStep = &Step{
	name: "Plumbing Commands",
}

func createplumbing_commandsAction(s *Step) func() {
	return func() {
		println("my Plumbing Commands function goes here")
	}
}
