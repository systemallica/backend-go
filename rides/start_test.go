package rides

import (
	"context"
	"testing"

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
