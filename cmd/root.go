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
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		config, err := config.New()
		if err != nil {
			return err
		}
		ctx := context.WithValue(cmd.Context(), cfgKey, config)

		db, err := db.Provide(config)
		if err != nil {
			return err
		}
		ctx = context.WithValue(ctx, dbKey, db)

		cmd.SetContext(ctx)
		return nil
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

func getDB(cmd *cobra.Command) (*db.DB, error) {
	db, ok := cmd.Context().Value(dbKey).(*db.DB)
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
