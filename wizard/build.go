package wizard

// Builder Pattern

func BuildComposite() Step {
	return &Ask{
		Id:       "Root",
		Question: "What would you like to do today?",
		Components: []Step{
			BuildSnapshottingBranch(),
		},
	}
}

func BuildSnapshottingBranch() Step {
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
