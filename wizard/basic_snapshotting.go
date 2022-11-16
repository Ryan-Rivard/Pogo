package wizard

func init() {
	basic_snapshottingStep.execute = createBasic_SnapshottingAction(basic_snapshottingStep)
}

var basic_snapshottingStep = &Step{
	name: "Basic Snapshotting",
}

func createBasic_SnapshottingAction(s *Step) func() {
	return func() {
		println("my basic snapshotting function goes here")
	}
}
