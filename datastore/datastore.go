package datastore

import (
	"github.com/t-bonatti/license-manager/datastore/database"
	"github.com/t-bonatti/license-manager/model"
	"gorm.io/gorm"
)

// DataStore interface
type DataStore interface {
	Get(id string, version string) (model.License, error)
	Create(license model.License) error
}

type dataStoreImpl struct {
	db *gorm.DB
}

// New creates a new datastore
func New() DataStore {
	db := database.GetDatabase()
	return dataStoreImpl{db: db}
}

func (ds dataStoreImpl) Get(id string, version string) (license model.License, err error) {
	err = ds.db.Where("id = ? AND version = ?", id, version).First(&license).Error
	return
}

func (ds dataStoreImpl) Create(license model.License) (err error) {
	err = ds.db.Create(license).Error
	return
}
