package wizard

func init() {
	inspection_comparisonStep.execute = createInspection_ComparisonAction(inspection_comparisonStep)
}

var inspection_comparisonStep = &Step{
	name: "Inspection and Comparison",
}

func createInspection_ComparisonAction(s *Step) func() {
	return func() {
		println("my Inspection and Comparison function goes here")
	}
}
