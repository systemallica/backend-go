package api

import (
	"backend/docs"
	"backend/rides"
	"fmt"

	h "backend/api/handlers"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	metrics "github.com/slok/go-http-metrics/metrics/prometheus"
	"github.com/slok/go-http-metrics/middleware"
	std "github.com/slok/go-http-metrics/middleware/std"

	"github.com/go-chi/chi/v5"
	chiMid "github.com/go-chi/chi/v5/middleware"
	"github.com/go-rel/rel"
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
	var (
		r            = chi.NewRouter()
		rides        = rides.New(repository)
		ridesHandler = h.NewRidesHandler(repository, rides)
		mdlw         = middleware.New(middleware.Config{
			Recorder: metrics.NewRecorder(metrics.Config{}),
		})
	)

	r.Use(chiMid.Logger)
	r.Use(chiMid.RequestID)
	r.Use(chiMid.RealIP)
	r.Use(chiMid.Recoverer)
	r.Use(std.HandlerProvider("", mdlw))

	r.Mount("/rides", ridesHandler)
	r.Mount("/metrics", promhttp.Handler())

	docs.SwaggerInfo.Version = "1.0"
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("http://localhost:%s/swagger/doc.json", port)),
	))

	return r
}
