package cmd

import (
	"github.com/Ryan-Rivard/Pogo/git"
	"github.com/Ryan-Rivard/Pogo/inquire"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(standaloneCmd)
}

var standaloneCmd = &cobra.Command{
	Use:   "standalone",
	Short: "Basic git commands for the standalone individual developer",
	Long:  `https://git-scm.com/docs/everyday#_individual_developer_standaloneindividual_developer_standalone`,
	Run: func(cmd *cobra.Command, args []string) {
		options := []string{
			"init",
			"show branch",
			"log",
			"checkout",
			"branch",
			"add",
			"diff",
			"status",
			"commit",
			"reset",
			"checkout",
			"merge",
			"rebase",
			"tag"}

		search := inquier.AskWithOptions("Standalone Developer", options)
		println("you selected", *search)

		git.ProcessLogCommand()
	},
}
