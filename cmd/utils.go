package cmd

import (
	"context"
	"strings"

	"github.com/spf13/cobra"
)

func projectsArgsFunction(onlyOngoing bool) cobra.CompletionFunc {
	return func(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
		db, err := getDB(cmd)
		cobra.CheckErr(err)

		var projects []string
		if onlyOngoing {
			projects, err = db.ListOngoingProjects(context.Background())
		} else {
			projects, err = db.ListProjects(context.Background())
		}
		cobra.CheckErr(err)

		var matches []string
		for _, target := range projects {
			if strings.HasPrefix(target, toComplete) {
				matches = append(matches, target)
			}
		}

		return matches, cobra.ShellCompDirectiveNoFileComp
	}
}
