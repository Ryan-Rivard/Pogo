package cmd

import (
	"fmt"
	"os"

	"github.com/Ryan-Rivard/Pogo/wizard"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pogo",
	Short: "Pogo is a git cli wiard",
	Long:  "What more do you need?",
	Run: func(cmd *cobra.Command, args []string) {
		println("root")
		tree := wizard.BuildDeleteBranchComposite()
		tree.Exec(nil)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
