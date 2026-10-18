package domain

import "fmt"

func validateScore(score int) error {
	if score < 1 || score > 10 {
		return fmt.Errorf("beginner score must be between 1 and 10, got %d", score)
	}
	return nil
}
func validateEnergy(energy float64) error {
	if energy < 0 {
		return fmt.Errorf("daily energy cannot be negative")
	}
	return nil
}

type ElectricScooter struct {
	inventoryNumber, model string
	dailyEnergy            float64
	beginnerScore          int
}

func NewElectricScooter(number, model string, energy float64, score int) (*ElectricScooter, error) {
	if err := validateScore(score); err != nil {
		return nil, err
	}
	if err := validateEnergy(energy); err != nil {
		return nil, err
	}
	return &ElectricScooter{number, model, energy, score}, nil
}
func (s *ElectricScooter) GetInventoryNumber() string         { return s.inventoryNumber }
func (s *ElectricScooter) GetName() string                    { return "Electric Scooter " + s.model }
func (s *ElectricScooter) GetDailyEnergyConsumption() float64 { return s.dailyEnergy }
func (s *ElectricScooter) GetBeginnerScore() int              { return s.beginnerScore }
func (s *ElectricScooter) IsSuitableForBeginners() bool       { return s.beginnerScore >= 6 }

type ElectricBike struct {
	inventoryNumber, model string
	dailyEnergy            float64
	beginnerScore          int
}

func NewElectricBike(number, model string, energy float64, score int) (*ElectricBike, error) {
	if err := validateScore(score); err != nil {
		return nil, err
	}
	if err := validateEnergy(energy); err != nil {
		return nil, err
	}
	return &ElectricBike{number, model, energy, score}, nil
}
func (b *ElectricBike) GetInventoryNumber() string         { return b.inventoryNumber }
func (b *ElectricBike) GetName() string                    { return "Electric Bike " + b.model }
func (b *ElectricBike) GetDailyEnergyConsumption() float64 { return b.dailyEnergy }
func (b *ElectricBike) GetBeginnerScore() int              { return b.beginnerScore }
func (b *ElectricBike) IsSuitableForBeginners() bool       { return b.beginnerScore >= 6 }

type RegularBike struct {
	inventoryNumber, model string
	beginnerScore          int
}

func NewRegularBike(number, model string, score int) (*RegularBike, error) {
	if err := validateScore(score); err != nil {
		return nil, err
	}
	return &RegularBike{number, model, score}, nil
}
func (b *RegularBike) GetInventoryNumber() string   { return b.inventoryNumber }
func (b *RegularBike) GetName() string              { return "Regular Bike " + b.model }
func (b *RegularBike) GetBeginnerScore() int        { return b.beginnerScore }
func (b *RegularBike) IsSuitableForBeginners() bool { return b.beginnerScore >= 6 }
