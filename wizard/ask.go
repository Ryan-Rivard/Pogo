package wizard

import (
	"log"

	"github.com/AlecAivazis/survey/v2"
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

	answer := struct {
		Answer string
	}{}

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
	b, ok := input.(*branches)
	if !ok {
		log.Fatal("boom1")
	}

	branchList := []string{}
	for _, branch := range b.branches {
		branchList = append(branchList, branch.refname)
	}

	return branchList
}
