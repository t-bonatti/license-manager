package datastore

import (
	"github.com/jmoiron/sqlx"
	"github.com/t-bonatti/license-manager/model"
)

type DataStore struct {
	db sqlx.DB
}

func New(db sqlx.DB) DataStore {
	return DataStore{db: db}
}

func (ds *DataStore) Get(id string, version string) (license model.License, err error) {
	return license, ds.db.Get(&license, "SELECT * FROM licenses WHERE id = $1 and license = $2", id, version)
}

func (ds *DataStore) Create(license model.License) (err error) {
	_, err = ds.db.Exec(
		"INSERT INTO licenses(id, version, created_at, info) VALUES($1, $2, $3, $4)",
		license.ID,
		license.Version,
		license.CreatedAt,
		license.Info,
	)
	return
}
