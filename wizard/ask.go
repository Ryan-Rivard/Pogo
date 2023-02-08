package wizard

import (
	"log"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/acarl005/stripansi"
)

type ask struct {
	id            string
	question      string
	questionType  string
	convertInput  func(interface{}) []string
	convertOutput func([]string) interface{}
	steps         []step
}

func (a *ask) Exec(input interface{}) {
	params := []string{}
	if a.convertInput != nil {
		params = a.convertInput(input)
	}

	if a.questionType == "single" {
		answer := askWithOptions(a.question, params)
		// figure out which step to send it to
		println(answer)

	} else if a.questionType == "multi" {
		answers := askMultiSelect(a.question, params)

		if len(answers) == 0 {
			println("no selection - exiting")
			return
		}
		// always going to be one step
		a.steps[0].Exec(a.convertOutput(answers))
	}
}

func askWithOptions(question string, options []string) string {
	prompt := []*survey.Question{
		{
			Name: "Answer",
			Prompt: &survey.Select{
				Message: question,
				Options: options,
			},
		},
	}

	answer := struct{ Answer string }{}
	err := survey.Ask(prompt, &answer)

	if err != nil {
		log.Panic(err.Error())
	}

	return answer.Answer
}

func askMultiSelect(message string, options []string) []string {
	prompt := &survey.MultiSelect{
		Message: message,
		Options: options,
	}

	answers := []string{}

	err := survey.AskOne(prompt, &answers)

	if err != nil {
		log.Panic(err.Error())
	}

	return answers
}

func importBranchListOptions(input interface{}) []string {
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
			Yellow+branch.refname+Reset+" \u0000 "+branch.authorName+" \u0000 "+Green+branch.relativeCommitDate+Reset)
	}

	return branchList
}

func exportBranchChoice(answers []string) interface{} {
	output := []string{}

	for _, answer := range answers {
		strip := stripansi.Strip(answer)
		split := strings.Split(strip, "\u0000")
		output = append(output, strings.TrimSpace(split[0]))
	}

	return output
}
