package binance_test

import (
	"context"
	"errors"
	"testing"

	binance_connector "github.com/binance/binance-connector-go"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/twk/trader-b/internal/connector/binance"
	mock_binance "github.com/twk/trader-b/internal/connector/binance/mocks"
)

func TestAccountService_GetAccount(t *testing.T) {
	type fields struct {
		mockOperation func(client *mock_binance.MockAccountClient)
	}

	type want struct {
		res *binance_connector.AccountResponse
		err error
	}

	tests := map[string]struct {
		fields fields
		want   want
	}{
		"GetAccountSuccess": {
			fields: fields{
				mockOperation: func(client *mock_binance.MockAccountClient) {
					client.EXPECT().Do(context.Background()).Return(&binance_connector.AccountResponse{}, nil)
				},
			},
			want: want{
				res: &binance_connector.AccountResponse{},
			},
		},
		"GetAccountError": {
			fields: fields{
				mockOperation: func(client *mock_binance.MockAccountClient) {
					client.EXPECT().Do(context.Background()).Return(nil, errors.New("error"))
				},
			},
			want: want{
				err: errors.New("error getting account: error"),
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockClient := mock_binance.NewMockAccountClient(ctrl)
			tt.fields.mockOperation(mockClient)
			service := binance.NewAccountService(mockClient)

			res, err := service.GetAccount(context.Background())
			if tt.want.err != nil {
				assert.EqualError(t, err, tt.want.err.Error())
				return
			}

			assert.Equal(t, tt.want.res, res)
		})
	}
}
