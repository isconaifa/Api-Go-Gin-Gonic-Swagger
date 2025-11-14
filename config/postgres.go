package config

import (
	"api/golang/gin-gonic/shemas"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func InitializePostges() (*gorm.DB, error) {
	logger := GetLogger("postgres")

	//Loading environment variables
	_ = godotenv.Load()

	//Building the connection DSN
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}

	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "5432"
	}

	dbname := os.Getenv("DB_DATABASE")
	if dbname == "" {
		dbname = "opening"
	}

	user := os.Getenv("DB_USER")
	if user == "" {
		user = "postgres"
	}

	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		password = "postgres"
	}

	sslmode := os.Getenv("DB_SSLMODE")
	if sslmode == "" {
		sslmode = "disable"
	}

	dbTimezone := os.Getenv("DB_TIMEZONE")
	if dbTimezone == "" {
		dbTimezone = "America/Cuiaba"
	}

	// Construção do DSN (note: sslmode, TimeZone)
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		host, user, password, dbname, port, sslmode, dbTimezone,
	)

	//create DB and connect
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Errorf("Failed to connect to database: %v", err)
	}

	//Migrate shemas
	err = db.AutoMigrate(&shemas.Opening{})
	if err != nil {
		logger.Errorf("Postgres failed to migrate: %v", err)
		return nil, err
	}
	logger.Info("Successfully connected to database")
	return db, err
}
