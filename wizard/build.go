package wizard

// Builder Pattern?

func BuildStandaloneComposite() Step {
	return &Ask{
		Id:       "Root",
		Question: "Standalone deverloper: What would you like to do today?",
		Components: []Step{
			&Cmd{
				Id:  "(init) Create an empty Git repository or reinitialize an existing one",
				arg: []string{"init"},
			},
			&Cmd{
				Id:  "(show-branch) Show branches and their commits",
				arg: []string{"show-branch"},
			},
			&Cmd{
				Id:  "(log) Show commit logs",
				arg: []string{"log"},
			},
			&Cmd{
				Id:  "(checkout) Switch branches or restore working tree files",
				arg: []string{"checkout"},
			},
			&Cmd{
				Id:  "(branch) List, create, or delete branches",
				arg: []string{"branch"},
			},
			&Cmd{
				Id:  "(add) Add file contents to the index",
				arg: []string{"add"},
			},
			&Cmd{
				Id:  "(diff) Show changes between commits, commit and working tree, etc",
				arg: []string{"diff"},
			},
			&Cmd{
				Id:  "(status) Show the working tree status",
				arg: []string{"status"},
			},
			&Cmd{
				Id:  "(commit) Record changes to the repository",
				arg: []string{"commit"},
			},
			&Cmd{
				Id:  "(reset) Reset current HEAD to the specified state",
				arg: []string{"reset"},
			},
			&Cmd{
				Id:  "(checkout) Switch branches or restore working tree files",
				arg: []string{"checkout"},
			},
			&Cmd{
				Id:  "(merge) Join two or more development histories together",
				arg: []string{"merge"},
			},
			&Cmd{
				Id:  "(rebase) Reapply commits on top of another base tip",
				arg: []string{"rebase"},
			},
			&Cmd{
				Id:  "(tag) Create, list, delete or verify a tag object signed with GPG",
				arg: []string{"tag"},
			},
		},
	}
}

func BuildBasicComposite() Step {
	return &Ask{
		Id:       "Root",
		Question: "Basic Commands: What would you like to do today?",
		Components: []Step{
			&Cmd{
				Id:  "(init) Create an empty Git repository or reinitialize an existing one",
				arg: []string{"init"},
			},
			&Cmd{
				Id:  "(clone) Clone a repository into a new directory",
				arg: []string{"init"},
			},
			&Cmd{
				Id:  "(show-branch) Show branches and their commits",
				arg: []string{"show-branch"},
			},
			&Cmd{
				Id:  "(log) Show commit logs",
				arg: []string{"log"},
			},
			&Cmd{
				Id:  "(checkout) Switch branches or restore working tree files",
				arg: []string{"checkout"},
			},
			&Cmd{
				Id:  "(branch) List, create, or delete branches",
				arg: []string{"branch"},
			},
			&Cmd{
				Id:  "(pull) Fetch from and integrate with another repository or a local branch",
				arg: []string{"branch"},
			},
			&Cmd{
				Id:  "(fetch) Download objects and refs from another repository",
				arg: []string{"branch"},
			},
			&Cmd{
				Id:  "(add) Add file contents to the index",
				arg: []string{"add"},
			},
			&Cmd{
				Id:  "(diff) Show changes between commits, commit and working tree, etc",
				arg: []string{"diff"},
			},
			&Cmd{
				Id:  "(status) Show the working tree status",
				arg: []string{"status"},
			},
			&Cmd{
				Id:  "(commit) Record changes to the repository",
				arg: []string{"commit"},
			},
			&Cmd{
				Id:  "(push) Update remote refs along with associated objects",
				arg: []string{"commit"},
			},
			&Cmd{
				Id:  "(reset) Reset current HEAD to the specified state",
				arg: []string{"reset"},
			},
			&Cmd{
				Id:  "(checkout) Switch branches or restore working tree files",
				arg: []string{"checkout"},
			},
			&Cmd{
				Id:  "(merge) Join two or more development histories together",
				arg: []string{"merge"},
			},
			&Cmd{
				Id:  "(rebase) Reapply commits on top of another base tip",
				arg: []string{"rebase"},
			},
			&Cmd{
				Id:  "(tag) Create, list, delete or verify a tag object signed with GPG",
				arg: []string{"tag"},
			},
			&Cmd{
				Id:  "(pull) Fetch from and integrate with another repository or a local branch",
				arg: []string{"tag"},
			},
			&Cmd{
				Id:  "(revert) Revert some existing commits",
				arg: []string{"tag"},
			},
		},
	}
}

func BuildRootComposite() Step {
	return &Ask{
		Id:       "Root",
		Question: "What would you like to do today?",
		Components: []Step{
			BuildStandaloneComposite(), // change later
		},
	}
}

// func BuildSnapshottingComposite() Step {
// 	return &Ask{
// 		Id:       "Basic Snapshotting",
// 		Question: "What git comand would you like to run?",
// 		Components: []Step{
// 			&Cmd{
// 				Id:  "(add) Add file contents to the index",
// 				arg: []string{"add", "."},
// 			},
// 			&Cmd{
// 				Id:  "(status) Show the working tree status",
// 				arg: []string{"status"},
// 			},
// 			&Cmd{
// 				Id:  "(diff) Show changes between commits, commit and working tree, etc",
// 				arg: []string{"diff"},
// 			},
// 			&Cmd{
// 				Id:  "(commit) Record changes to the repository",
// 				arg: []string{"commit", "."},
// 			},
// 			&Cmd{
// 				Id:  "(notes) Add or inspect object notes",
// 				arg: []string{"notes"},
// 			},
// 			&Cmd{
// 				Id:  "(restore) Restore working tree files",
// 				arg: []string{"restore", "."},
// 			},
// 			&Cmd{
// 				Id:  "(reset) Reset current HEAD to the specified state",
// 				arg: []string{"reset"},
// 			},
// 			&Cmd{
// 				Id:  "(rm) Remove files from the working tree and from the index",
// 				arg: []string{"rm"},
// 			},
// 			&Cmd{
// 				Id:  "(mv) Move or rename a file, a directory, or a symlink",
// 				arg: []string{"mv"},
// 			},
// 		},
// 	}
// }
