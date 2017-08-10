package datastore

import (
	"github.com/jmoiron/sqlx"
	"github.com/t-bonatti/license-manager/model"
)

type DataStore interface {
	Get(id string, version string) (license model.License, err error)
	Create(license model.License) (err error)
}

type dataStore struct {
	db sqlx.DB
}

// New creates a new datastore
func New(db sqlx.DB) DataStore {
	return dataStore{db: db}
}

func (ds dataStore) Get(id string, version string) (license model.License, err error) {
	return license, ds.db.Get(&license, "SELECT * FROM licenses WHERE id = $1 and version = $2", id, version)
}

func (ds dataStore) Create(license model.License) (err error) {
	_, err = ds.db.Exec(
		"INSERT INTO licenses(id, version, created_at, info) VALUES($1, $2, $3, $4)",
		license.ID,
		license.Version,
		license.CreatedAt,
		license.Info,
	)
	return
}
