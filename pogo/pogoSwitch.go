package main

type pogoSwitch struct {
	git iGit
}

func (s *pogoSwitch) process() {
	s.git.executeGitCommand("switch")
}
