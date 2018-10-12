package main

import (
	"fmt"

	"github.com/apex/log"
	_ "github.com/lib/pq"
	"github.com/t-bonatti/license-manager/config"
	"github.com/t-bonatti/license-manager/datastore"
	"github.com/t-bonatti/license-manager/datastore/database"
	"github.com/t-bonatti/license-manager/server"
)

func main() {
	var cfg = config.Get()

	var db = database.Connect(cfg.DatabaseURL)
	defer func() {
		if err := db.Close(); err != nil {
			log.WithError(err).Error("failed to close database connections")
		}
	}()

	server := server.New(datastore.New(*db), fmt.Sprintf(":%s", cfg.Port))
	if err := server.ListenAndServe(); err != nil {
		log.WithError(err).Fatal("failed to start up server")
	}
}
