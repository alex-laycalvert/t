package cmd

import (
	"alex-laycalvert/t/internal/utils"
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop <project>",
	Short: "Stops the current timer for a project.",
	Args:  cobra.ExactArgs(1),
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
		db, err := getDB(cmd)
		cobra.CheckErr(err)

		activeProjects, err := db.ListOngoingProjects(context.Background())
		cobra.CheckErr(err)

		var matches []string
		for _, target := range activeProjects {
			if strings.HasPrefix(target, toComplete) {
				matches = append(matches, target)
			}
		}

		return matches, cobra.ShellCompDirectiveNoFileComp
	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			cobra.CheckErr(fmt.Errorf("stop needs a project name"))
		}

		db, err := getDB(cmd)
		cobra.CheckErr(err)

		projectName := args[0]
		projectTimer, err := db.StopTimer(context.Background(), projectName)
		cobra.CheckErr(err)

		fmt.Printf("Stopped timer for %s\n", projectTimer.ProjectName)

		started := time.Unix(projectTimer.StartSeconds, 0)
		stopped := time.Unix(projectTimer.StopSeconds, 0)
		elapsed := projectTimer.StopSeconds - projectTimer.StartSeconds

		fmt.Printf("  Started: %s\n", started.Format(utils.DateTimeLayout))
		fmt.Printf("  Stopped: %s\n", stopped.Format(utils.DateTimeLayout))
		fmt.Printf("  Elapsed: %s\n", utils.FormatElapsedTime(elapsed))
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
