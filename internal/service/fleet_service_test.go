package service_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"vyshkat/internal/domain"
	"vyshkat/internal/service"
)

type stubServiceCenter struct {
	accepted bool
	reason   string
}

func (s stubServiceCenter) Inspect(domain.RentalVehicle) (bool, string) { return s.accepted, s.reason }

func TestFleetServiceAcceptsVehicleAfterInspection(t *testing.T) {
	fleet := service.NewFleetService(stubServiceCenter{true, "accepted"})
	vehicle, err := domain.NewRegularBike("RB-001", "City", 6)
	assert.NoError(t, err)
	accepted, reason := fleet.AddVehicle(vehicle)
	assert.True(t, accepted)
	assert.Equal(t, "accepted", reason)
	assert.Equal(t, 1, fleet.GetVehiclesCount())
	assert.Equal(t, 1, fleet.GetItemsCount())
}

func TestFleetServiceRejectsVehicleAfterInspection(t *testing.T) {
	fleet := service.NewFleetService(stubServiceCenter{false, "unsafe"})
	vehicle, err := domain.NewRegularBike("RB-002", "City", 6)
	assert.NoError(t, err)
	accepted, reason := fleet.AddVehicle(vehicle)
	assert.False(t, accepted)
	assert.Equal(t, "unsafe", reason)
	assert.Zero(t, fleet.GetVehiclesCount())
	assert.Zero(t, fleet.GetItemsCount())
}

func TestFleetServiceRejectsDuplicateInventoryNumber(t *testing.T) {
	fleet := service.NewFleetService(stubServiceCenter{true, "accepted"})
	first, err := domain.NewRegularBike("RB-001", "City", 6)
	assert.NoError(t, err)
	second := domain.NewHelmet("RB-001", "M")

	accepted, _ := fleet.AddVehicle(first)
	assert.True(t, accepted)
	added, reason := fleet.AddItem(second)
	assert.False(t, added)
	assert.Contains(t, reason, "RB-001")
	assert.Equal(t, 1, fleet.GetItemsCount())
}

func TestFleetServiceSumsOnlyEnergyConsumers(t *testing.T) {
	fleet := service.NewFleetService(stubServiceCenter{true, ""})
	scooter, err := domain.NewElectricScooter("SC-001", "Pro", 2.5, 8)
	assert.NoError(t, err)
	bike, err := domain.NewRegularBike("RB-001", "Road", 8)
	assert.NoError(t, err)
	fleet.AddVehicle(scooter)
	fleet.AddVehicle(bike)
	fleet.AddItem(domain.NewHelmet("H-001", "M"))
	station, err := domain.NewChargingStation("CS-001", 4)
	assert.NoError(t, err)
	fleet.AddItem(station)
	assert.Equal(t, 6.5, fleet.GetTotalDailyEnergy())
}

func TestFleetServiceFiltersBeginnerVehiclesByScore(t *testing.T) {
	fleet := service.NewFleetService(stubServiceCenter{true, ""})
	for i, score := range []int{5, 6, 9} {
		vehicle, err := domain.NewRegularBike("RB-00"+strconv.Itoa(i), "Test", score)
		assert.NoError(t, err)
		fleet.AddVehicle(vehicle)
	}
	vehicles := fleet.GetVehiclesForBeginners()
	assert.Len(t, vehicles, 2)
	assert.Equal(t, 6, vehicles[0].GetBeginnerScore())
	assert.Equal(t, 9, vehicles[1].GetBeginnerScore())
}

func TestFleetServiceListsCompleteInventory(t *testing.T) {
	fleet := service.NewFleetService(stubServiceCenter{true, ""})
	vehicle, err := domain.NewRegularBike("RB-001", "City", 6)
	assert.NoError(t, err)
	fleet.AddVehicle(vehicle)
	fleet.AddItem(domain.NewHelmet("H-001", "M"))
	items := fleet.GetAllItems()
	assert.Len(t, items, 2)
	assert.Equal(t, "RB-001", items[0].GetInventoryNumber())
	assert.Equal(t, "H-001", items[1].GetInventoryNumber())
}
