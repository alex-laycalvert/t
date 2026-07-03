package cmd

import (
	"alex-laycalvert/t/internal/utils"
	"context"
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:               "stop <project>",
	Short:             "Stops the current timer for a project.",
	Args:              cobra.ExactArgs(1),
	ValidArgsFunction: projectsArgsFunction(true),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("stop needs a project name")
		}

		db, err := getDB(cmd)
		if err != nil {
			return err
		}

		projectName := args[0]
		projectTimer, err := db.Queries.StopTimer(context.Background(), projectName)
		if err != nil {
			return err
		}

		fmt.Printf("Stopped timer for %s\n", projectTimer.ProjectName)

		started := time.Unix(projectTimer.StartSeconds, 0)
		stopped := time.Unix(projectTimer.StopSeconds, 0)
		elapsed := projectTimer.StopSeconds - projectTimer.StartSeconds

		fmt.Printf("  Started: %s\n", started.Format(utils.DateTimeLayout))
		fmt.Printf("  Stopped: %s\n", stopped.Format(utils.DateTimeLayout))
		fmt.Printf("  Elapsed: %s\n", utils.FormatElapsedTime(elapsed))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
