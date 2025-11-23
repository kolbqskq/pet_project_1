package handler

import (
	"MinersGame/internal/api/payload"
	"MinersGame/internal/usecase/equipment"
	"MinersGame/pkg/req"
	"MinersGame/pkg/res"
	"net/http"
)

type EquipmentHandler struct {
	equipment.EquipmentManager
	equipment.BallanceManager
}

type EquipmentHandlerDeps struct {
	equipment.EquipmentManager
	equipment.BallanceManager
}

func NewEquipmentHandler(router *http.ServeMux, deps EquipmentHandlerDeps) {
	handler := &EquipmentHandler{
		EquipmentManager: deps.EquipmentManager,
		BallanceManager:  deps.BallanceManager,
	}

	router.HandleFunc("POST /equipment", handler.BuyEquipment())
	router.HandleFunc("GET /equipment", handler.GetEquipmentsInfo())
	router.HandleFunc("GET /equipment/prices", handler.GetEquipmentsPrices())
}

func (handler *EquipmentHandler) BuyEquipment() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[payload.EquipmentBuyRequest](w, r)
		if err != nil {
			res.Json(w, 0, err)
			return
		}
		if err := equipment.BuyEquipment(equipment.BuyEquipmentDeps{
			Name:            body.Equipment,
			BallanceManager: handler.BallanceManager,
			EquipmentBuy:    handler.EquipmentManager,
		}); err != nil {
			res.Json(w, 0, err)
			return
		}
		res.Json(w, 201, body.Equipment+" is bought")
	}
}

func (handler *EquipmentHandler) GetEquipmentsInfo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		equipments := equipment.GetEquipmentsInfo(equipment.EquipmentInfoProviderDeps{
			EquipmentInfo: handler.EquipmentManager,
		})

		response := payload.EquipmentsInfoResponse{
			Equipments: equipments,
		}

		res.Json(w, 200, response)
	}
}

func (handler *EquipmentHandler) GetEquipmentsPrices() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prices := equipment.GetEquipmentsPrices(equipment.EquipmentInfoProviderDeps{
			EquipmentInfo: handler.EquipmentManager,
		})
		response := payload.EquipmentsPricesResponse{
			Equipments: prices,
		}
		res.Json(w, 200, response)
	}
}
