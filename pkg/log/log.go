package main

import (
	"github.com/gookit/slog"
	"github.com/gookit/slog/handler"
)

func main() {
	defer slog.MustClose()
	h := handler.MustFileHandler("error.log", handler.WithLogLevels(slog.DangerLevels))
	slog.PushHandler(h)
}
