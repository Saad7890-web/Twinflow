package cli

import (
	"fmt"

	"github.com/Saad7890-web/twinflow/internal/proxy"
	"github.com/spf13/cobra"
)

var listenAddr string
var targetAddr string

var recordCmd = &cobra.Command{
	Use:   "record",
	Short: "Record HTTP traffic",
	Run: func(cmd *cobra.Command, args []string) {

	fmt.Println("Starting TwinFlow recorder")

	err := proxy.Start(listenAddr, targetAddr)
	if err != nil {
		fmt.Println("Error:", err)
	}
},
}

func init() {
	recordCmd.Flags().StringVar(&listenAddr, "listen", ":8080", "proxy listen address")
	recordCmd.Flags().StringVar(&targetAddr, "target", "", "target service URL")

	recordCmd.MarkFlagRequired("target")
}
