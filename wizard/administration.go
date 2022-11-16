package wizard

func init() {
	administrationStep.execute = createAdministrationAction(administrationStep)
}

var administrationStep = &Step{
	name: "Administration",
}

func createAdministrationAction(s *Step) func() {
	return func() {
		println("my Administration function goes here")
	}
}
