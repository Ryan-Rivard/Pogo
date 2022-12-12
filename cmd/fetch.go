package cmd

import (
	"github.com/Ryan-Rivard/Pogo/wizard"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(fetchCmd)
}

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "fetch the most recent changes",
	Long:  `fetch the mostmostmost recent changes`,
	Run: func(cmd *cobra.Command, args []string) {
		//wizard.Process("fetch")
		tree := wizard.BuildComposite()
		tree.Execute()
	},
}
