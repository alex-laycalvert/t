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
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			cobra.CheckErr(fmt.Errorf("start needs a project name"))
		}

		db, err := getDB(cmd)
		cobra.CheckErr(err)

		projectName := args[0]
		err = db.StartTimer(context.Background(), projectName)
		cobra.CheckErr(err)
		fmt.Printf("%s started.\n", projectName)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
