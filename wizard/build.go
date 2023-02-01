package wizard

func BuildDeleteBranchComposite() step {
	return &cmd{
		id:            "list-branch",
		args:          []string{"for-each-ref", "--sort=committerdate", "refs/heads/", "--format=%(refname:short)|%(committerdate:relative)"},
		convertOutput: exportBranchList,
		step: &ask{
			id:            "confirm-delete",
			question:      "Which Branches would you like to delete?",
			questionType:  "multi",
			convertInput:  importBranchListOptions,
			convertOutput: exportBranchChoice,
			steps: []step{
				&cmd{
					id:           "delete-branch",
					args:         []string{"branch", "-d"},
					convertInput: importStringSlice,
				},
			},
		},
	}
}
