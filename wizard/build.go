package wizard

func BuildDeleteBranchComposite() step {
	return &cmd{
		id:           "list-branch",
		args:         []string{"for-each-ref", "--sort=committerdate", "refs/heads/", "--format=%(refname:short)"}, // "--format='%(HEAD) %(refname:short) -  %(contents:subject) -  %(authorname) %(committerdate:relative)"},
		formatOutput: branchListFormat,
		step: &ask{
			id:           "confirm-delete",
			question:     "Which Branches would you like to delete?",
			questionType: "multi",
			steps: []step{
				&cmd{
					id:   "delete-branch",
					args: []string{"branch", "-d"},
				},
			},
		},
	}
}
