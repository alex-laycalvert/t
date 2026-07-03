package cmd

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:               "delete <project>",
	Short:             "Deletes a project and all timers.",
	Args:              cobra.ExactArgs(1),
	ValidArgsFunction: projectsArgsFunction(false),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("start needs a project name")
		}
		projectName := args[0]

		forceDelete, _ := cmd.Flags().GetBool("yes")
		if !forceDelete {
			fmt.Printf("Are you sure you want to delete %s? (y/N): ", projectName)
			reader := bufio.NewReader(os.Stdin)
			response, _ := reader.ReadString('\n')
			response = strings.ToLower(strings.TrimSpace(response))

			if response != "y" && response != "yes" {
				fmt.Println("Delete canceled.")
				return nil
			}
		}

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
		if err := qtx.DeleteProject(ctx, projectName); err != nil {
			return err
		}
		if err := qtx.DeleteTimers(ctx, projectName); err != nil {
			return err
		}
		defer func() {
			fmt.Printf("%s deleted.\n", projectName)
		}()

		return tx.Commit()
	},
}

func init() {
	deleteCmd.Flags().BoolP("yes", "y", false, "Skip confirmation")
	rootCmd.AddCommand(deleteCmd)
}
