package database

import (
	"time"

	"github.com/apex/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// StartDB initiate connection with postgres
func StartDB(dsn string) {
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.WithError(err).Fatal("Could not connect to the Postgres Database")
	}

	db = database
	config, _ := db.DB()
	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)

	// TODO Implements migrations: migrations.RunMigrations(db)
}

// CloseConn close connection with postgres
func CloseConn() error {
	config, err := db.DB()
	if err != nil {
		return err
	}

	err = config.Close()
	if err != nil {
		return err
	}

	return nil
}

// GetDatabase return database connection
func GetDatabase() *gorm.DB {
	return db
}
