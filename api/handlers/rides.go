package handlers

import (
	"errors"
	"net/http"
	"time"

	"backend/rides"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/go-rel/rel"
	"github.com/go-rel/rel/where"
)

type Rides struct {
	*chi.Mux
	repository rel.Repository
	rides      rides.Service
}

type RideRequest struct {
	UserID    string `json:"user_id"`
	VehicleID string `json:"vehicle_id"`
}

var (
	ErrRideUserIDBlank    = errors.New("missing required UserID field.")
	ErrRideVehicleIDBlank = errors.New("missing required VehicleID field.")
)

func (ride *RideRequest) Bind(r *http.Request) error {
	if ride.UserID == "" {
		return ErrRideUserIDBlank
	}
	if ride.VehicleID == "" {
		return ErrRideVehicleIDBlank
	}
	return nil
}

// RideResponse is the response payload for the Ride data model.
type RideResponse struct {
	*rides.Ride
}

func (rd *RideResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// Rides godoc
// @Summary starts a ride.
// @Description create ride
// @Tags rides
// @Accept json
// @Produce json
// @Param params body RideRequest true "Ride request parameters"
// @Success 201 {object} rides.Ride
// @Router /rides [post]
func (r Rides) RideStartHandler(w http.ResponseWriter, req *http.Request) {
	data := &RideRequest{}
	if err := render.Bind(req, data); err != nil {
		if err := render.Render(w, req, ErrInvalidRequest(err)); err != nil {
			return
		}
		return
	}

	ride := rides.Ride{
		UserID:    data.UserID,
		VehicleID: data.VehicleID,
	}

	if err, savedRide := r.rides.StartRide(req.Context(), &ride); err != nil {
		if err := render.Render(w, req, ErrStartDB(err)); err != nil {
			return
		}
		return
	} else {
		resp := &RideResponse{Ride: savedRide}

		render.Status(req, http.StatusCreated)
		if err := render.Render(w, req, resp); err != nil {
			return
		}
	}
}

// Rides godoc
// @Summary finishes the ride that matches the given ID.
// @Description finish ride
// @Tags rides
// @Accept json
// @Produce json
// @Param id path string true "Ride ID"
// @Success 200 {object} rides.Ride
// @Router /rides/:id/finish [post]
func (r Rides) RideFinishHandler(w http.ResponseWriter, req *http.Request) {
	rideID := chi.URLParam(req, "id")

	var ride rides.Ride
	if err := r.repository.Find(req.Context(), &ride, where.Eq("id", rideID)); err != nil {
		if err := render.Render(w, req, ErrFindDB(err)); err != nil {
			return
		}
		return
	}

	if err, savedRide := r.rides.FinishRide(req.Context(), &ride, time.Now()); err != nil {
		if err := render.Render(w, req, ErrFinishDB(err)); err != nil {
			return
		}
		return
	} else {
		resp := &RideResponse{Ride: savedRide}

		render.Status(req, http.StatusOK)
		if err := render.Render(w, req, resp); err != nil {
			return
		}
	}
}

func NewRidesHandler(repository rel.Repository, rides rides.Service) Rides {
	r := Rides{
		Mux:        chi.NewRouter(),
		repository: repository,
		rides:      rides,
	}

	r.Post("/", r.RideStartHandler)
	r.Post("/{id}/finish", r.RideFinishHandler)

	return r
}
