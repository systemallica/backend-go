package rides

import (
	"context"

	"github.com/go-rel/rel"
)

type Service interface {
	StartRide(ctx context.Context, ride *Ride) (error, *Ride)
}

type service struct {
	startRide
}

func New(repository rel.Repository) Service {
	return service{
		startRide: startRide{repository: repository},
	}
}