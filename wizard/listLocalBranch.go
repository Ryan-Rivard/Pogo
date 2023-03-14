package wizard

import (
	"strings"

	"github.com/Ryan-Rivard/Pogo/git"
)

type ListLocalBranch struct {
	step step
}

func (l *ListLocalBranch) Exec(input interface{}) {
	args := []string{"for-each-ref", "--sort=committerdate", "refs/heads/", "--format=%(refname:short)|%(authorname)|%(committerdate:relative)"}

	output := git.ExecuteGitCommand(args...)

	slice := strings.Split(strings.TrimSpace(output), "\n")
	b := []branch{}

	for _, line := range slice {
		split := strings.Split(line, "|")
		b = append(
			b, branch{
				commit: commit{
					refname: split[0], authorName: split[1], relativeCommitDate: split[2],
				},
			},
		)
	}

	l.step.Exec(b)
}
