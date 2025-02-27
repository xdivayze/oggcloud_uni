package db

import (
	"fmt"
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Create_DB() error {
	dsn := os.Getenv("POSTGRES_URI")
	fmt.Println(dsn)
	var err error
	DB, err = setupDatabase(dsn)
	if err != nil {
		return fmt.Errorf("error while creating db:\n\t%w", err)
	}

	return nil
}

func setupDatabase(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to db %w", err)
	}

	return db, nil

}
