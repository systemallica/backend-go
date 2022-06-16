package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"

	"backend/app"
	"backend/docs"
)

const httpPort = 8080

// @title Rides Swagger API
// @description This is a basic Rides API using Chi and go-rel.

// @contact.name systemallica
// @contact.url http://www.andres.reveronmolina.me
// @contact.email andres@reveronmolina.me

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http https
func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/rides", app.RideStartHandler)
	r.Post("/rides/{id}/finish", app.RideFinishHandler)

	docs.SwaggerInfo.Version = "1.0"
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", httpPort)),
	))

	if err := http.ListenAndServe(fmt.Sprintf("localhost:%d", httpPort), r); err != http.ErrServerClosed && err != nil {
		log.Fatalf("Error starting http server <%s>", err)
	}
}
