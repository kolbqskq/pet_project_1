package equipment

type BuyEquipmentDeps struct {
	Name string
	BalanceManager
	EquipmentBuy
}

func BuyEquipment(deps BuyEquipmentDeps) error {
	price := deps.EquipmentBuy.GetPrice(deps.Name)
	if err := deps.BalanceManager.SpendCoal(price); err != nil {
		return err
	}
	if err := deps.EquipmentBuy.Buy(deps.Name); err != nil {
		return err
	}
	return nil
}
