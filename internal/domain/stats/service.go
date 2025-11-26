package stats

import (
	"MinersGame/internal/domain/miner"
	"MinersGame/pkg/errs"
	"MinersGame/pkg/event"
	"sync"
)

type StatsService struct {
	StatsMiners
	EventBus *event.EventBus
	mu       sync.RWMutex
}

type StatsServiceDeps struct {
	EventBus *event.EventBus
}

func NewStatsMinersService(deps StatsServiceDeps) *StatsService {
	s := &StatsService{
		StatsMiners: StatsMiners{
			StatMiners: make(map[string][]*miner.Miner),
		},
		EventBus: deps.EventBus,
	}
	s.startListening()
	return s
}

func (s *StatsService) startListening() {
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
			s.mu.Lock()
			if _, ok := s.StatMiners[class]; !ok {
				s.StatMiners[class] = make([]*miner.Miner, 0)
			}

			s.StatMiners[class] = append(s.StatMiners[class], m)

			s.mu.Unlock()
		}
	}()
}

func (s *StatsService) GetStats() ([]MinerInfo, error) {
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

func (s *StatsService) GetStatsByClass(class string) ([]MinerInfo, error) {
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

func (s *StatsService) GetCountsAllClass() []CountsMiners {
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

func (s *StatsService) Load() {
	s.mu.Lock()
	s.StatMiners = make(map[string][]*miner.Miner)
	s.mu.Unlock()
}
