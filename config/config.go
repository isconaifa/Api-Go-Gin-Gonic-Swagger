package config

import (
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error
	db, err = InitializePostges()
	if err != nil {
		logger.Errorf("Error to initialize database: %v", err)
	}
	return nil
}

func GetPostgres() *gorm.DB {
	return db
}
func GetLogger(p string) *Logger {
	logger := NewLogger(p)
	return logger
}
