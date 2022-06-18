package rides

import (
	"context"
	"testing"

	"github.com/go-rel/rel"
	"github.com/go-rel/reltest"
	"github.com/stretchr/testify/assert"
)

func TestStart(t *testing.T) {
	var (
		ctx        = context.TODO()
		repository = reltest.New()
		service    = New(repository)
		ride       = Ride{UserID: "1", VehicleID: "1"}
	)

	repository.ExpectCount("rides", rel.Or(
		rel.And(
			rel.Eq("user_id", ride.UserID), rel.Eq("finished", false),
		),
	),
		rel.And(
			rel.Eq("vehicle_id", ride.VehicleID), rel.Eq("finished", false),
		),
	).Result(0)

	repository.ExpectInsert().For(&ride)

	err, _ := service.StartRide(ctx, &ride)
	assert.Nil(t, err)
	assert.NotEmpty(t, ride.ID)
	assert.Equal(t, ride.Price, 18)
	assert.NotEmpty(t, ride.CreatedAt)
	assert.NotEmpty(t, ride.UpdatedAt)
	assert.False(t, ride.Finished)

	repository.AssertExpectations(t)
}

func TestStartRideValidationErrorUserID(t *testing.T) {
	var (
		ctx        = context.TODO()
		repository = reltest.New()
		service    = New(repository)
		ride       = Ride{UserID: "", VehicleID: "1"}
	)

	err, _ := service.StartRide(ctx, &ride)
	assert.Equal(t, ErrRideUserIDBlank, err)

	repository.AssertExpectations(t)
}

func TestStartRideValidationErrorVehicleID(t *testing.T) {
	var (
		ctx        = context.TODO()
		repository = reltest.New()
		service    = New(repository)
		ride       = Ride{UserID: "1", VehicleID: ""}
	)

	err, _ := service.StartRide(ctx, &ride)
	assert.Equal(t, ErrRideVehicleIDBlank, err)

	repository.AssertExpectations(t)
}

func TestStartAnotherRideAlreadyStarted(t *testing.T) {
	var (
		ctx        = context.TODO()
		repository = reltest.New()
		service    = New(repository)
		ride       = Ride{UserID: "2", VehicleID: "2"}
	)
	repository.ExpectCount("rides", rel.Or(
		rel.And(
			rel.Eq("user_id", ride.UserID), rel.Eq("finished", false),
		),
	),
		rel.And(
			rel.Eq("vehicle_id", ride.VehicleID), rel.Eq("finished", false),
		),
	).Result(1)

	err, _ := service.StartRide(ctx, &ride)
	assert.Equal(t, ErrRideAlreadyStarted, err)
	assert.Empty(t, ride.ID)

	repository.AssertExpectations(t)
}
