package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Analyze replay results",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Analyzing replay results...")
	},
}