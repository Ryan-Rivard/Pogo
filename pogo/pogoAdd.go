package main

type pogoAdd struct {
	git iGit
}

func (a *pogoAdd) process() {
	// 	changedFiles := Execute("git", "ls-files", "-m", "-d", "-o", "--exclude-standard")

	a.git.executeGitCommand("add")
}
