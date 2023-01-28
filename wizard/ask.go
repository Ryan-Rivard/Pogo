package wizard

import (
	"log"

	"github.com/AlecAivazis/survey/v2"
)

type ask struct {
	id           string
	question     string
	questionType string
	steps        []step
}

func (a *ask) Exec(params []string) {
	if a.questionType == "single" {
		answer := askWithOptions(a.question, params)
		// figure out which step to send it to
		println(answer)

	} else if a.questionType == "multi" {
		answers := askMultiSelect(a.question, params)
		// always going to be one step
		a.steps[0].Exec(answers)
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
