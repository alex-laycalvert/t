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
	RunE: func(cmd *cobra.Command, args []string) error {
		db, err := getDB(cmd)
		if err != nil {
			return err
		}

		status, err := cmd.Flags().GetString("status")
		if err != nil {
			return err
		}

		var projects []string
		switch status {
		case "ongoing":
			projects, err = db.Queries.ListOngoingProjects(context.Background())
		case "stopped":
			projects, err = db.Queries.ListStoppedProjects(context.Background())
		case "all":
			projects, err = db.Queries.ListProjects(context.Background())
		default:
			return fmt.Errorf("unknown status (must be 'all', 'ongoing', or 'stopped')")
		}
		if err != nil {
			return err
		}

		for _, project := range projects {
			fmt.Println(project)
		}
		return nil
	},
}

func init() {
	projectsCmd.Flags().StringP("status", "s", "all", "status of projects to list ('all', 'ongoing', 'stopped')")
	projectsCmd.RegisterFlagCompletionFunc("status", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"all", "ongoing", "stopped"}, cobra.ShellCompDirectiveDefault
	})

	rootCmd.AddCommand(projectsCmd)
}
