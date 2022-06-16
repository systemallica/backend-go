package app

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func RideStartHandler(w http.ResponseWriter, r *http.Request) {
}

func RideFinishHandler(w http.ResponseWriter, r *http.Request) {
	log.Print(chi.URLParam(r, "id"))
}
