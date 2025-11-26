package main

import (
	"MinersGame/internal/app"
	"MinersGame/pkg/db"
	"context"
)

func main() {
	db := db.NewDb()
	ctx := context.Background()
	app := app.New(ctx, db)
	app.StartGame()
	app.RunServer(":8081")
}
