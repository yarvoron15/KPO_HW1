package service

import (
	"strings"
	"vyshkat/internal/domain"
)

type StandardServiceCenter struct{}

func NewStandardServiceCenter() *StandardServiceCenter { return &StandardServiceCenter{} }
func (s *StandardServiceCenter) Inspect(vehicle domain.RentalVehicle) (bool, string) {
	if vehicle == nil {
		return false, "vehicle is nil"
	}
	if strings.Contains(strings.ToUpper(vehicle.GetInventoryNumber()), "BROKEN") {
		return false, "inventory number contains BROKEN"
	}
	if vehicle.GetInventoryNumber() == "" {
		return false, "inventory number is empty"
	}
	if vehicle.GetBeginnerScore() <= 0 {
		return false, "beginner score must be positive"
	}
	return true, "vehicle passed technical inspection"
}
