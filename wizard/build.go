package wizard

func BuildDeleteBranchComposite() step {
	return &cmd{
		id:            "list-local-branch",
		args:          []string{"for-each-ref", "--sort=committerdate", "refs/heads/", "--format=%(refname:short)|%(authorname)|%(committerdate:relative)"},
		convertOutput: exportLocalBranchList,
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

func BuildCheckoutBranchComposite() step {
	return &cmd{
		id:   "fetch-branches",
		args: []string{"fetch"},
		step: &cmd{
			id:            "list-all-branch",
			args:          []string{"for-each-ref", "refs/heads/", "refs/remotes/origin/", "--format=%(refname:short)|%(authorname)|%(committerdate:relative)"},
			convertOutput: exportBranchList,
			step: &ask{
				id:           "confirm-checkout",
				question:     "Which branch would you like to checkout?",
				questionType: "single",
				steps: []step{
					&cmd{
						id:   "checkout-branch",
						args: []string{"checkout"},
					},
				},
			},
		},
	}
}
