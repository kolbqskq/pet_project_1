package db

import (
	"github.com/gookit/slog"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct {
	*gorm.DB
}

func NewDb() *Db {
	dsn := "host=localhost user=postgres password=my_pass dbname=saves port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}
	return &Db{db}
}
