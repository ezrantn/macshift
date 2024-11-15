package cmd

import (
	"fmt"
	"github.com/briandowns/spinner"
	"github.com/ezrantn/macshift/adapter"
	"github.com/spf13/cobra"
	"os"
	"time"
)

var restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restore the original MAC address",
	Run: func(cmd *cobra.Command, args []string) {
		s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
		s.Writer = os.Stdout

		s.Start()

		err := adapter.RestoreOriginalMAC(iface)
		if err != nil {
			s.Stop()
			fmt.Fprintf(os.Stderr, "Error: Failed to restore the original MAC address on interface %s. %v\n", iface, err)
			fmt.Fprintln(os.Stderr, "Check if the original MAC address is available and if you have sufficient permissions.")
			os.Exit(1)
		}

		s.Stop()
		fmt.Print("\r")

		fmt.Printf("Original MAC address restored successfully on interface %s\n", iface)
	},
}

func init() {
	restoreCmd.Flags().StringVarP(&iface, "interface", "i", "", "Network interface to restore MAC address (required)")
	restoreCmd.MarkFlagRequired("interface")
}
