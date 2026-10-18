package domain_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"vyshkat/internal/domain"
)

func TestElectricScooterRejectsInvalidScore(t *testing.T) {
	_, err := domain.NewElectricScooter("SC-001", "Test", 1, 0)
	assert.Error(t, err)
}

func TestElectricVehiclesRejectNegativeEnergy(t *testing.T) {
	_, scooterErr := domain.NewElectricScooter("SC-001", "Test", -1, 8)
	_, bikeErr := domain.NewElectricBike("EB-001", "Test", -1, 8)
	assert.Error(t, scooterErr)
	assert.Error(t, bikeErr)
}

func TestChargingStationRejectsNegativeEnergy(t *testing.T) {
	_, err := domain.NewChargingStation("CS-001", -1)
	assert.Error(t, err)
}

func TestDomainObjectsExposeInventoryAndEnergyData(t *testing.T) {
	bike, err := domain.NewElectricBike("EB-001", "Urban", 3.5, 7)
	assert.NoError(t, err)
	regular, err := domain.NewRegularBike("RB-001", "Road", 5)
	assert.NoError(t, err)
	helm := domain.NewHelmet("H-001", "L")
	station, err := domain.NewChargingStation("CS-001", 4)
	assert.NoError(t, err)

	assert.Equal(t, "EB-001", bike.GetInventoryNumber())
	assert.Equal(t, "Electric Bike Urban", bike.GetName())
	assert.Equal(t, 3.5, bike.GetDailyEnergyConsumption())
	assert.True(t, bike.IsSuitableForBeginners())
	assert.Equal(t, "RB-001", regular.GetInventoryNumber())
	assert.False(t, regular.IsSuitableForBeginners())
	assert.Equal(t, "Helmet (L)", helm.GetName())
	assert.Equal(t, "CS-001", station.GetInventoryNumber())
	assert.Equal(t, 4.0, station.GetDailyEnergyConsumption())
}

func TestVehicleAndEquipmentImplementSeparateCapabilities(t *testing.T) {
	scooter, err := domain.NewElectricScooter("SC-001", "Test", 2, 8)
	assert.NoError(t, err)
	bike, err := domain.NewRegularBike("RB-001", "Test", 8)
	assert.NoError(t, err)
	helm := domain.NewHelmet("H-001", "M")
	station, err := domain.NewChargingStation("CS-001", 4)
	assert.NoError(t, err)

	_, scooterConsumesEnergy := interface{}(scooter).(domain.EnergyConsumer)
	_, bikeConsumesEnergy := interface{}(bike).(domain.EnergyConsumer)
	_, helmetConsumesEnergy := interface{}(helm).(domain.EnergyConsumer)
	_, stationConsumesEnergy := interface{}(station).(domain.EnergyConsumer)

	assert.True(t, scooterConsumesEnergy)
	assert.False(t, bikeConsumesEnergy)
	assert.False(t, helmetConsumesEnergy)
	assert.True(t, stationConsumesEnergy)
}
