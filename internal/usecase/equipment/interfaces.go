package equipment

import "MinersGame/internal/domain/equipment"

type EquipmentBuy interface {
	Buy(name string) error
	GetPrice(name string) int
}

type BalanceManager interface {
	SpendCoal(amount int) error
}

type EquipmentInfo interface {
	GetEquipmentsStatusOwn() map[string]equipment.Equipment
}

type EquipmentManager interface {
	EquipmentInfo
	EquipmentBuy
}
