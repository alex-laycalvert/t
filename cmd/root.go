package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "t",
	Short: "Yet another lightweight, terminal time-tracking tool.",
	Long: `This does what every other time-tracking tool does, and probably less. I built it because
I wanted to.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
