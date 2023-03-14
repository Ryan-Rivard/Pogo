package inquier

import (
	"log"

	"github.com/AlecAivazis/survey/v2"
)

func AskPrompt(message string) string {
	ans := ""
	prompt := &survey.Input{
		Message: message,
	}

	survey.AskOne(prompt, &ans)
	return ans
}

func AskWithOptions(message string, options []string) string {
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

	return answers.Answer
}

func AskMultiSelect(message string, options []string) []string {
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
