package main

import (
	"MinersGame/internal/domain/save"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=my_pass dbname=saves port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&save.GameSaveJSON{})
	if err != nil {
		log.Fatal(err)
	}

}
