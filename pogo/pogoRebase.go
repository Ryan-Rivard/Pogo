package main

type pogoRebase struct {
	git iGit
}

func (r *pogoRebase) process() {
	r.git.executeGitCommand("rebase")
}
