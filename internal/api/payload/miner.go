package payload

type MinerBuyRequest struct {
	Class string `json:"class" validate:"required"`
}

type MinerGetPricesResponse struct {
	Class     string `json:"class"`
	Price     int    `json:"price"`
	Power     int    `json:"power"`
	Energy    int    `json:"energy"`
	BreakTime string `json:"breakTime"`
	Progress  int    `json:"progress"`
}
