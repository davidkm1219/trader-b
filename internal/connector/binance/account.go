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

// AccountService is a service for interacting with the Binance account.
type AccountService struct {
	client AccountClient
}

// NewAccountService creates a new account service.
func NewAccountService(client AccountClient) *AccountService {
	return &AccountService{client: client}
}

// GetAccount gets the account information from Binance.
func (s *AccountService) GetAccount(ctx context.Context) (*binance_connector.AccountResponse, error) {
	res, err := s.client.Do(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting account: %w", err)
	}

	return res, nil
}