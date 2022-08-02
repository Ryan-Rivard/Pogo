package main

import (
	"fmt"
	"log"

	"github.com/Ryan-Rivard/pogo/git"
)

// git add .
// git reset

func main() {

	options := []string{"add", "exit"}

	question := CreateQuestion("Answer", "What can pogo do for you today?", options, options[0])
	answer := AskSelection(question)

	if answer == "exit" {
		log.Println("Thank you for using pogo")
		return
	}

	if answer == "add" {
		changedFiles := git.ListChangedFiles()
		fmt.Println("Your modified files are...")

		for _, changedFile := range changedFiles {
			fmt.Println(changedFile)
		}

		allFilesQuestion := CreateConfirm("Would you like to stage all of your changes?")
		allFilesAnswer := AskConfirm(allFilesQuestion)

		if !allFilesAnswer {
			fileQuestion := CreateMultiSelect("What files would you like to stage?", changedFiles)
			fileAnswer := AskMultiSelect(fileQuestion)

			fmt.Println("You selected...")

			for _, file := range fileAnswer {
				fmt.Println(file)
			}
		}
	}

	// process := flag.String("Role", "standalone", "What development role do you have?")
	// flag.Parse()

	// var options []string

	// if *process == "standalone" {
	// 	options = []string{
	// 		//"init",
	// 		//"log",
	// 		//"switch",
	// 		//"branch",
	// 		"add",
	// 		//"diff",//
	// 		"status",
	// 		"commit",
	// 		"restore",
	// 		"merge",
	// 		"rebase",
	// 		"tag",
	// 		"exit",
	// 	}
	// }

	// // options = []string{"Setup and Config",
	// // 	"Getting and Creating Projects",
	// // 	"Basic Snapshotting",
	// // 	"Branching and Merging",
	// // 	"Sharing and Updating Projects",
	// // 	"Inspection and Comparison",
	// // 	"Patching",
	// // 	"Debugging",
	// // }

	// question := CreateQuestion("Answer", "What would you like to do today?", options, options[0])

	// answer := AskSelection(question)

	// if answer == "exit" {
	// 	log.Println("Thank you for using pogo")
	// 	return
	// }

	// fmt.Printf("executing command: %s.\n", answer)

	// var answers []string
	// //--color --graph --pretty=format:'%Cred%h%Creset %Cgreen(%cr) %C(bold blue)<%an>%Creset -%C(yellow)%d%Creset %s' --abbrev-commit
	// if answer == "log" {
	// 	answers = append(answers, "log")
	// 	answers = append(answers, "")
	// }

	// Execute(answer)

	// // ExecGitCmd(answers.GitCommand)
	// // Execute(answers.GitCommand)

	// myConfig.TestMe()
	// // ExecGitCmd("add", ".") asdf

}
