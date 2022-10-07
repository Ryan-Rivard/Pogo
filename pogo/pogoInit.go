package main

import "fmt"

type pogoInit struct {
	git iGit
}

func (i *pogoInit) process() {
	fmt.Println("initializing git repo")

	i.git.executeGitCommand("init")
}
