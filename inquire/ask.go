package inquier

import (
	"github.com/AlecAivazis/survey/v2"
	"log"
)

func AskWithOptions(message string, options []string) *string {
	question := []*survey.Question{
		{
			Name: "Answer",
			Prompt: &survey.Select{
				Message: message,
				Options: options,
				Default: options[0],
			},
		},
	}

	answers := struct {
		Answer string
	}{}

	err := survey.Ask(question, &answers)

	if err != nil {
		log.Panic(err.Error())
	}

	return &answers.Answer
}
