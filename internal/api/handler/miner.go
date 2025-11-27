package handler

import (
	"MinersGame/internal/api/payload"
	"MinersGame/internal/config"
	"MinersGame/internal/usecase/miner"
	"MinersGame/pkg/req"
	"MinersGame/pkg/res"
	"context"
	"net/http"
)

type MinersHandler struct {
	Ctx context.Context
	miner.MinerManager
	miner.BalanceManager
}

type MinersHandlerDeps struct {
	Ctx context.Context
	miner.MinerManager
	miner.BalanceManager
}

func NewMinersHandler(router *http.ServeMux, deps MinersHandlerDeps) {
	handler := &MinersHandler{
		Ctx:            deps.Ctx,
		MinerManager:   deps.MinerManager,
		BalanceManager: deps.BalanceManager,
	}

	router.HandleFunc("POST /miners", handler.BuyMiner())
	router.HandleFunc("GET /miners/prices", handler.GetMinersPrices())
}

func (handler *MinersHandler) BuyMiner() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[payload.MinerBuyRequest](w, r)
		if err != nil {
			res.Json(w, 0, err)
			return
		}
		if err := miner.BuyMiner(miner.BuyMinerDeps{
			MinerManager:   handler.MinerManager,
			BalanceManager: handler.BalanceManager,
			Class:          body.Class,
			Ctx:            handler.Ctx,
		}); err != nil {
			res.Json(w, 0, err)
			return
		}
		res.Json(w, 201, body.Class+" miner is bought")

	}
}

func (handler *MinersHandler) GetMinersPrices() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		minersList := config.GetMinersPrices()
		response := make([]payload.MinerGetPricesResponse, 0, len(minersList))
		for _, v := range minersList {
			response = append(response, payload.MinerGetPricesResponse{
				Class:     v.Class,
				Price:     int(v.Price),
				Power:     v.Power,
				Energy:    v.Energy,
				BreakTime: v.BreakTime.String(),
				Progress:  v.Progress,
			})
		}
		res.Json(w, 200, response)
	}
}
