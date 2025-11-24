package wallet

import (
	"MinersGame/internal/types"
	"MinersGame/pkg/errs"
	"MinersGame/pkg/event"
	"sync"
)

type WalletService struct {
	Wallet
	EventBus *event.EventBus
	mu       sync.RWMutex
}

type ServiceWalletDeps struct {
	EventBus *event.EventBus
}

func NewWalletService(deps ServiceWalletDeps) *WalletService {
	s := &WalletService{
		Wallet:   Wallet{},
		EventBus: deps.EventBus,
	}
	s.startListening()
	return s
}

func (s *WalletService) startListening() {
	go func() {
		for msg := range s.EventBus.Subscribe() {
			if msg.Type == event.EventMinerMined {
				coal, ok := msg.Data.(types.Coal)
				if !ok {
					continue
				}
				s.mu.Lock()
				s.Coal += coal
				s.mu.Unlock()
			}
		}
	}()
}

func (s *WalletService) GetBallance() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return int(s.Coal)
}

func (s *WalletService) validateBalance(price int) error {
	if int(s.Coal) < price {
		return errs.NewNotEnoughBalance()
	}
	return nil
}

func (s *WalletService) SpendCoal(amount int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if err := s.validateBalance(amount); err != nil {
		return err
	}
	s.Coal -= types.Coal(amount)
	return nil
}
