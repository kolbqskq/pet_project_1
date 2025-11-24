package handler

import (
	"MinersGame/internal/api/payload"
	"MinersGame/internal/usecase/enterprise"
	"MinersGame/internal/usecase/stats"
	"MinersGame/pkg/res"
	"context"
	"log"
	"net/http"
	"time"
)

type EnterpriseHandler struct {
	enterprise.EnterpriseManager
	enterprise.BallanceManager
	stats.StatsProvider
	MinersCancel context.CancelFunc
	Ctx          context.Context
	Server       *http.Server
	TimeStart    *time.Time
}

type EnterpriseHandlerDeps struct {
	enterprise.EnterpriseManager
	enterprise.BallanceManager
	stats.StatsProvider
	MinersCancel context.CancelFunc
	Ctx          context.Context
	Server       *http.Server
	TimeStart    *time.Time
}

func NewEnterpriseHandler(router *http.ServeMux, deps EnterpriseHandlerDeps) {
	handler := &EnterpriseHandler{
		EnterpriseManager: deps.EnterpriseManager,
		BallanceManager:   deps.BallanceManager,
		StatsProvider:     deps.StatsProvider,
		MinersCancel:      deps.MinersCancel,
		Ctx:               deps.Ctx,
		Server:            deps.Server,
		TimeStart:         deps.TimeStart,
	}

	router.HandleFunc("GET /enterprise/end", handler.EndGame())
	router.HandleFunc("GET /enterprise/balance", handler.GetCoal())
}

func (handler *EnterpriseHandler) EndGame() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := handler.ValidateEndGame(); err != nil {
			res.Json(w, 0, err)
			return
		}
		miners, err := handler.GetStats()
		if err != nil {
			res.Json(w, 0, err)
			return
		}
		response := payload.EndGameResponse{
			Balance:      handler.GetBallance(),
			Miners:       miners,
			CountMiners:  handler.GetCountsAllClass(),
			GameDuration: time.Since(*handler.TimeStart).String(),
		}
		handler.MinersCancel()
		res.Json(w, 200, response)

		go func() {
			if err := handler.Server.Shutdown(handler.Ctx); err != nil {
				log.Printf("Server shutdown error: %v", err)
			}
		}()
	}
}

func (handler *EnterpriseHandler) GetCoal() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := payload.GetBalanceResponse{
			Coal: handler.GetBallance(),
		}
		res.Json(w, 200, response)
	}
}
