package service_test

import (
	"github.com/Waelson/go-tests/internal/mocks"
	"github.com/Waelson/go-tests/internal/model"
	"github.com/Waelson/go-tests/internal/service"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserService_FindAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mocks.NewMockUserRepository(ctrl)

	userService := service.NewUserService(mockUserRepo)

	mockUsers := []model.User{
		{ID: 1, Username: "John Doe"},
		{ID: 2, Username: "Jane Doe"},
	}

	mockUserRepo.EXPECT().FindAll().Return(mockUsers, nil)

	users, err := userService.FindAll()

	assert.NoError(t, err)
	assert.Equal(t, mockUsers, users)
}

func TestUserService_Save(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mocks.NewMockUserRepository(ctrl)

	userService := service.NewUserService(mockUserRepo)

	mockUser := model.User{ID: 1, Username: "John Doe"}

	mockUserRepo.EXPECT().Save(mockUser).Return(mockUser, nil)

	savedUser, err := userService.Save(mockUser)

	assert.NoError(t, err)
	assert.Equal(t, mockUser, savedUser)
}

// Erro ao Listar dos os usuarios
func TestUserService_FindAll_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mocks.NewMockUserRepository(ctrl)

	userService := service.NewUserService(mockUserRepo)

	mockUserRepo.EXPECT().FindAll().Return(nil, assert.AnError)

	users, err := userService.FindAll()

	assert.Error(t, err)
	assert.Nil(t, users)
}

func TestUserService_Save_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mocks.NewMockUserRepository(ctrl)

	userService := service.NewUserService(mockUserRepo)

	mockUser := model.User{ID: 1, Username: "John Doe"}

	mockUserRepo.EXPECT().Save(mockUser).Return(model.User{}, assert.AnError)

	savedUser, err := userService.Save(mockUser)

	assert.Error(t, err)
	assert.Equal(t, model.User{}, savedUser)
}
