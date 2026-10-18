package domain

import "fmt"

type Helmet struct{ inventoryNumber, size string }

func NewHelmet(number, size string) *Helmet  { return &Helmet{number, size} }
func (h *Helmet) GetInventoryNumber() string { return h.inventoryNumber }
func (h *Helmet) GetName() string            { return "Helmet (" + h.size + ")" }

type ChargingStation struct {
	inventoryNumber string
	dailyEnergy     float64
}

func NewChargingStation(number string, energy float64) (*ChargingStation, error) {
	if energy < 0 {
		return nil, fmt.Errorf("daily energy cannot be negative")
	}
	return &ChargingStation{number, energy}, nil
}
func (s *ChargingStation) GetInventoryNumber() string         { return s.inventoryNumber }
func (s *ChargingStation) GetName() string                    { return "Charging Station" }
func (s *ChargingStation) GetDailyEnergyConsumption() float64 { return s.dailyEnergy }
