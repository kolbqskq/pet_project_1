package payload

import "MinersGame/internal/domain/save"

type GetSaveInfoResponse struct {
	Saves []save.SaveInfo `json:"saves"`
}
