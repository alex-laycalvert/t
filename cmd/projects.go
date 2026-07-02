package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

var projectsCmd = &cobra.Command{
	Use:   "projects",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		db, err := getDB(cmd)
		cobra.CheckErr(err)

		status, err := cmd.Flags().GetString("status")
		cobra.CheckErr(err)

		var projects []string
		switch status {
		case "ongoing":
			projects, err = db.ListOngoingProjects(context.Background())
			cobra.CheckErr(err)
		case "stopped":
			projects, err = db.ListStoppedProjects(context.Background())
			cobra.CheckErr(err)
		case "all":
			projects, err = db.ListProjects(context.Background())
			cobra.CheckErr(err)
		default:
			cobra.CheckErr(fmt.Errorf("unknown status (must be 'all', 'ongoing', or 'stopped')"))
		}

		for _, project := range projects {
			fmt.Println(project)
		}
	},
}

func init() {
	projectsCmd.Flags().StringP("status", "s", "all", "status of projects to list ('all', 'ongoing', 'stopped')")
	projectsCmd.RegisterFlagCompletionFunc("status", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"all", "ongoing", "stopped"}, cobra.ShellCompDirectiveDefault
	})

	rootCmd.AddCommand(projectsCmd)
}
