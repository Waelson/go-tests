package controller_test

import (
	"bytes"
	"encoding/json"
	"github.com/Waelson/go-tests/internal/controller"
	"github.com/Waelson/go-tests/internal/mocks"
	"github.com/Waelson/go-tests/internal/model"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserController_List(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := mocks.NewMockUserService(ctrl)

	mockUsers := []model.User{
		{ID: 1, Username: "John Doe", Password: "123456", Email: "john@gmail.com"},
		{ID: 2, Username: "Jane Doe", Password: "123456", Email: "jane@gmail.com"},
	}

	mockUserService.EXPECT().FindAll().Return(mockUsers, nil)

	userController := controller.NewUserController(mockUserService)

	req, err := http.NewRequest("GET", "/users", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(userController.List)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	expected := `[{"email":"john@gmail.com", "id":1, "password":"123456", "username":"John Doe"},{"email":"jane@gmail.com", "id":2, "password":"123456", "username":"Jane Doe"}]`
	assert.JSONEq(t, expected, rr.Body.String())
}

func TestUserController_Save(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := mocks.NewMockUserService(ctrl)

	userToSave := model.User{ID: 1, Username: "John Doe", Password: "123456", Email: "john@gmail.com"}
	mockUserService.EXPECT().Save(userToSave).Return(userToSave, nil)

	userController := controller.NewUserController(mockUserService)

	userJSON, err := json.Marshal(userToSave)
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(userJSON))
	assert.NoError(t, err)

	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(userController.Save)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
}
