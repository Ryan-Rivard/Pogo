package main

type pogoDiff struct {
	git iGit
}

func (d *pogoDiff) process() {
	d.git.executeGitCommand("diff")
}
