package main

type iInquier interface {
	askWithPresetAnswers(name string, message string, options []string) string
}
