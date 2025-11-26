package app

import (
	"MinersGame/internal/domain/equipment"
	"MinersGame/internal/domain/miner"
	"MinersGame/internal/domain/save"
	"MinersGame/internal/domain/stats"
	"MinersGame/internal/domain/wallet"
	"MinersGame/pkg/db"
	"MinersGame/pkg/event"
	"context"
	"net/http"
	"time"
)

type App struct {
	Server   *http.Server
	EventBus *event.EventBus
	Db       *db.Db

	MinerService     *miner.MinerService
	WalletService    *wallet.WalletService
	StatsService     *stats.StatsService
	EquipmentService *equipment.EquipmentService

	AppCtx    context.Context
	AppCancel context.CancelFunc

	MinersCtx    context.Context
	MinersCancel context.CancelFunc

	TimeStart *time.Time
}

func New(ctx context.Context, db *db.Db) App {
	timeStart := time.Now()
	appCtx, appCancel := context.WithCancel(ctx)
	minersCtx, minersCancel := context.WithCancel(appCtx)

	eventBus := event.NewEventBus()

	//Services:

	minerService := miner.NewMinerService(miner.MinerServiceDeps{
		EventBus: eventBus,
	})

	walletService := wallet.NewWalletService(wallet.ServiceWalletDeps{
		EventBus: eventBus,
	},
	)

	statsService := stats.NewStatsMinersService(stats.StatsServiceDeps{
		EventBus: eventBus,
	},
	)

	equipmentService := equipment.NewEquipmentService()

	return App{
		Server:   &http.Server{},
		EventBus: eventBus,
		Db:       db,

		MinerService:     minerService,
		WalletService:    walletService,
		StatsService:     statsService,
		EquipmentService: equipmentService,

		AppCtx:    appCtx,
		AppCancel: appCancel,

		MinersCtx:    minersCtx,
		MinersCancel: minersCancel,

		TimeStart: &timeStart,
	}
}

func (a *App) LoadSave(save *save.GameSave) {
	a.MinersCancel()

	minersCtx, minersCancel := context.WithCancel(a.AppCtx)
	a.MinersCtx = minersCtx
	a.MinersCancel = minersCancel

	a.StatsService.Load()
	a.WalletService.Load(save.Wallet)
	a.EquipmentService.Load(save.Equipments)
	for _, v := range save.Miners {
		miner, _ := miner.NewMiner(v.Class)
		miner.EnergyLeft = v.EnergyLeft
		a.MinerService.Mine(minersCtx, miner)
	}
}
