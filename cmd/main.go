package main

import (
	"database/sql"
	"fmt"
	"github.com/Waelson/go-tests/internal/controller"
	_ "github.com/Waelson/go-tests/internal/docs"
	"github.com/Waelson/go-tests/internal/repository"
	"github.com/Waelson/go-tests/internal/service"
	"github.com/Waelson/go-tests/internal/util"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
	"os"
)

// @title Simples App Golang
// @version 1.0.0
// @contact.name Waelson Nunes
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {

	fmt.Println("[main] Starting application")
	logFile, err := os.OpenFile("/logs/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Failed to open log file: %v", err)
	}
	log.SetOutput(logFile)
	log.SetFormatter(&log.TextFormatter{})
	log.SetLevel(log.InfoLevel)

	log.Infof("[main] Starting application")

	fmt.Println("[main] Open database")
	database, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("[main] Creating tables")
	err = util.CreateTables(database)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	fmt.Println("[main] Creating components")
	log.Infof("[main] Creating components")
	userRepository := repository.NewUserRepository(database)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	r := chi.NewRouter()

	fmt.Println("[main] Registering routes")
	log.Infof("[main] Registering routes")
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/swagger/*", httpSwagger.WrapHandler)

	r.Route("/users", func(r chi.Router) {
		r.Get("/", userController.List)  // GET /users
		r.Post("/", userController.Save) // POST /users
	})

	log.Infof("[main] Web server started at :8080")
	http.ListenAndServe(":8080", r)
}
