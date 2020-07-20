package cryptodini_server_test

import (
	"cryptodini/cryptodini_server"
	mock "cryptodini/cryptodini_server/mock"
	"cryptodini/models"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCryptodiniService_Adjust(t *testing.T) {

	mockUserID := uuid.New()
	mockPort := models.Portfolio{
		UserID: mockUserID,
		Assets: []models.Asset{{
			Symbol:          "BTC",
			Amount:          decimal.NewFromFloat(0.3011),
			TransactionDate: time.Now(),
			Status:          "BUY",
		}},
	}

	mockActiveUser := &models.User{
		UserID:   uuid.UUID{},
		Username: "John",
		Status:   "ACTIVE",
	}

	mockCurrentPort := &models.Portfolio{
		UserID: uuid.UUID{},
		Assets: nil,
		Status: nil,
	}

	mockOrders := []models.Order{
		{
			Symbol: "BTC",
			Type:   "BUY",
			Amount: decimal.NewFromFloat(0.2201),
		}, {
			Symbol: "USDT",
			Type:   "SELL",
			Amount: decimal.NewFromFloat(100.20),
		},
	}

	t.Run("successfully adjust", func(t *testing.T) {
		repoMock := mock.NewMockRepoService(gomock.NewController(t))
		repoMock.EXPECT().GetUser(mockUserID).Return(mockActiveUser, nil)
		repoMock.EXPECT().GetPort(mockUserID).Return(mockCurrentPort, nil)
		repoMock.EXPECT().ExecuteOrders(mockOrders).Return(nil)
		repoMock.EXPECT().SavePort(mockPort).Return(nil)

		service := cryptodini_server.Init(repoMock)
		err := service.Adjust(mockUserID, &mockPort)
		assert.NoError(t, err)
	})

	t.Run("fail, cannot find userID", func(t *testing.T) {
		expectedError := errors.New("not found user")
		repoMock := mock.NewMockRepoService(gomock.NewController(t))
		repoMock.EXPECT().GetUser(mockUserID).Return(nil, expectedError)

		service := cryptodini_server.Init(repoMock)
		err := service.Adjust(mockUserID, &mockPort)
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})

	t.Run("fail, specific user is not active", func(t *testing.T) {
		expectedError := errors.New("invalid user status")
		repoMock := mock.NewMockRepoService(gomock.NewController(t))
		repoMock.EXPECT().GetUser(mockUserID).Return(&models.User{
			UserID:   uuid.UUID{},
			Username: "John",
			Status:   "INACTIVE",
		}, nil)

		service := cryptodini_server.Init(repoMock)
		err := service.Adjust(mockUserID, &mockPort)
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})

	t.Run("fail, port of specific user not valid", func(t *testing.T) {
		expectedError := errors.New("port not valid for specific user")
		repoMock := mock.NewMockRepoService(gomock.NewController(t))
		repoMock.EXPECT().GetUser(mockUserID).Return(mockActiveUser, nil)
		repoMock.EXPECT().GetPort(mockUserID).Return(nil, expectedError)

		service := cryptodini_server.Init(repoMock)
		err := service.Adjust(mockUserID, &mockPort)
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})

	t.Run("fail, port adjust amount is invalid", func(t *testing.T) {})

	t.Run("fail, cannot send command to custodian", func(t *testing.T) {
		expectedError := errors.New("port not valid for specific user")
		repoMock := mock.NewMockRepoService(gomock.NewController(t))
		repoMock.EXPECT().GetUser(mockUserID).Return(mockActiveUser, nil)
		repoMock.EXPECT().GetPort(mockUserID).Return(mockCurrentPort, nil)
		repoMock.EXPECT().ExecuteOrders(mockOrders).Return(expectedError)

		service := cryptodini_server.Init(repoMock)
		err := service.Adjust(mockUserID, &mockPort)
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})

}
