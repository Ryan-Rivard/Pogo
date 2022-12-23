package wizard

import (
	"log"

	"github.com/AlecAivazis/survey/v2"
)

// Composite pattern
// Container
// Element that contains sub-containers and/or leafs

type Ask struct {
	Id         string
	Question   string
	Answer     string
	Components []Step
}

func (a *Ask) GetId() string {
	return a.Id
}

func (a *Ask) Execute() {
	options := []string{}

	for _, com := range a.Components {
		options = append(options, com.GetId())
	}

	answer := askWithOptions(a.Question, options)

	for _, com := range a.Components {
		if com.GetId() == answer {
			com.Execute()
			break
		}
	}
}

// func askPrompt(message string) *string {
// 	ans := ""
// 	prompt := &survey.Input{
// 		Message: message,
// 	}

// 	survey.AskOne(prompt, &ans)
// 	return &ans
// }

func askWithOptions(message string, options []string) string {
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
