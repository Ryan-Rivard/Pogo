package main

import "fmt"

type pogoStandalone struct {
	inquier iInquier
}

func (s *pogoStandalone) process() {
	options := []string{"init", "log", "switch", "branch", "add", "diff", "status", "commit", "restore", "merge", "rebase", "tag", "quit"}

	answer := s.inquier.askWithPresetAnswers("Answer", "Standalone - What can pogo do for you today?", options)

	if *answer == "quit" {
		fmt.Println("Thank you for using pogo")
	}

	git := &gitService{}

	if *answer == "init" {
		processInit := &pogoInit{
			git: git,
		}

		processInit.process()
	}

	if *answer == "log" {
		processLog := &pogoLog{
			inquier: s.inquier,
			git:     git,
		}

		processLog.process()
	}

	if *answer == "switch" {
		processSwitch := &pogoSwitch{
			git: git,
		}

		processSwitch.process()
	}

	if *answer == "branch" {
		processBranch := &pogoBranch{
			git: git,
		}

		processBranch.process()
	}

	if *answer == "add" {
		processAdd := &pogoAdd{
			git: git,
		}

		processAdd.process()
	}

	if *answer == "diff" {
		processDiff := &pogoDiff{
			git: git,
		}

		processDiff.process()
	}

	if *answer == "status" {
		processStatus := &pogoStatus{
			inquier: s.inquier,
			git:     git,
		}

		processStatus.process()
	}

	if *answer == "commit" {
		processCommit := &pogoCommit{
			git: git,
		}

		processCommit.process()
	}

	if *answer == "restore" {
		processRestore := &pogoRestore{
			git: git,
		}

		processRestore.process()
	}

	if *answer == "merge" {
		processMerge := &pogoMerge{
			git: git,
		}

		processMerge.process()
	}

	if *answer == "rebase" {
		processRebase := &pogoRebase{
			git: git,
		}

		processRebase.process()
	}

	if *answer == "tag" {
		processTag := &pogoTag{
			git: git,
		}

		processTag.process()
	}
}
