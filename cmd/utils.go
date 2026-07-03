package cmd

import (
	"context"
	"strings"

	"github.com/spf13/cobra"
)

func projectsArgsFunction(onlyOngoing bool) cobra.CompletionFunc {
	return func(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
		db, err := getDB(cmd)
		if err != nil {
			return nil, cobra.ShellCompDirectiveError
		}

		var projects []string
		if onlyOngoing {
			projects, err = db.Queries.ListOngoingProjects(context.Background())
		} else {
			projects, err = db.Queries.ListProjects(context.Background())
		}
		if err != nil {
			return nil, cobra.ShellCompDirectiveError
		}

		var matches []string
		for _, target := range projects {
			if strings.HasPrefix(target, toComplete) {
				matches = append(matches, target)
			}
		}

		return matches, cobra.ShellCompDirectiveNoFileComp
	}
}
