package main

import "fmt"

type pogoStatus struct {
	inquier iInquier
	git     iGit
}

func (s *pogoStatus) process() {
	fmt.Println("outputting status")

	s.git.executeGitCommand("status")

}
