// Package config provides the configuration for the application. It contains the configuration schema and the default values for the configuration.
package config

import "time"

// Config represents the configuration for the application.
type Config struct {
	ConfigPath string    `mapstructure:"config_path"`
	LogLevel   string    `mapstructure:"log_level"`
	Stacktrace bool      `mapstructure:"stacktrace"`
	Get        Get       `mapstructure:"get"`
	Connector  Connector `mapstructure:"connector"`
}

// Get represents the configuration for the get command.
type Get struct {
	Timeout time.Duration `mapstructure:"timeout"`
}

// Connector represents the configuration for the connector.
type Connector struct {
	Binance Binance `mapstructure:"binance"`
}

// Binance represents the configuration for the Binance connector.
type Binance struct {
	BaseURL string `mapstructure:"base_url"`
}
