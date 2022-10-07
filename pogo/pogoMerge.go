package main

type pogoMerge struct {
	git iGit
}

func (m *pogoMerge) process() {
	m.git.executeGitCommand("merge")
}
