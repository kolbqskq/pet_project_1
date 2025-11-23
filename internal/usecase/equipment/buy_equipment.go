package equipment

type BuyEquipmentDeps struct {
	Name string
	BallanceManager
	EquipmentBuy
}

func BuyEquipment(deps BuyEquipmentDeps) error {
	price := deps.EquipmentBuy.GetPrice(deps.Name)
	if err := deps.BallanceManager.SpendCoal(price); err != nil {
		return err
	}
	if err := deps.EquipmentBuy.Buy(deps.Name); err != nil {
		return err
	}
	return nil
}
