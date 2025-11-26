package equipment

import (
	"MinersGame/pkg/errs"
	"sync"
)

type EquipmentService struct {
	Equipments map[string]Equipment
	mu         sync.RWMutex
}

/* type EquipmentServiceDeps struct {
} */

func NewEquipmentService() *EquipmentService {
	return &EquipmentService{
		Equipments: NewEquipments(),
	}
}

func (e *EquipmentService) isOwned(name string) error {
	eq, ok := e.Equipments[name]
	if !ok {
		return errs.NewInvalidEquipment()
	}
	if eq.Owned {
		return errs.NewEquipmentAlreadyOwn()
	}
	return nil
}

func (e *EquipmentService) Buy(name string) error {
	e.mu.Lock()
	defer e.mu.Unlock()
	if err := e.isOwned(name); err != nil {
		return err
	}
	eq := e.Equipments[name]
	eq.Owned = true
	e.Equipments[name] = eq
	return nil
}

func (e *EquipmentService) ValidateEndGame() error {
	e.mu.RLock()
	defer e.mu.RUnlock()
	for _, v := range e.Equipments {
		if !v.Owned {
			return errs.NewYouShouldBuyAllTools()
		}
	}
	return nil
}

func (e *EquipmentService) GetEquipmentsStatusOwn() map[string]Equipment {
	e.mu.RLock()
	defer e.mu.RUnlock()
	response := make(map[string]Equipment, len(e.Equipments))
	for k, v := range e.Equipments {
		response[k] = v
	}
	return response
}

func (e *EquipmentService) GetPrice(name string) int {
	e.mu.RLock()
	defer e.mu.RUnlock()
	price := e.Equipments[name].Price
	return price
}

func (e *EquipmentService) Load(equipments map[string]Equipment) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.Equipments = equipments
}
