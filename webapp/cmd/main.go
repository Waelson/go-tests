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
	"github.com/prometheus/client_golang/prometheus/promhttp"
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

	if os.Getenv("ENV") == "production" {
		logFile, err := os.OpenFile("/logs/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("Failed to open log file: %v", err)
		}
		log.SetOutput(logFile)
	}

	log.SetFormatter(&log.TextFormatter{})
	log.SetLevel(log.InfoLevel)

	log.Infof("[main] Starting application")
	database, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}

	err = util.CreateTables(database)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	log.Infof("[main] Creating components")
	metricsRecord := util.NewMetricsRecord()
	userRepository := repository.NewUserRepository(database)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService, metricsRecord)
	metricMiddleware := util.NewMetricMiddleware(metricsRecord)

	r := chi.NewRouter()

	log.Infof("[main] Registering routes")
	r.Use(metricMiddleware.Handler)
	r.Use(middleware.Recoverer)
	r.Handle("/metrics", promhttp.Handler())
	r.Get("/swagger/*", httpSwagger.WrapHandler)
	r.Route("/users", func(r chi.Router) {
		r.Get("/", userController.List)  // GET /users
		r.Post("/", userController.Save) // POST /users
	})

	log.Infof("[main] Web server started at :8080")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
