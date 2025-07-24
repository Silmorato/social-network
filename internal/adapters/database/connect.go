package database

import (
	"fmt"
	"github.com/magiconair/properties"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func InitDB() (*gorm.DB, error) {
	p, err := properties.LoadFile("config/render.properties", properties.UTF8)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	host := p.GetString("DB_HOST", "")
	port := p.GetString("DB_PORT", "")
	user := p.GetString("DB_USER", "")
	password := p.GetString("DB_PASSWORD", "")
	dbName := p.GetString("DB_NAME", "")

	sslmode := p.GetString("DB_SSLMODE", "disable")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host, user, password, dbName, port, sslmode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to DB: %w", err)
	}

	log.Println("database connected successfully")
	return db, nil
}
