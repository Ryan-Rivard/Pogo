package main

type pogoPull struct {
	git iGit
}

func (p *pogoPull) process() {
	p.git.executeGitCommand("pull")
}
