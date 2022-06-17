package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/go-rel/postgres"
	"github.com/go-rel/rel"
	_ "github.com/lib/pq"

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
	var (
		httpPort   = os.Getenv("PORT")
		adapter    = initDbAdapter()
		repository = rel.New(adapter)
	)
	defer adapter.Close()
		log.Fatalf("Error starting http server <%s>", err)
	}
}

func initDbAdapter() rel.Adapter {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("POSTGRESQL_USERNAME"),
		os.Getenv("POSTGRESQL_PASSWORD"),
		os.Getenv("POSTGRESQL_HOST"),
		os.Getenv("POSTGRESQL_PORT"),
		os.Getenv("POSTGRESQL_DATABASE"))

	adapter, err := postgres.Open(dsn)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return adapter
}
