package rides

import (
	"context"
	"math"
	"time"

	"backend/utils"

	"github.com/go-rel/rel"
)

type finishRide struct {
	repository rel.Repository
}

func (c startRide) FinishRide(ctx context.Context, ride *Ride, now time.Time) (error, *Ride) {
	var (
		duration    = now.Sub(ride.CreatedAt)
		seconds     = duration.Seconds()
		minutes     = int(math.Ceil(seconds / 60))
		minutePrice = utils.GetEnvAsInt("RIDE_MINUTE_PRICE", 100)
	)

	count, _ := c.repository.Count(ctx, "rides", rel.And(rel.Eq("id", ride.ID), rel.Eq("finished", true)))
	if count != 0 {
		return ErrRideAlreadyFinished, nil
	}

	ride.Price = ride.Price + minutes*minutePrice
	ride.Finished = true

	err := c.repository.Update(ctx, ride)
	return err, ride
}
