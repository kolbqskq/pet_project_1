package handler

import (
	"MinersGame/internal/api/payload"
	"MinersGame/internal/usecase/enterprise"
	"MinersGame/internal/usecase/stats"
	"MinersGame/pkg/res"
	"context"
	"net/http"
	"time"
)

type EnterpriseHandler struct {
	enterprise.EnterpriseManager
	enterprise.BallanceManager
	stats.StatsProvider
	Ctx       context.Context
	Cancel    context.CancelFunc
	TimeStart *time.Time
}

type EnterpriseHandlerDeps struct {
	enterprise.EnterpriseManager
	enterprise.BallanceManager
	stats.StatsProvider
	Ctx       context.Context
	AppCancel context.CancelFunc
	TimeStart *time.Time
}

func NewEnterpriseHandler(router *http.ServeMux, deps EnterpriseHandlerDeps) {
	handler := &EnterpriseHandler{
		EnterpriseManager: deps.EnterpriseManager,
		BallanceManager:   deps.BallanceManager,
		StatsProvider:     deps.StatsProvider,
		Ctx:               deps.Ctx,
		Cancel:            deps.AppCancel,
		TimeStart:         deps.TimeStart,
	}

	router.HandleFunc("POST /enterprise/end", handler.EndGame())
	router.HandleFunc("GET /enterprise/ballance", handler.GetCoal())
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
		res.Json(w, 200, response)

		handler.Cancel()
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
