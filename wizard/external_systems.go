package wizard

func init() {
	external_systemsStep.execute = createExternal_SystemsAction(external_systemsStep)
}

var external_systemsStep = &Step{
	name: "External Systems",
}

func createExternal_SystemsAction(s *Step) func() {
	return func() {
		println("my External Systems function goes here")
	}
}
