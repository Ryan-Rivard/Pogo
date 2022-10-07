package main

type pogoBranch struct {
	git iGit
}

func (b *pogoBranch) process() {
	b.git.executeGitCommand("branch")
}
