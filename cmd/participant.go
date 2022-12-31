package cmd

import (
	// "github.com/Ryan-Rivard/Pogo/git"
	// "github.com/Ryan-Rivard/Pogo/inquire"
	"github.com/Ryan-Rivard/Pogo/wizard"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(participantCmd)
}

var participantCmd = &cobra.Command{
	Use:   "participant",
	Short: "Basic git commands for the participant individual developer",
	Long:  `https://git-scm.com/docs/everyday#_individual_developer_standaloneindividual_developer_standalone`,
	Run: func(cmd *cobra.Command, args []string) {
		tree := wizard.BuildBasicComposite()
		tree.Execute()
	},
}
