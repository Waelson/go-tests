package controller

import (
	"encoding/json"
	"github.com/Waelson/go-tests/internal/model"
	"github.com/Waelson/go-tests/internal/service"
	"github.com/Waelson/go-tests/internal/util"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type UserController interface {
	List(w http.ResponseWriter, r *http.Request)
	Save(w http.ResponseWriter, r *http.Request)
}

type userController struct {
	userService   service.UserService
	metricsRecord util.MetricsRecord
}

// List godoc
// @Summary      List user
// @Description  get users
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {array}   model.User
// @Router       /users [get]
func (u *userController) List(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	defer func() {
		log.Infof("[user_controller] List users took %v", time.Since(start))
	}()

	users, err := u.userService.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(response)
}

// Save godoc
// @Summary      Save user
// @Description  save user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user body model.User true "User"
// @Success      201  {string} string
// @Router       /users [post]
func (u *userController) Save(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	defer func() {
		log.Infof("[user_controller] Save user took %v", time.Since(start))
	}()

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

func NewUserController(userService service.UserService, metricsRecord util.MetricsRecord) UserController {
	return &userController{
		userService:   userService,
		metricsRecord: metricsRecord,
	}
}
