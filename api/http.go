package api

import (
	"backend/docs"
	"backend/rides"
	"fmt"

	h "backend/api/handlers"

	"github.com/go-rel/rel"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)


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
func NewRouter(repository rel.Repository, port string) *chi.Mux {
	var(
		r = chi.NewRouter()
		rides = rides.New(repository)
		ridesHandler = h.NewRidesHandler(repository, rides)
	)

	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)

	r.Mount("/rides", ridesHandler)

	docs.SwaggerInfo.Version = "1.0"
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("http://localhost:%s/swagger/doc.json", port)),
	))

	return r
}
