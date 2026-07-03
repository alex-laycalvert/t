package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start <project>",
	Short: "Starts a timer for a project.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("start needs a project name")
		}

		db, err := getDB(cmd)
		if err != nil {
			return err
		}

		projectName := args[0]
		if err := db.Queries.StartTimer(context.Background(), projectName); err != nil {
			return err
		}
		fmt.Printf("%s started.\n", projectName)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
