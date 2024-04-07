// Package commands provides the command line interface for the application. It contains the root command and all the subcommands.
package commands

import (
	"fmt"

	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/twk/trader-b/cmd/trader-b/commands/binance"
	"github.com/twk/trader-b/internal/config"
)

const appName = "trader-b"

// NewRootCommand creates a new cobra command for the root command
func NewRootCommand(logger *zap.Logger) (*cobra.Command, error) {
	v := config.NewViper()

	b := []config.BindDetail{
		{Flag: config.FlagDetail{Name: "config", Description: fmt.Sprintf("Specifies the path to the configuration file for %s.", appName), DefaultValue: "./config.yaml"}, MapKey: "config_path"},
		{Flag: config.FlagDetail{Name: "log-level", Description: "Determines the logging verbosity level for the application. Available options are 'debug', 'info', 'warn', and 'error'.", DefaultValue: ""}, EnvName: "LOG_LEVEL", MapKey: "log_level"},
		{Flag: config.FlagDetail{Name: "stacktrace", Description: "Enables or disables the inclusion of stack traces in the log output.", DefaultValue: false}, EnvName: "STACKTRACE", MapKey: "stacktrace"},
	}

	rootCmd := &cobra.Command{
		Use:   appName,
		Short: "CLI for the trader-b application",
		Long: `CLI for the trader-b application.
This CLI is used to interact with the trader-b application.`,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return cmd.Help()
		},
		SilenceUsage: true,
	}

	if err := v.SetFlagAndBind(rootCmd, b); err != nil {
		return nil, fmt.Errorf("error initializing flags: %w", err)
	}

	rootCmd.AddCommand(NewGetCmd(v, logger))
	rootCmd.AddCommand(binance.NewBinanceCommand(v, logger))

	return rootCmd, nil
}
