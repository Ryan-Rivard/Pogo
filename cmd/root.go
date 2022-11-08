package cmd

import (
	"fmt"
	"github.com/Ryan-Rivard/Pogo/inquire"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "pogo",
	Short: "Pogo is a git cli wiard",
	Long:  "What more do you need?",
	Run: func(cmd *cobra.Command, args []string) {
		options := []string{"init", "clone", "branch", "fetch", "pull", "status", "add", "commit", "push"}
		search := inquier.AskWithOptions("What can pogo do for you?", options)

		println("you selected", *search)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
