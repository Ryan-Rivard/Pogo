package cmd

import (
	// "github.com/Ryan-Rivard/Pogo/git"
	// "github.com/Ryan-Rivard/Pogo/inquire"
	"github.com/Ryan-Rivard/Pogo/wizard"
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
		tree := wizard.BuildStandaloneComposite()
		tree.Execute()
	},
}
