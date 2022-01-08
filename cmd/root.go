package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd *cobra.Command = &cobra.Command{Use: "gohelper"}

func init() {
	initCmdInit()
	rootCmd.AddCommand(cmdInit)
}

func Execute() error {
	return rootCmd.Execute()
}
