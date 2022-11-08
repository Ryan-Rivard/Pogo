package cmd

import (
	"github.com/Ryan-Rivard/Pogo/inquire"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(searchCmd)
}

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Give all options that pogo can do",
	Long:  "All git functions that pogo can provide to you",
	Run: func(cmd *cobra.Command, args []string) {
		options := []string{"init", "clone", "branch", "fetch", "pull", "status", "add", "commit", "push"}

		search := inquier.AskWithOptions("search all", options)

		println("you selected", *search)
	},
}
