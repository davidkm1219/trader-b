package binance_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	binance_connector "github.com/binance/binance-connector-go"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/twk/trader-b/internal/connector/binance"
	mock_binance "github.com/twk/trader-b/internal/connector/binance/mocks"
)

func TestService_GetAccount_ErrorFromDo(t *testing.T) {
	type fields struct {
		mockOperation func(client *mock_binance.MockClient, accountClient *mock_binance.MockAccountClient)
	}

	type want struct {
		res *binance_connector.AccountResponse
		err error
	}

	tests := map[string]struct {
		fields fields
		want   want
	}{
		"GetAccountDoError": {
			fields: fields{
				mockOperation: func(client *mock_binance.MockClient, accountClient *mock_binance.MockAccountClient) {
					accountClient.EXPECT().Do(gomock.Any()).Return(nil, errors.New("do error"))
					client.EXPECT().NewGetAccountService().Return(accountClient)
				},
			},
			want: want{
				err: fmt.Errorf("error getting account: %w", errors.New("do error")),
			},
		},
		"Success": {
			fields: fields{
				mockOperation: func(client *mock_binance.MockClient, accountClient *mock_binance.MockAccountClient) {
					accountClient.EXPECT().Do(gomock.Any()).Return(&binance_connector.AccountResponse{}, nil)
					client.EXPECT().NewGetAccountService().Return(accountClient)
				},
			},
			want: want{
				res: &binance_connector.AccountResponse{},
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockClient := mock_binance.NewMockClient(ctrl)
			mockAccountClient := mock_binance.NewMockAccountClient(ctrl)
			tt.fields.mockOperation(mockClient, mockAccountClient)
			service := binance.NewService(mockClient)

			res, err := service.GetAccount(context.Background())
			if tt.want.err != nil {
				assert.EqualError(t, err, tt.want.err.Error())
				return
			}

			assert.Equal(t, tt.want.res, res)
		})
	}
}

func TestService_GetExchangeInfo_ErrorFromDo(t *testing.T) {
	type fields struct {
		mockOperation func(client *mock_binance.MockClient, exchangeInfoClient *mock_binance.MockExchangeInfoClient)
	}

	type want struct {
		res *binance_connector.ExchangeInfoResponse
		err error
	}

	tests := map[string]struct {
		fields fields
		want   want
	}{
		"GetExchangeInfoDoError": {
			fields: fields{
				mockOperation: func(client *mock_binance.MockClient, exchangeInfoClient *mock_binance.MockExchangeInfoClient) {
					exchangeInfoClient.EXPECT().Do(gomock.Any()).Return(nil, errors.New("do error"))
					client.EXPECT().NewExchangeInfoService().Return(exchangeInfoClient)
				},
			},
			want: want{
				err: fmt.Errorf("error getting exchange info: %w", errors.New("do error")),
			},
		},
		"Success": {
			fields: fields{
				mockOperation: func(client *mock_binance.MockClient, exchangeInfoClient *mock_binance.MockExchangeInfoClient) {
					exchangeInfoClient.EXPECT().Do(gomock.Any()).Return(&binance_connector.ExchangeInfoResponse{}, nil)
					client.EXPECT().NewExchangeInfoService().Return(exchangeInfoClient)
				},
			},
			want: want{
				res: &binance_connector.ExchangeInfoResponse{},
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockClient := mock_binance.NewMockClient(ctrl)
			mockExchangeInfoClient := mock_binance.NewMockExchangeInfoClient(ctrl)
			tt.fields.mockOperation(mockClient, mockExchangeInfoClient)
			service := binance.NewService(mockClient)

			res, err := service.GetExchangeInfo(context.Background())
			if tt.want.err != nil {
				assert.EqualError(t, err, tt.want.err.Error())
				return
			}

			assert.Equal(t, tt.want.res, res)
		})
	}
}
