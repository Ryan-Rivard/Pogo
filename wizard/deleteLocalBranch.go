package wizard

import (
	"log"

	"github.com/Ryan-Rivard/Pogo/git"
)

type DeleteLocalBranch struct {
}

func (l *DeleteLocalBranch) Exec(input interface{}) {
	b, ok := input.([]string)
	if !ok {
		log.Fatal("boomie")
	}
	args := []string{"branch", "-d"}

	args = append(args, b...)

	git.ExecuteGitCommand(args...)
}
