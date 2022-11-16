package wizard

func init() {
	patchingStep.execute = createPatchingAction(patchingStep)
}

var patchingStep = &Step{
	name: "Patching",
}

func createPatchingAction(s *Step) func() {
	return func() {
		println("my Patching function goes here")
	}
}
