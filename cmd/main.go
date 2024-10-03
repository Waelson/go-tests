package main

import (
	"database/sql"
	"github.com/Waelson/go-tests/internal/controller"
	_ "github.com/Waelson/go-tests/internal/docs"
	"github.com/Waelson/go-tests/internal/repository"
	"github.com/Waelson/go-tests/internal/service"
	"github.com/Waelson/go-tests/internal/util"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/mattn/go-sqlite3"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
)

// @title Simples App Golang
// @version 1.0.0
// @contact.name Waelson Nunes
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {

	database, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}

	err = util.CreateTables(database)
	if err != nil {
		log.Fatal(err)
	}

	userRepository := repository.NewUserRepository(database)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/swagger/*", httpSwagger.WrapHandler)

	r.Route("/users", func(r chi.Router) {
		r.Get("/", userController.List)  // GET /users
		r.Post("/", userController.Save) // POST /users
	})

	log.Println("Servidor executando na porta 8080")
	http.ListenAndServe(":8080", r)
}
