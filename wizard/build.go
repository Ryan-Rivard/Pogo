package wizard

// Builder Pattern

func BuildComposite() Step {
	return &Ask{
		Question: "Question",
		Components: []Step{
			&Cmd{
				arg: "status",
			},
		},
	}
}
