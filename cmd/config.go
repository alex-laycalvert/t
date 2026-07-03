package cmd

import (
	"alex-laycalvert/t/internal/config"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage configuration",
	Long: `View and modify configuration settings.
Config file defaults to ~/.config/t/config.
Override with the T_CONFIG_FILE environment variable.`,
	// TODO:
	// PersistentPreRun: func(cmd *cobra.Command, args []string) {
	// Override root PersistentPreRun — config commands don't need DB access.
	// },
}

var configGetCmd = &cobra.Command{
	Use:   "get <key>",
	Short: "Get the value of a config property",
	Args:  cobra.ExactArgs(1),
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		if len(args) != 0 {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		return config.ValidKeys(), cobra.ShellCompDirectiveNoFileComp
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		keyStr := args[0]
		if !config.IsConfigKey(keyStr) {
			return fmt.Errorf("invalid config key")
		}
		key := config.ConfigKey(keyStr)

		cfg, err := getCfg(cmd)
		if err != nil {
			return err
		}
		value := cfg.Get(key)
		fmt.Println(value)
		return nil
	},
}

var configSetCmd = &cobra.Command{
	Use:   "set <key> <value>",
	Short: "Set the value of a config property",
	Args:  cobra.ExactArgs(2),
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		if len(args) == 0 {
			return config.ValidKeys(), cobra.ShellCompDirectiveNoFileComp
		}
		return nil, cobra.ShellCompDirectiveDefault
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		keyStr := args[0]
		if !config.IsConfigKey(keyStr) {
			return fmt.Errorf("invalid config key")
		}
		key := config.ConfigKey(keyStr)
		value := args[1]

		cfg, err := getCfg(cmd)
		if err != nil {
			return err
		}
		if err := cfg.Set(key, value); err != nil {
			return err
		}
		fmt.Printf("%s = %s\n", args[0], args[1])
		return nil
	},
}

var configOpenCmd = &cobra.Command{
	Use:   "open",
	Short: "Open config file in your preferred editor",
	Long:  "Opens the config file in the editor specified by the EDITOR environment variable.",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		editor := os.Getenv("EDITOR")
		if editor == "" {
			return fmt.Errorf("EDITOR environment variable is not set")
		}

		cfg, err := getCfg(cmd)
		if err != nil {
			return err
		}

		// Ensure the config file exists before opening.
		if err := os.MkdirAll(filepath.Dir(cfg.FilePath), 0755); err != nil {
			return err
		}

		if _, err := os.Stat(cfg.FilePath); os.IsNotExist(err) {
			f, err := os.Create(cfg.FilePath)
			if err != nil {
				return err
			}
			f.Close()
		}

		c := exec.Command(editor, cfg.FilePath)
		c.Stdin = os.Stdin
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		return c.Run()
	},
}

func init() {
	configCmd.AddCommand(configGetCmd)
	configCmd.AddCommand(configSetCmd)
	configCmd.AddCommand(configOpenCmd)
	rootCmd.AddCommand(configCmd)
}
