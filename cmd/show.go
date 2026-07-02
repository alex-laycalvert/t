package cmd

import (
	"alex-laycalvert/t/internal/config"
	"alex-laycalvert/t/internal/db"
	"alex-laycalvert/t/internal/utils"
	"context"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show <project>",
	Short: "Shows all recorded timers and progress for a project.",
	Args:  cobra.ExactArgs(1),
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
		config, err := config.New()
		cobra.CheckErr(err)
		db, err := db.Provide(config.DatabasePath)
		cobra.CheckErr(err)

		projects, err := db.ListProjects(context.Background())
		cobra.CheckErr(err)

		var matches []string
		for _, target := range projects {
			if strings.HasPrefix(target, toComplete) {
				matches = append(matches, target)
			}
		}

		return matches, cobra.ShellCompDirectiveNoFileComp
	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			cobra.CheckErr(fmt.Errorf("show needs a project name"))
		}

		config, err := config.New()
		cobra.CheckErr(err)
		db, err := db.Provide(config.DatabasePath)
		cobra.CheckErr(err)

		projectName := args[0]
		projectTimers, err := db.GetProject(context.Background(), projectName)
		cobra.CheckErr(err)

		header := "Time entries for "
		fmt.Printf("%s%s\n", header, projectName)

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

			fmt.Fprintf(w, "%d\t%s\t%s\t%s\n",
				i,
				utils.FormatElapsedTime(elapsed),
				started,
				stopped,
			)
		}
		w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
