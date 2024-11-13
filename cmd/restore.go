package cmd

import (
	"fmt"
	"github.com/ezrantn/macshift/adapter"
	"github.com/spf13/cobra"
	"log"
)

var restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restore the original MAC address",
	Run: func(cmd *cobra.Command, args []string) {
		err := adapter.RestoreOriginalMAC(iface)
		if err != nil {
			log.Fatalf("Failed to restore original MAC address: %v", err)
		}

		fmt.Printf("Original MAC address restored successfully on interface %s\n", iface)
	},
}

func init() {
	restoreCmd.Flags().StringVarP(&iface, "interface", "i", "", "Network interface to restore MAC address (required)")
	restoreCmd.MarkFlagRequired("interface")
}
