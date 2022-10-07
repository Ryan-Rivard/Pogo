package main

type pogoTag struct {
	git iGit
}

func (t *pogoTag) process() {
	t.git.executeGitCommand("tag")
}
