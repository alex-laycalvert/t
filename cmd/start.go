package cmd

import (
	"alex-laycalvert/t/internal/config"
	"alex-laycalvert/t/internal/db"
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start <project>",
	Short: "Starts a timer for a project.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			cobra.CheckErr(fmt.Errorf("start needs a project name"))
		}

		config, err := config.New()
		cobra.CheckErr(err)
		db, err := db.Provide(config.DatabasePath)
		cobra.CheckErr(err)

		projectName := args[0]
		err = db.StartTimer(context.Background(), projectName)
		cobra.CheckErr(err)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
