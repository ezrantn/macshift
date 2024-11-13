package cmd

import (
	"github.com/spf13/cobra"
)

var (
	iface   string
	rootCmd = &cobra.Command{
		Use:   "macshift",
		Short: "MAC Address Changing Tool",
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(changeCmd, restoreCmd, listCmd, versionCmd)
}
