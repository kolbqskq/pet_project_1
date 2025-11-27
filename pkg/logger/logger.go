package logger

import (
	"github.com/gookit/slog"
	"github.com/gookit/slog/handler"
)

func Init() {
	h := handler.MustFileHandler("error.log", handler.WithLogLevels(slog.DangerLevels))
	slog.PushHandler(h)
	h2 := handler.MustFileHandler("info.log", handler.WithLogLevels(slog.NormalLevels))
	slog.PushHandler(h2)
}
