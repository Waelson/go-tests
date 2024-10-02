package controller

import (
	"encoding/json"
	"github.com/Waelson/go-tests/internal/model"
	"github.com/Waelson/go-tests/internal/service"
	"net/http"
)

type UserController interface {
	List(w http.ResponseWriter, r *http.Request)
	Save(w http.ResponseWriter, r *http.Request)
}

type userController struct {
	userService service.UserService
}

func (u *userController) List(w http.ResponseWriter, r *http.Request) {
	users, err := u.userService.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(response)
}

func (u *userController) Save(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = u.userService.Save(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func NewUserController(userService service.UserService) UserController {
	return &userController{userService}
}
