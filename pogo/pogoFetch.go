package main

type pogoFetch struct {
	git iGit
}

func (f *pogoFetch) process() {
	f.git.executeGitCommand("fetch")
}
