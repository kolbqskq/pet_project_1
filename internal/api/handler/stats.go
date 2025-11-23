package handler

import (
	"MinersGame/internal/api/payload"
	"MinersGame/internal/usecase/stats"
	"MinersGame/pkg/res"
	"net/http"
)

type StatsHandler struct {
	stats.StatsProvider
}

type StatsHandlerDeps struct {
	stats.StatsProvider
}

func NewStatsHandler(router *http.ServeMux, deps StatsHandlerDeps) {
	handler := &StatsHandler{
		StatsProvider: deps.StatsProvider,
	}

	router.HandleFunc("GET /stats", handler.GetStats())
}

func (handler *StatsHandler) GetStats() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		by := r.URL.Query().Get("by")
		if by == "" {
			stats, err := handler.StatsProvider.GetStats()
			if err != nil {
				res.Json(w, 0, err)
				return
			}
			response := payload.GetStatsResponse{
				Miners: stats,
			}
			res.Json(w, 200, response)
		} else {
			stats, err := handler.GetStatsByClass(by)
			if err != nil {
				res.Json(w, 0, err)
				return
			}
			response := payload.GetStatsResponse{
				Miners: stats,
			}
			res.Json(w, 200, response)
		}

	}
}
