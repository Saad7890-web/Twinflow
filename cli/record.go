package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listenAddr string
var targetAddr string

var recordCmd = &cobra.Command{
	Use:   "record",
	Short: "Record HTTP traffic",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Record mode started")
		fmt.Println("Listening on:", listenAddr)
		fmt.Println("Forwarding to:", targetAddr)
	},
}

func init() {
	recordCmd.Flags().StringVar(&listenAddr, "listen", ":8080", "proxy listen address")
	recordCmd.Flags().StringVar(&targetAddr, "target", "", "target service URL")

	recordCmd.MarkFlagRequired("target")
}
