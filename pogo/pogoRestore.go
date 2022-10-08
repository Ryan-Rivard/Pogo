package main

import "fmt"

type pogoRestore struct {
	git iGit
}

func (r *pogoRestore) process() {

	fmt.Println("removing all files in directory from stage")

	r.git.executeGitCommand("restore", "--staged", ".")
}
