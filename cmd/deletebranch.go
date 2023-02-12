package cmd

import (
	"github.com/Ryan-Rivard/Pogo/wizard"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(deleteBranchCmd)
}

var deleteBranchCmd = &cobra.Command{
	Use:   "delete-branch",
	Short: "Delete local branch(es)",
	Long:  "List out the local branches and confirm any to delete",
	Run: func(cmd *cobra.Command, args []string) {
		tree := wizard.BuildDeleteBranchComposite()
		tree.Exec(nil)
	},
}
