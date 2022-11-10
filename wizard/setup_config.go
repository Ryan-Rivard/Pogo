package wizard

func init() {
	rootStep.addNext(setup_configStep)

	setup_configStep.execute = createSetup_ConfigAction(setup_configStep)
}

var setup_configStep = &Step{
	name: "Setup and Config",
}

func createSetup_ConfigAction(s *Step) func() {
	return func() {
		println("You chose Setup and Config")
	}
}
