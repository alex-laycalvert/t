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
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			cobra.CheckErr(fmt.Errorf("start needs a project name"))
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
				return
			}
		}

		db, err := getDB(cmd)
		cobra.CheckErr(err)

		err = db.DeleteProject(context.Background(), projectName)
		cobra.CheckErr(err)
		fmt.Printf("%s deleted.\n", projectName)
	},
}

func init() {
	deleteCmd.Flags().BoolP("yes", "y", false, "Skip confirmation")
	rootCmd.AddCommand(deleteCmd)
}
