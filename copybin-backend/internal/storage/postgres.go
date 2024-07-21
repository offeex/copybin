package storage

import (
	"copybin/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func New(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	models := []interface{}{
		&model.User{},
		&model.Pasta{},
		&model.Integration{},
	}
	if err := db.AutoMigrate(models...); err != nil {
		log.Fatalln(err)
	}

	return db
}
