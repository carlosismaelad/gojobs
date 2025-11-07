package config

import (
	"os"

	"github.com/carlosismaelad/gojobs/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")

	dbPath := "./db/main.db"

	// Check if the database file already exists
	logger.Info("Checking if the database file already exists...")
	_, err := os.Stat(dbPath)

	if os.IsNotExist(err) {
		logger.Info("Database file not found. Creating...")
		// Create the database file and directory
		err := os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}


	// Create DB and connect
	db, err := gorm.Open(
		sqlite.Open(dbPath), 
		&gorm.Config{},
	)
	
	if err != nil {
		logger.Errorf("SQLite opening error: %v", err)
		return nil, err
	}

	// Migrate the schema
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("SQLite automigration Error: %v", err)
		return nil, err
	}

	return db, nil	
}