package binance

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

// AccountClient is a client for interacting with the Binance account.
type AccountClient interface {
	Do(ctx context.Context, opts ...binance_connector.RequestOption) (res *binance_connector.AccountResponse, err error)
}

// ExchangeInfoClient is a client for interacting with the Binance exchange info.
type ExchangeInfoClient interface {
	Do(ctx context.Context, opts ...binance_connector.RequestOption) (res *binance_connector.ExchangeInfoResponse, err error)
}

// Client is a client for interacting with Binance.
type Client interface {
	NewGetAccountService() AccountClient
	NewExchangeInfoService() ExchangeInfoClient
}

// Service is a service for interacting with Binance.
type Service struct {
	client Client
}

// NewService creates a new service.
func NewService(client Client) *Service {
	return &Service{client: client}
}

// GetAccount gets the account information from Binance.
func (s *Service) GetAccount(ctx context.Context) (*binance_connector.AccountResponse, error) {
	accountService := s.client.NewGetAccountService()

	res, err := accountService.Do(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting account: %w", err)
	}

	return res, nil
}

// GetExchangeInfo gets the exchange info from Binance.
func (s *Service) GetExchangeInfo(ctx context.Context) (*binance_connector.ExchangeInfoResponse, error) {
	exchangeInfoService := s.client.NewExchangeInfoService()

	res, err := exchangeInfoService.Do(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting exchange info: %w", err)
	}

	return res, nil
}
