package equipment

import "MinersGame/internal/domain/equipment"

type EquipmentInfoProviderDeps struct {
	EquipmentInfo
}

type EquipmentsPrices struct {
	Equipment string
	Price     int
}

func GetEquipmentsInfo(deps EquipmentInfoProviderDeps) map[string]equipment.Equipment {
	equipments := deps.GetEquipmentsStatusOwn()
	return equipments
}

func GetEquipmentsPrices(deps EquipmentInfoProviderDeps) []EquipmentsPrices {
	equipments := deps.GetEquipmentsStatusOwn()
	response := make([]EquipmentsPrices, 0, len(equipments))
	for k, v := range equipments {
		response = append(response, EquipmentsPrices{
			Equipment: k,
			Price:     v.Price,
		})
	}
	return response
}
