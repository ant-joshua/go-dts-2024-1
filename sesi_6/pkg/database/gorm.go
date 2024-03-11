package database

import (
	"database/sql"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GormInit(sqlDB *sql.DB) *gorm.DB {
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	// db.AutoMigrate(&models.Product{}, &models.Category{})

	if err != nil {
		log.Fatalf("Error when connect to database: %v", err)
	}

	return db
}
