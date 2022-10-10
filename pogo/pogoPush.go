package main

type pogoPush struct {
	git iGit
}

func (p *pogoPush) process() {
	p.git.executeGitCommand("push")
}
