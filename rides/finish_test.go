package rides

import (
	"context"
	"testing"
	"time"

	"github.com/go-rel/rel"
	"github.com/go-rel/reltest"
	"github.com/stretchr/testify/assert"
)

func TestFinish1Second(t *testing.T) {
	var (
		ctx        = context.TODO()
		repository = reltest.New()
		service    = New(repository)
		now        = time.Now()
		createdAt  = now.Add(-time.Second * 1)
		ride       = Ride{ID: 1, UserID: "1", VehicleID: "1", Price: 18, CreatedAt: createdAt, UpdatedAt: createdAt, Finished: false}
	)

	repository.ExpectCount("rides", rel.And(rel.Eq("id", ride.ID), rel.Eq("finished", true))).Result(0)
	repository.ExpectUpdate().For(&ride)

	err, _ := service.FinishRide(ctx, &ride, now)
	assert.Nil(t, err)
	assert.Equal(t, ride.Price, 118)
	assert.NotEmpty(t, ride.UpdatedAt)
	assert.True(t, ride.Finished)

	repository.AssertExpectations(t)
}

func TestFinish10Seconds(t *testing.T) {
	var (
		ctx        = context.TODO()
		repository = reltest.New()
		service    = New(repository)
		now        = time.Now()
		createdAt  = now.Add(-time.Second * 10)
		ride       = Ride{ID: 1, UserID: "1", VehicleID: "1", Price: 18, CreatedAt: createdAt, UpdatedAt: createdAt, Finished: false}
	)

	repository.ExpectCount("rides", rel.And(rel.Eq("id", ride.ID), rel.Eq("finished", true))).Result(0)
	repository.ExpectUpdate().For(&ride)

	err, _ := service.FinishRide(ctx, &ride, now)
	assert.Nil(t, err)
	assert.Equal(t, ride.Price, 118)
	assert.True(t, ride.Finished)

	repository.AssertExpectations(t)
}

func TestFinish59Seconds(t *testing.T) {
	var (
		ctx        = context.TODO()
		repository = reltest.New()
		service    = New(repository)
		now        = time.Now()
		createdAt  = now.Add(-time.Second * 59)
		ride       = Ride{ID: 1, UserID: "1", VehicleID: "1", Price: 18, CreatedAt: createdAt, UpdatedAt: createdAt, Finished: false}
	)

	repository.ExpectCount("rides", rel.And(rel.Eq("id", ride.ID), rel.Eq("finished", true))).Result(0)
	repository.ExpectUpdate().For(&ride)

	err, _ := service.FinishRide(ctx, &ride, now)
	assert.Nil(t, err)
	assert.Equal(t, ride.Price, 118)
	assert.True(t, ride.Finished)

	repository.AssertExpectations(t)
}

func TestFinish60Seconds(t *testing.T) {
	var (
		ctx        = context.TODO()
		repository = reltest.New()
		service    = New(repository)
		now        = time.Now()
		createdAt  = now.Add(-time.Second * 60)
		ride       = Ride{ID: 1, UserID: "1", VehicleID: "1", Price: 18, CreatedAt: createdAt, UpdatedAt: createdAt, Finished: false}
	)

	repository.ExpectCount("rides", rel.And(rel.Eq("id", ride.ID), rel.Eq("finished", true))).Result(0)
	repository.ExpectUpdate().For(&ride)

	err, _ := service.FinishRide(ctx, &ride, now)
	assert.Nil(t, err)
	assert.Equal(t, ride.Price, 118)
	assert.True(t, ride.Finished)

	repository.AssertExpectations(t)
}

func TestFinish61Seconds(t *testing.T) {
	var (
		ctx        = context.TODO()
		repository = reltest.New()
		service    = New(repository)
		now        = time.Now()
		createdAt  = now.Add(-time.Second * 61)
		ride       = Ride{ID: 1, UserID: "1", VehicleID: "1", Price: 18, CreatedAt: createdAt, UpdatedAt: createdAt, Finished: false}
	)

	repository.ExpectCount("rides", rel.And(rel.Eq("id", ride.ID), rel.Eq("finished", true))).Result(0)
	repository.ExpectUpdate().For(&ride)

	err, _ := service.FinishRide(ctx, &ride, now)
	assert.Nil(t, err)
	assert.Equal(t, ride.Price, 218)
	assert.True(t, ride.Finished)

	repository.AssertExpectations(t)
}

func TestFinish1565Seconds(t *testing.T) {
	var (
		ctx        = context.TODO()
		repository = reltest.New()
		service    = New(repository)
		now        = time.Now()
		createdAt  = now.Add(-time.Second * 1565)
		ride       = Ride{ID: 1, UserID: "1", VehicleID: "1", Price: 18, CreatedAt: createdAt, UpdatedAt: createdAt, Finished: false}
	)

	repository.ExpectCount("rides", rel.And(rel.Eq("id", ride.ID), rel.Eq("finished", true))).Result(0)
	repository.ExpectUpdate().For(&ride)

	err, _ := service.FinishRide(ctx, &ride, now)
	assert.Nil(t, err)
	assert.Equal(t, ride.Price, 2718)
	assert.True(t, ride.Finished)

	repository.AssertExpectations(t)
}

func TestFinishRideAlreadyFinished(t *testing.T) {
	var (
		ctx        = context.TODO()
		repository = reltest.New()
		service    = New(repository)
		now        = time.Now()
		createdAt  = now.Add(-time.Second * 1565)
		ride       = Ride{ID: 1, UserID: "1", VehicleID: "1", Price: 18, CreatedAt: createdAt, UpdatedAt: createdAt, Finished: false}
	)

	repository.ExpectCount("rides", rel.And(rel.Eq("id", ride.ID), rel.Eq("finished", true))).Result(1)

	err, _ := service.FinishRide(ctx, &ride, now)
	assert.Equal(t, ErrRideAlreadyFinished, err)

	repository.AssertExpectations(t)
}
