package service_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"vyshkat/internal/domain"
	"vyshkat/internal/service"
)

func TestStandardServiceCenterAcceptsValidVehicle(t *testing.T) {
	center := service.NewStandardServiceCenter()
	vehicle, err := domain.NewRegularBike("RB-001", "City", 6)
	assert.NoError(t, err)
	accepted, reason := center.Inspect(vehicle)
	assert.True(t, accepted)
	assert.Equal(t, "vehicle passed technical inspection", reason)
}

func TestStandardServiceCenterRejectsBrokenVehicle(t *testing.T) {
	center := service.NewStandardServiceCenter()
	vehicle, err := domain.NewRegularBike("BROKEN-001", "City", 6)
	assert.NoError(t, err)
	accepted, reason := center.Inspect(vehicle)
	assert.False(t, accepted)
	assert.Equal(t, "inventory number contains BROKEN", reason)
}
