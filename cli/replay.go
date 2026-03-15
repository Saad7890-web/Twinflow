package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var captureDir string
var replayTarget string

var replayCmd = &cobra.Command{
	Use:   "replay",
	Short: "Replay captured traffic",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Replaying traffic from:", captureDir)
		fmt.Println("Target service:", replayTarget)
	},
}

func init() {
	replayCmd.Flags().StringVar(&captureDir, "capture", "./captures", "capture directory")
	replayCmd.Flags().StringVar(&replayTarget, "target", "", "target service URL")

	replayCmd.MarkFlagRequired("target")
}