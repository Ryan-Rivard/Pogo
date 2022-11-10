package wizard

func init() {
	rootStep.addNext(versionStep)

	versionStep.execute = createVersionAction(versionStep)
}

var versionStep = &Step{
	name: "version",
}

func createVersionAction(s *Step) func() {
	return func() {
		println("You chose Version")
	}
}
