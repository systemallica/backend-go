package rides

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRideValidation(t *testing.T) {
	var ride Ride

	t.Run("UserID is blank", func(t *testing.T) {
		assert.Equal(t, ErrRideUserIDBlank, ride.Validate())
	})

	t.Run("UserID is blank", func(t *testing.T) {
		ride.UserID = "1"
		assert.Equal(t, ErrRideVehicleIDBlank, ride.Validate())
	})

	t.Run("valid", func(t *testing.T) {
		ride.VehicleID = "1"
		assert.Nil(t, ride.Validate())
	})
}
