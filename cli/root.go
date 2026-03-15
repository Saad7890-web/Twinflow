package cli

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "twinflow",
	Short: "TwinFlow is a traffic capture and replay tool",
	Long: `TwinFlow records real HTTP traffic and replays it
to detect breaking behavior changes between service versions.`,
}


func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(recordCmd)
	rootCmd.AddCommand(replayCmd)
	rootCmd.AddCommand(analyzeCmd)
}