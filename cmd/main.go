package main

import (
	"MinersGame/internal/app"
	"context"
)

func main() {
	ctx := context.Background()
	app := app.New(ctx)
	app.StartGame()
	app.RunServer(":8081")
}
