package payload

import (
	"MinersGame/internal/usecase/equipment"
)

type EquipmentBuyRequest struct {
	Equipment string `json:"equipment" validate:"required"`
}

type EquipmentsPricesResponse struct {
	Equipments []equipment.EquipmentsPrices `json:"equipments"`
}

type EquipmentsInfoResponse struct {
	Equipments any `json:"equipments"`
}
