package cmd

import (
	"fmt"
	"github.com/Ryan-Rivard/Pogo/wizard"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "pogo",
	Short: "Pogo is a git cli wiard",
	Long:  "What more do you need?",
	Run: func(cmd *cobra.Command, args []string) {
		wizard.ProcessNoArgs()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
