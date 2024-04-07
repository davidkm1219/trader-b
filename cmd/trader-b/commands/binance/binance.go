// Package binance provides the Binance command for the application. It contains the NewBinanceCommand function and the binanceRun function.
package binance

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/twk/trader-b/internal/config"
	"go.uber.org/zap"
)

// NewBinanceCommand creates a new Binance command.
func NewBinanceCommand(v *config.Viper, l *zap.Logger) *cobra.Command {
	b := []config.BindDetail{
		{Flag: config.FlagDetail{Name: "binance-api-key", Description: "The API key for the Binance API", DefaultValue: ""}, MapKey: "connector.binance.api_key", EnvName: "BINANCE_API_KEY"},
		{Flag: config.FlagDetail{Name: "binance-api-secret", Description: "The API secret for the Binance API", DefaultValue: ""}, MapKey: "connector.binance.secret_key", EnvName: "BINANCE_API_SECRET"},
		{Flag: config.FlagDetail{Name: "binance-base-url", Description: "The base URL for the Binance API", DefaultValue: ""}, MapKey: "connector.binance.base_url", EnvName: "BINANCE_BASE_URL"},
	}

	cmd := &cobra.Command{
		Use:   "binance",
		Short: "Interact with the Binance API",
		RunE: func(_ *cobra.Command, _ []string) error {
			return binanceRun(v, l)
		},
	}

	if err := v.SetFlagAndBind(cmd, b); err != nil {
		return nil
	}

	return cmd
}

func binanceRun(v *config.Viper, l *zap.Logger) error {
	cfg, err := v.BuildConfig()
	if err != nil {
		return fmt.Errorf("error building config: %w", err)
	}

	l.Info("running binance command", zap.Any("config", cfg))

	return nil
}
