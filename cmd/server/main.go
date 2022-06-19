package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-rel/postgres"
	"github.com/go-rel/rel"
	_ "github.com/lib/pq"

	"backend/api"
)

func main() {
	var (
		httpPort   = os.Getenv("PORT")
		adapter    = initDbAdapter()
		repository = rel.New(adapter)
		r          = api.NewRouter(repository, httpPort)
	)
	defer adapter.Close()

	log.Printf("Server listening at %s", httpPort)
	if err := http.ListenAndServe(fmt.Sprintf("localhost:%s", httpPort), r); err != http.ErrServerClosed && err != nil {
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
