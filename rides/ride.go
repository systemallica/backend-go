package rides

import (
	"errors"
	"time"
)

type Ride struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Price     int       `json:"price"`
	UserID    string    `json:"user_id"`
	VehicleID string    `json:"vehicle_id"`
	Finished  bool      `json:"finished"`
}

var (
	ErrRideUserIDBlank    = errors.New("UserID can't be blank")
	ErrRideVehicleIDBlank = errors.New("VehicleID can't be blank")
	ErrRideAlreadyStarted = errors.New("A ride is already started for this vehicle or user")
)

func (r Ride) Validate() error {
	var err error
	switch {
	case r.UserID == "":
		err = ErrRideUserIDBlank
	case r.VehicleID == "":
		err = ErrRideVehicleIDBlank
	}

	return err
}
