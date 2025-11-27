package event

import (
	"sync"
	"time"

	"github.com/gookit/slog"
)

const (
	EventMinerStart = "miner.start"
	EventMinerMined = "miner.mined"
)

type Event struct {
	Type string
	Data any
}

type EventBus struct {
	subs []chan Event
	mu   sync.RWMutex
}

func NewEventBus() *EventBus {
	return &EventBus{
		subs: make([]chan Event, 0),
	}
}

func (e *EventBus) Publish(event Event) {
	e.mu.RLock()
	for _, ch := range e.subs {
		select {
		case ch <- event:
		case <-time.After(time.Millisecond * 10):
			slog.Warn("Event lost:%v", event)
		}
	}
	e.mu.RUnlock()
}

func (e *EventBus) Subscribe() <-chan Event {
	ch := make(chan Event, 1000)
	e.mu.Lock()
	e.subs = append(e.subs, ch)
	e.mu.Unlock()

	return ch
}
