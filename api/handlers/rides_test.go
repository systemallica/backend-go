package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"backend/api/handlers"
	"backend/rides"
	"backend/utils"

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
