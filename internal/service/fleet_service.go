package service

import (
	"fmt"

	"vyshkat/internal/domain"
)

type FleetService struct {
	serviceCenter domain.ServiceCenter
	items         []domain.InventoryItem
	vehicles      []domain.RentalVehicle
	knownNumbers  map[string]struct{}
}

func NewFleetService(sc domain.ServiceCenter) *FleetService {
	return &FleetService{serviceCenter: sc, items: make([]domain.InventoryItem, 0), vehicles: make([]domain.RentalVehicle, 0), knownNumbers: make(map[string]struct{})}
}
func (f *FleetService) AddVehicle(v domain.RentalVehicle) (bool, string) {
	if v == nil {
		return false, "vehicle is nil"
	}
	if _, exists := f.knownNumbers[v.GetInventoryNumber()]; exists {
		return false, fmt.Sprintf("предмет с инвентарным номером %s уже существует", v.GetInventoryNumber())
	}
	accepted, reason := f.serviceCenter.Inspect(v)
	if !accepted {
		return false, reason
	}
	f.knownNumbers[v.GetInventoryNumber()] = struct{}{}
	f.vehicles = append(f.vehicles, v)
	f.items = append(f.items, v)
	return true, reason
}
func (f *FleetService) AddItem(item domain.InventoryItem) (bool, string) {
	if item == nil {
		return false, "item is nil"
	}
	if _, exists := f.knownNumbers[item.GetInventoryNumber()]; exists {
		return false, fmt.Sprintf("предмет с инвентарным номером %s уже существует", item.GetInventoryNumber())
	}
	f.knownNumbers[item.GetInventoryNumber()] = struct{}{}
	f.items = append(f.items, item)
	return true, "item added to inventory"
}
func (f *FleetService) GetTotalDailyEnergy() float64 {
	var total float64
	for _, item := range f.items {
		if consumer, ok := item.(domain.EnergyConsumer); ok {
			total += consumer.GetDailyEnergyConsumption()
		}
	}
	return total
}
func (f *FleetService) GetVehiclesForBeginners() []domain.RentalVehicle {
	result := make([]domain.RentalVehicle, 0)
	for _, v := range f.vehicles {
		if v.IsSuitableForBeginners() {
			result = append(result, v)
		}
	}
	return result
}
func (f *FleetService) GetAllItems() []domain.InventoryItem {
	result := make([]domain.InventoryItem, len(f.items))
	copy(result, f.items)
	return result
}
func (f *FleetService) GetVehiclesCount() int { return len(f.vehicles) }
func (f *FleetService) GetItemsCount() int    { return len(f.items) }
