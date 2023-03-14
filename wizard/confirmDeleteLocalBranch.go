package wizard

import (
	"log"

	inquier "github.com/Ryan-Rivard/Pogo/inquire"
)

type ConfirmDeleteLocalBranch struct {
	step step
}

func (c *ConfirmDeleteLocalBranch) Exec(input interface{}) {
	b, ok := input.([]branch)
	if !ok {
		log.Fatal("boom1")
	}

	var Reset = "\033[0m"
	// var Red = "\033[31m"
	var Green = "\033[32m"
	var Yellow = "\033[33m"
	// var Blue = "\033[34m"
	// var Purple = "\033[35m"
	// var Cyan = "\033[36m"
	// var Gray = "\033[37m"
	// var White = "\033[97m"

	branchList := []string{}
	for _, branch := range b {
		branchList = append(branchList,
			Yellow+branch.commit.refname+Reset+" \u0000 "+
				branch.commit.authorName+" \u0000 "+
				Green+branch.commit.relativeCommitDate+Reset)
	}

	answers := inquier.AskMultiSelect("Which local branches would you like to delete?", branchList)

	if len(answers) == 0 {
		println("no selection - exiting")
		return
	}

	c.step.Exec(answers)
}
