package main

type pogoCommit struct {
	git iGit
}

func (c *pogoCommit) process() {
	c.git.executeGitCommand("commit")
}
