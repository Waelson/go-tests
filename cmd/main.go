package main

import (
	"database/sql"
	"github.com/Waelson/go-tests/internal/controller"
	"github.com/Waelson/go-tests/internal/repository"
	"github.com/Waelson/go-tests/internal/service"
	"github.com/Waelson/go-tests/internal/util"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

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
	controller := controller.NewUserController(userService)

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/users", func(r chi.Router) {
		r.Get("/", controller.List)  // GET /users
		r.Post("/", controller.Save) // POST /users
	})

	log.Println("Servidor executando na porta 8080")
	http.ListenAndServe(":8080", r)
}
