package cmd

import (
	"alex-laycalvert/t/internal/config"
	"alex-laycalvert/t/internal/db"
	"context"
	"errors"
	"os"

	"github.com/spf13/cobra"
)

type ctxKey string

const dbKey ctxKey = "database"
const cfgKey ctxKey = "config"

var rootCmd = &cobra.Command{
	Use:   "t",
	Short: "Yet another lightweight, terminal time-tracking tool.",
	Long: `This does what every other time-tracking tool does, and probably less. I built it because
I wanted to.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		config, err := config.New()
		cobra.CheckErr(err)
		ctx := context.WithValue(cmd.Context(), cfgKey, config)

		db, err := db.Provide(config)
		cobra.CheckErr(err)
		ctx = context.WithValue(ctx, dbKey, db)

		cmd.SetContext(ctx)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}

func getDB(cmd *cobra.Command) (*db.Queries, error) {
	db, ok := cmd.Context().Value(dbKey).(*db.Queries)
	if !ok {
		return nil, errors.New("db missing")
	}

	return db, nil
}

func getCfg(cmd *cobra.Command) (*config.Config, error) {
	cfg, ok := cmd.Context().Value(cfgKey).(*config.Config)
	if !ok {
		return nil, errors.New("config missing")
	}

	return cfg, nil
}
