package rides

import (
	"context"

	"backend/utils"

	"github.com/go-rel/rel"
)

type startRide struct {
	repository rel.Repository
}

func (c startRide) StartRide(ctx context.Context, ride *Ride) (error, *Ride) {
	if err := ride.Validate(); err != nil {
		return err, nil
	}

	count, _ := c.repository.Count(ctx, "rides",
		rel.Eq("finished", false).And(
			rel.Eq("user_id", ride.UserID).OrEq("vehicle_id", ride.VehicleID),
		),
	)
	if count != 0 {
		return ErrRideAlreadyStarted, nil
	}

	ride.Price = utils.GetEnvAsInt("RIDE_INITIAL_PRICE", 18)

	err := c.repository.Insert(ctx, ride)
	return err, ride
}
