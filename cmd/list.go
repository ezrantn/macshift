package cmd

import (
	"fmt"
	"github.com/ezrantn/macshift/adapter"
	"github.com/spf13/cobra"
	"log"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all network adapters",
	Run: func(cmd *cobra.Command, args []string) {
		adapters, err := adapter.ListAdapters()
		if err != nil {
			log.Fatalf("Failed to list adapters: %v", err)
		}

		fmt.Println("\nAvailable network adapters:")
		for _, adapter := range adapters {
			fmt.Printf("\nName: %s\nDescription: %s\nMAC: %s\n\n",
				adapter.Name, adapter.Description, adapter.MacAddress)
		}
	},
}
