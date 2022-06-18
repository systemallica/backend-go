package handlers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"backend/api/handlers"
	"backend/rides"
	"backend/utils"

	"github.com/go-rel/rel"
	"github.com/go-rel/rel/where"
	"github.com/go-rel/reltest"
	"github.com/stretchr/testify/assert"
)

func TestStartRide(t *testing.T) {
	var (
		request    = handlers.RideRequest{UserID: "1", VehicleID: "1"}
		body, _    = json.Marshal(request)
		req, _     = http.NewRequest("POST", "/", bytes.NewBuffer(body))
		rr         = httptest.NewRecorder()
		repository = reltest.New()
		service    = rides.New(repository)
		handler    = handlers.NewRidesHandler(repository, service)
	)
	req.Header.Add("Content-Type", "application/json")

	repository.ExpectCount("rides", rel.Or(
		rel.And(
			rel.Eq("user_id", request.UserID), rel.Eq("finished", false),
		),
	),
		rel.And(
			rel.Eq("vehicle_id", request.VehicleID), rel.Eq("finished", false),
		),
	).Result(0)
	repository.ExpectInsert().ForType("*rides.Ride")

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)

	var ride rides.Ride
	if err := json.NewDecoder(rr.Body).Decode(&ride); err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}
	assert.Equal(t, ride.Price, utils.GetEnvAsInt("RIDE_INITIAL_PRICE", 18))
	assert.NotNil(t, ride.CreatedAt)
	assert.NotNil(t, ride.UpdatedAt)
	assert.False(t, ride.Finished)
}

func TestStartRideBadRequestVehicle(t *testing.T) {
	var (
		request    = handlers.RideRequest{UserID: "1"}
		body, _    = json.Marshal(request)
		req, _     = http.NewRequest("POST", "/", bytes.NewBuffer(body))
		rr         = httptest.NewRecorder()
		repository = reltest.New()
		service    = rides.New(repository)
		handler    = handlers.NewRidesHandler(repository, service)
	)
	req.Header.Add("Content-Type", "application/json")

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)

	var resp handlers.ErrResponse
	json.Unmarshal([]byte(rr.Body.String()), &resp)
	assert.Equal(t, resp.StatusText, "Invalid request.")
	assert.Equal(t, resp.ErrorText, "missing required VehicleID field.")
}

func TestStartRideBadRequestUser(t *testing.T) {
	var (
		request    = handlers.RideRequest{VehicleID: "1"}
		body, _    = json.Marshal(request)
		req, _     = http.NewRequest("POST", "/", bytes.NewBuffer(body))
		rr         = httptest.NewRecorder()
		repository = reltest.New()
		service    = rides.New(repository)
		handler    = handlers.NewRidesHandler(repository, service)
	)
	req.Header.Add("Content-Type", "application/json")

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)

	var resp handlers.ErrResponse
	json.Unmarshal([]byte(rr.Body.String()), &resp)
	assert.Equal(t, resp.StatusText, "Invalid request.")
	assert.Equal(t, resp.ErrorText, "missing required UserID field.")
}

func TestStartRideAlreadyStarted(t *testing.T) {
	var (
		request    = handlers.RideRequest{UserID: "1", VehicleID: "1"}
		body, _    = json.Marshal(request)
		req, _     = http.NewRequest("POST", "/", bytes.NewBuffer(body))
		rr         = httptest.NewRecorder()
		repository = reltest.New()
		service    = rides.New(repository)
		handler    = handlers.NewRidesHandler(repository, service)
	)
	req.Header.Add("Content-Type", "application/json")

	repository.ExpectCount("rides", rel.Or(
		rel.And(
			rel.Eq("user_id", request.UserID), rel.Eq("finished", false),
		),
	),
		rel.And(
			rel.Eq("vehicle_id", request.VehicleID), rel.Eq("finished", false),
		),
	).Result(1)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)

	var resp handlers.ErrResponse
	json.Unmarshal([]byte(rr.Body.String()), &resp)
	assert.Equal(t, resp.StatusText, "Error while starting ride.")
	assert.Equal(t, resp.ErrorText, "A ride is already started for this vehicle or user")
}

func TestFinishRide(t *testing.T) {
	var (
		rideID      = uint(1)
		now         = time.Now()
		startedRide = rides.Ride{ID: rideID, VehicleID: "1", UserID: "1", Price: 18, Finished: false, CreatedAt: now, UpdatedAt: now}
		path        = fmt.Sprintf("/%d/finish", rideID)
		req, _      = http.NewRequest("POST", path, nil)
		rr          = httptest.NewRecorder()
		repository  = reltest.New()
		service     = rides.New(repository)
		handler     = handlers.NewRidesHandler(repository, service)
	)
	req.Header.Add("Content-Type", "application/json")

	repository.ExpectFind(where.Eq("id", "1")).Result(startedRide)
	repository.ExpectCount("rides", rel.And(rel.Eq("id", rideID), rel.Eq("finished", true))).Result(0)
	repository.ExpectUpdate().ForType("*rides.Ride")

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var ride rides.Ride
	if err := json.NewDecoder(rr.Body).Decode(&ride); err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}
	assert.Greater(t, ride.Price, utils.GetEnvAsInt("RIDE_INITIAL_PRICE", 18))
	assert.Less(t, ride.Price, 218)
	assert.True(t, ride.Finished)
}

func TestFinishRideAlreadyFinished(t *testing.T) {
	var (
		rideID     = uint(1)
		path       = fmt.Sprintf("/%d/finish", rideID)
		req, _     = http.NewRequest("POST", path, nil)
		rr         = httptest.NewRecorder()
		repository = reltest.New()
		service    = rides.New(repository)
		handler    = handlers.NewRidesHandler(repository, service)
	)
	req.Header.Add("Content-Type", "application/json")

	startedRide := rides.Ride{ID: rideID, VehicleID: "1", UserID: "1", Price: 18, Finished: false}
	repository.ExpectFind(where.Eq("id", "1")).Result(startedRide)
	repository.ExpectCount("rides", rel.And(rel.Eq("id", rideID), rel.Eq("finished", true))).Result(1)
	repository.ExpectUpdate().ForType("*rides.Ride")

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)

	var resp handlers.ErrResponse
	json.Unmarshal([]byte(rr.Body.String()), &resp)
	assert.Equal(t, resp.StatusText, "Error while finishing ride.")
	assert.Equal(t, resp.ErrorText, "This ride is already finished")
}
