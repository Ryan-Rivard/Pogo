package wizard

func init() {
	guidesStep.execute = createGuidesAction(guidesStep)
}

var guidesStep = &Step{
	name: "Guides",
}

func createGuidesAction(s *Step) func() {
	return func() {
		println("my Guides function goes here")
	}
}
