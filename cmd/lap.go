package cmd

import (
	"alex-laycalvert/t/internal/utils"
	"context"
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var lapCmd = &cobra.Command{
	Use:   "lap <project>",
	Short: "Lap the current current timer by stopping and restarting the project",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("lap needs a project name")
		}

		projectName := args[0]

		db, err := getDB(cmd)
		if err != nil {
			return err
		}

		tx, err := db.DB.Begin()
		if err != nil {
			return err
		}

		ctx := context.Background()
		qtx := db.Queries.WithTx(tx)
		project, err := qtx.StopTimer(ctx, projectName)
		if err != nil {
			return err
		}

		started := time.Unix(project.StartSeconds, 0)
		stopped := time.Unix(project.StopSeconds, 0)
		elapsed := project.StopSeconds - project.StartSeconds

		fmt.Printf("Stopped timer for %s\n", project.ProjectName)
		fmt.Printf("  Started: %s\n", started.Format(utils.DateTimeLayout))
		fmt.Printf("  Stopped: %s\n", stopped.Format(utils.DateTimeLayout))
		fmt.Printf("  Elapsed: %s\n", utils.FormatElapsedTime(elapsed))

		if err := qtx.StartTimer(ctx, projectName); err != nil {
			return err
		}

		defer func() {
			fmt.Printf("%s lapped.\n", projectName)
		}()

		return tx.Commit()
	},
}

func init() {
	rootCmd.AddCommand(lapCmd)
}
