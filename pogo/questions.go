package main

import (
	"log"

	"github.com/AlecAivazis/survey/v2"
)

func AskSelection(question []*survey.Question) string {
	answers := struct {
		Answer string
	}{}

	err := survey.Ask(question, &answers)

	if err != nil {
		log.Panic(err.Error())
	}

	return answers.Answer
}

func CreateQuestion(name string, message string, options []string, startingPoint string) []*survey.Question {
	return []*survey.Question{
		{
			Name: name,
			Prompt: &survey.Select{
				Message: message,
				Options: options,
				Default: startingPoint,
			},
		},
	}
}
