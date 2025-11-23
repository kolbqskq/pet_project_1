package stats

import (
	"MinersGame/internal/domain/miner"
	"MinersGame/pkg/errs"
	"MinersGame/pkg/event"
	"sync"
)

type StatsMinersService struct {
	StatsMiners
	EventBus *event.EventBus
	mu       sync.RWMutex
}

type StatsMinersServiceDeps struct {
	EventBus *event.EventBus
}

func NewStatsMinersService(deps StatsMinersServiceDeps) *StatsMinersService {
	s := &StatsMinersService{
		StatsMiners: StatsMiners{
			StatMiners: make(map[string]map[string]*miner.Miner),
		},
		EventBus: deps.EventBus,
	}
	s.startListening()
	return s
}

func (s *StatsMinersService) startListening() {
	go func() {

		for msg := range s.EventBus.Subscribe() {
			if msg.Type != event.EventMinerStart {
				continue
			}
			m, ok := msg.Data.(*miner.Miner)
			if !ok {
				continue
			}
			class := m.Config.Class
			id := m.ID
			s.mu.Lock()
			if _, ok := s.StatMiners[class]; !ok {
				s.StatMiners[class] = make(map[string]*miner.Miner)
			}

			s.StatMiners[class][id] = m

			s.mu.Unlock()
		}
	}()
}

func (s *StatsMinersService) GetStats() ([]MinerInfo, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	response := []MinerInfo{}
	for _, miners := range s.StatMiners {
		for _, miner := range miners {
			response = append(response, MinerInfo{
				Class:      miner.Config.Class,
				EnergyLeft: miner.EnergyLeft,
			})
		}
	}
	if len(response) == 0 {
		return nil, errs.NewHaveNotMiners() //error
	}
	return response, nil
}

func (s *StatsMinersService) GetStatsByClass(class string) ([]MinerInfo, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if len(s.StatMiners[class]) == 0 {
		errMsg := "you have not " + class + " miners"

		return nil, errs.NewCustomErr(409, errMsg) //error
	}

	response := []MinerInfo{}
	for _, miner := range s.StatMiners[class] {
		response = append(response, MinerInfo{
			Class:      miner.Config.Class,
			EnergyLeft: miner.EnergyLeft,
		})

	}

	return response, nil
}

func (s *StatsMinersService) GetCountsAllClass() []CountsMiners {
	s.mu.RLock()
	defer s.mu.RUnlock()
	response := []CountsMiners{}
	for class, miner := range s.StatMiners {
		response = append(response, CountsMiners{
			Class: class,
			Count: len(miner),
		})
	}
	return response
}
