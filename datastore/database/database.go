package database

import (
	"github.com/apex/log"
	"github.com/jmoiron/sqlx"
)

func Connect(url string) *sqlx.DB {
	var log = log.WithField("url", url)
	db, err := sqlx.Open("postgres", url)
	if err != nil {
		log.WithError(err).Fatal("failed to open connection to database")
	}
	if err := db.Ping(); err != nil {
		log.WithError(err).Fatal("failed to ping database")
	}
	return db
}
