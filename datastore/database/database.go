package database

import (
	"time"

	"github.com/apex/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

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

func GetDatabase() *gorm.DB {
	return db
}
