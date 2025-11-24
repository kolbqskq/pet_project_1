package miner

import (
	"MinersGame/internal/types"
	"MinersGame/pkg/event"
	"context"
	"time"
)

type MinerService struct {
	EventBus *event.EventBus
}

type MinerServiceDeps struct {
	EventBus *event.EventBus
}

func NewMinerService(deps MinerServiceDeps) *MinerService {
	return &MinerService{
		EventBus: deps.EventBus,
	}
}

func (s *MinerService) Mine(ctx context.Context, m *Miner) {

	ticker := time.NewTicker(m.Config.BreakTime)

	go func() {
		defer ticker.Stop()

		s.EventBus.Publish(event.Event{
			Type: event.EventMinerStart,
			Data: m,
		})

		coalPerTick := m.Config.Power

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				m.mu.Lock()
				if m.EnergyLeft == 0 {
					m.mu.Unlock()
					return
				}
				m.EnergyLeft--
				m.mu.Unlock()

				s.EventBus.Publish(event.Event{
					Type: event.EventMinerMined,
					Data: types.Coal(coalPerTick),
				})
			}
			coalPerTick += m.Config.Progress

		}
	}()

}
