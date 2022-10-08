package main

import "fmt"

type pogoAdd struct {
	git iGit
}

func (a *pogoAdd) process() {
	// 	changedFiles := Execute("git", "ls-files", "-m", "-d", "-o", "--exclude-standard")

	fmt.Println("adding all files in directory to stage")

	a.git.executeGitCommand("add", ".")
}
