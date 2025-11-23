package config

import (
	"MinersGame/internal/types"
	"time"
)

type MinerConfig struct {
	Class     string
	Price     types.Coal
	Power     int
	Energy    int
	BreakTime time.Duration
	Progress  int
}

var MinerPresets = map[string]MinerConfig{
	"small": {
		Class:     "small",
		Price:     5,
		Energy:    30,
		Power:     1,
		BreakTime: 3 * time.Second,
		Progress:  0,
	},
	"normal": {
		Class:     "normal",
		Price:     50,
		Energy:    45,
		Power:     3,
		BreakTime: 2 * time.Second,
		Progress:  0,
	},
	"strong": {
		Class:     "strong",
		Price:     450,
		Energy:    60,
		Power:     10,
		BreakTime: 1 * time.Second,
		Progress:  3,
	},
}

var minerList = []string{"small", "normal", "strong"}

func GetMinersPrices() []MinerConfig {
	response := make([]MinerConfig, 0, len(minerList))
	for _, v := range minerList {
		if cfg, ok := MinerPresets[v]; ok {
			response = append(response, cfg)
		}

	}
	return response
}
