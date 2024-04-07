// Package binance provides the Binance service for the application. It contains the service and the methods to interact with the Binance API.
package binance

import (
	binance_connector "github.com/binance/binance-connector-go"

	"github.com/twk/trader-b/internal/config"
)

// NewBinanceClient creates a new Binance client.
func NewBinanceClient(cfg *config.Config, apiKey, apiSecret string) *binance_connector.Client {
	if cfg.Connector.Binance.BaseURL == "" {
		return binance_connector.NewClient(apiKey, apiSecret)
	}

	return binance_connector.NewClient(apiKey, apiSecret, cfg.Connector.Binance.BaseURL)
}
