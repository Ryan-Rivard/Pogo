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
				Id:  "(add) Add file contents to the index",
				arg: []string{"add", "."},
			},
			&Cmd{
				Id:  "(status) Show the working tree status",
				arg: []string{"status"},
			},
			&Cmd{
				Id:  "(diff) Show changes between commits, commit and working tree, etc",
				arg: []string{"diff"},
			},
			&Cmd{
				Id:  "(commit) Record changes to the repository",
				arg: []string{"commit", "."},
			},
			&Cmd{
				Id:  "(notes) Add or inspect object notes",
				arg: []string{"notes"},
			},
			&Cmd{
				Id:  "(restore) Restore working tree files",
				arg: []string{"restore", "."},
			},
			&Cmd{
				Id:  "(reset) Reset current HEAD to the specified state",
				arg: []string{"reset"},
			},
			&Cmd{
				Id:  "(rm) Remove files from the working tree and from the index",
				arg: []string{"rm"},
			},
			&Cmd{
				Id:  "(mv) Move or rename a file, a directory, or a symlink",
				arg: []string{"mv"},
			},
		},
	}
}
