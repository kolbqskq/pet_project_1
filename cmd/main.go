package main

import (
	"MinersGame/internal/app"
	"MinersGame/pkg/db"
	"MinersGame/pkg/logger"
	"context"

	"github.com/gookit/slog"
)

func main() {
	logger.Init()
	defer slog.MustClose()
	db := db.NewDb()
	ctx := context.Background()
	app := app.New(ctx, db)
	app.StartGame()
	app.RunServer(":8081")
}
