package main

import (
	"log"

	"github.com/AlecAivazis/survey/v2"
)

type inquierService struct {
}

func (i *inquierService) askWithPresetAnswers(name string, message string, options []string) string {
	question := i.CreateQuestion(name, message, options, options[0])
	return i.AskSelection(question)
}

func (i *inquierService) CreateQuestion(name string, message string, options []string, startingPoint string) []*survey.Question {
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

func (i *inquierService) AskSelection(question []*survey.Question) string {
	answers := struct {
		Answer string
	}{}

	err := survey.Ask(question, &answers)

	if err != nil {
		log.Panic(err.Error())
	}

	return answers.Answer
}

// func AskSelection(question []*survey.Question) string {
// 	answers := struct {
// 		Answer string
// 	}{}

// 	err := survey.Ask(question, &answers)

// 	if err != nil {
// 		log.Panic(err.Error())
// 	}

// 	return answers.Answer
// }

// func AskMultiSelect(question *survey.MultiSelect) []string {
// 	selection := []string{}
// 	survey.AskOne(question, &selection)
// 	return selection
// }

// func AskConfirm(question *survey.Confirm) bool {
// 	name := false
// 	survey.AskOne(question, &name)
// 	return name
// }

// func CreateQuestion(name string, message string, options []string, startingPoint string) []*survey.Question {
// 	return []*survey.Question{
// 		{
// 			Name: name,
// 			Prompt: &survey.Select{
// 				Message: message,
// 				Options: options,
// 				Default: startingPoint,
// 			},
// 		},
// 	}
// }

// func CreateMultiSelect(message string, options []string) *survey.MultiSelect {
// 	return &survey.MultiSelect{
// 		Message: message,
// 		Options: options,
// 	}
// }

// func CreateConfirm(message string) *survey.Confirm {
// 	return &survey.Confirm{
// 		Default: true,
// 		Message: message,
// 	}
// }
