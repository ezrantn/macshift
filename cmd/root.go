package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var (
	iface   string
	rootCmd = &cobra.Command{
		Use:   "macshift",
		Short: "MAC Address Changing Tool",
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(changeCmd, restoreCmd, listCmd, versionCmd)
}
