package cmd

import (
	"fmt"
	"github.com/ezrantn/macshift/adapter"
	"github.com/spf13/cobra"
	"log"
)

var changeCmd = &cobra.Command{
	Use:   "change",
	Short: "Generate and set a random MAC address",
	Run: func(cmd *cobra.Command, args []string) {
		newMac, err := adapter.GenerateMac()
		if err != nil {
			log.Fatalf("Failed to generate random MAC address: %v", err)
		}

		fmt.Println("Generated random MAC address:", newMac)

		err = adapter.ChangeMACAddress(iface, newMac)
		if err != nil {
			log.Fatalf("Failed to change MAC address: %v", err)
		}

		fmt.Printf("MAC address changed successfully to %s on interface %s\n", newMac, iface)
	},
}

func init() {
	changeCmd.Flags().StringVarP(&iface, "interface", "i", "", "Network interface to change MAC address (required)")
	changeCmd.MarkFlagRequired("interface")
}
