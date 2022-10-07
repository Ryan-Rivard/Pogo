package main

type pogoRestore struct {
	git iGit
}

func (r *pogoRestore) process() {
	r.git.executeGitCommand("restore")
}
