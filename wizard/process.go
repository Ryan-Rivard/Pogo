package wizard

import (
// "github.com/Ryan-Rivard/Pogo/inquire"
)

func ProcessNoArgs() {
	// do something here

	// options := []string{
	// 	"init",
	// 	"clone",
	// 	"branch",
	// 	"fetch",
	// 	"pull",
	// 	"status",
	// 	"add",
	// 	"commit",
	// 	"push",
	// }

	// options := []string{
	// 	"status",
	// 	"version",
	// }

	// myStep := &Step{
	// 	prev:    nil,
	// 	options: options,
	// }

	// search := inquier.AskWithOptions("What git command would you like to execute?", options)

	// println("you selected", *search)

	// wiz := createTree()

	// ans := wiz.doer()
	// wiz.

	rootStep.execute()

}
