package cryptodini_server_test

import (
	"cryptodini/cryptodini_server"
	mock "cryptodini/cryptodini_server/mock"
	"cryptodini/models"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCryptodiniService_GetPort(t *testing.T) {

	mockUserID := uuid.New()
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

	t.Run("successfully get port", func(t *testing.T) {
		repoMock := mock.NewMockRepoService(gomock.NewController(t))
		repoMock.EXPECT().GetUser(mockUserID).Return(mockActiveUser, nil)
		repoMock.EXPECT().GetPort(mockUserID).Return(mockCurrentPort, nil)

		service := cryptodini_server.Init(repoMock)
		port, err := service.GetPort(mockUserID)
		assert.NoError(t, err)
		assert.NotNil(t, port)
	})

}
