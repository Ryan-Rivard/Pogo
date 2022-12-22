package wizard

// Builder Pattern

func BuildRootComposite() Step {
	return &Ask{
		Id:       "Root",
		Question: "What would you like to do today?",
		Components: []Step{
			BuildSnapshottingComposite(),
		},
	}
}

func BuildSnapshottingComposite() Step {
	return &Ask{
		Id:       "Basic Snapshotting",
		Question: "What git comand would you like to run?",
		Components: []Step{
			&Cmd{
				Id:  "status",
				arg: "status",
			},
		},
	}
}
