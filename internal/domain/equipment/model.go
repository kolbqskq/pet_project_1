package equipment

type Equipment struct {
	Price int
	Owned bool
}

func NewEquipments() map[string]Equipment {
	return map[string]Equipment{
		"pickaxe": {
			Price: 3000,
			Owned: false,
		},
		"ventilation": {
			Price: 15000,
			Owned: false,
		},
		"trolleys": {
			Price: 50000,
			Owned: false,
		},
	}
}
