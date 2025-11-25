package app

import (
	"MinersGame/internal/types"
	"MinersGame/pkg/event"
	"time"
)

func (app *App) StartGame() {
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for {
			select {
			case <-app.AppCtx.Done():
				return
			case <-ticker.C:
				app.EventBus.Publish(event.Event{
					Type: event.EventMinerMined,
					Data: types.Coal(1),
				})
			}
		}

	}()

}
