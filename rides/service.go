package rides

import (
	"context"
	"time"

	"github.com/go-rel/rel"
)

type Service interface {
	StartRide(ctx context.Context, ride *Ride) (error, *Ride)
	FinishRide(ctx context.Context, ride *Ride, now time.Time) (error, *Ride)
}

type service struct {
	startRide
	finishRide
}

func New(repository rel.Repository) Service {
	return service{
		startRide:  startRide{repository: repository},
		finishRide: finishRide{repository: repository},
	}
}
