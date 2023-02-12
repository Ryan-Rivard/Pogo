package cmd

import (
	"github.com/Ryan-Rivard/Pogo/wizard"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(checkoutBranchCmd)
}

var checkoutBranchCmd = &cobra.Command{
	Use:   "checkout-branch",
	Short: "Checkout branch",
	Long:  "List out the local and remote branches and confirm branch to checkout",
	Run: func(cmd *cobra.Command, args []string) {
		tree := wizard.BuildCheckoutBranchComposite()
		tree.Exec(nil)
	},
}
