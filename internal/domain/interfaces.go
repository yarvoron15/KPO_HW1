package domain

type InventoryItem interface {
	GetInventoryNumber() string
	GetName() string
}
type EnergyConsumer interface{ GetDailyEnergyConsumption() float64 }
type RentalVehicle interface {
	InventoryItem
	GetBeginnerScore() int
	IsSuitableForBeginners() bool
}
type ServiceCenter interface {
	Inspect(vehicle RentalVehicle) (bool, string)
}
