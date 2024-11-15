package cmd

import (
	"fmt"
	"github.com/briandowns/spinner"
	"github.com/ezrantn/macshift/adapter"
	"github.com/spf13/cobra"
	"os"
	"time"
)

var changeCmd = &cobra.Command{
	Use:   "change",
	Short: "Generate and set a random MAC address",
	Run: func(cmd *cobra.Command, args []string) {
		s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
		s.Writer = os.Stdout

		newMac, err := adapter.GenerateMac()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: Failed to generate random MAC address. %v\n", err)
			os.Exit(1)
		}

		fmt.Println("Generated random MAC address:", newMac)

		s.Start()

		err = adapter.ChangeMACAddress(iface, newMac)
		if err != nil {
			s.Stop()
			fmt.Fprintf(os.Stderr, "Error: Failed to change MAC address on interface %s. %v\n", iface, err)
			fmt.Fprintln(os.Stderr, "Ensure you have the required permissions or try running the program as administrator.")
			os.Exit(1)
		}

		s.Stop()
		fmt.Print("\r")

		fmt.Printf("MAC address changed successfully to %s on interface %s\n", newMac, iface)
	},
}

func init() {
	changeCmd.Flags().StringVarP(&iface, "interface", "i", "", "Network interface to change MAC address (required)")
	changeCmd.MarkFlagRequired("interface")
}
