package cmd

import (
	"alex-laycalvert/t/internal/utils"
	"context"
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:               "show <project>",
	Short:             "Shows all recorded timers and progress for a project.",
	Args:              cobra.ExactArgs(1),
	ValidArgsFunction: projectsArgsFunction(false),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("show needs a project name")
		}

		db, err := getDB(cmd)
		if err != nil {
			return err
		}

		projectName := args[0]
		projectTimers, err := db.Queries.GetProject(context.Background(), projectName)
		if err != nil {
			return err
		}

		fmt.Printf("Time entries for %s\n", projectName)

		var totalTime int64 = 0

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
		fmt.Fprintf(w, "i\tElapsed\tStarted\tStopped\n")
		for i, timer := range projectTimers {
			started := time.Unix(timer.StartSeconds, 0).Format(utils.DateTimeLayout)
			stopped := "Ongoing"
			if timer.StopSeconds != -1 {
				stopped = time.Unix(timer.StopSeconds, 0).Format(utils.DateTimeLayout)
			}

			var elapsed int64 = 0
			if timer.StopSeconds == -1 {
				elapsed = time.Now().Unix() - timer.StartSeconds
			} else {
				elapsed = timer.StopSeconds - timer.StartSeconds
			}
			totalTime += elapsed

			fmt.Fprintf(w, "%d\t%s\t%s\t%s\n",
				i+1,
				utils.FormatElapsedTime(elapsed),
				started,
				stopped,
			)
		}
		w.Flush()

		fmt.Printf("Total time: %s\n", utils.FormatElapsedTime(totalTime))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
