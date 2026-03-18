package cli

import (
	"fmt"

	"github.com/Saad7890-web/twinflow/internal/replay"
	"github.com/spf13/cobra"
)

var captureDir string
var replayTarget string

var replayCmd = &cobra.Command{
	Use:   "replay",
	Short: "Replay captured traffic",
Run: func(cmd *cobra.Command, args []string) {

	engine := replay.NewEngine(replayTarget)

	err := engine.Replay(captureDir)
	if err != nil {
		fmt.Println("Replay failed:", err)
	}
},
}

func init() {
	replayCmd.Flags().StringVar(&captureDir, "capture", "/captures/traffic.log", "capture file")
	replayCmd.Flags().StringVar(&replayTarget, "target", "", "target service URL")

	replayCmd.MarkFlagRequired("target")
}