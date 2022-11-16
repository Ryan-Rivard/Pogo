package wizard

func init() {
	debuggingStep.execute = createDebuggingAction(debuggingStep)
}

var debuggingStep = &Step{
	name: "Debugging",
}

func createDebuggingAction(s *Step) func() {
	return func() {
		println("my debugging function goes here")
	}
}
