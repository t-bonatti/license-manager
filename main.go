package main

import (
	"fmt"

	"github.com/apex/log"
	_ "github.com/lib/pq"
	"github.com/t-bonatti/license-manager/config"
	"github.com/t-bonatti/license-manager/datastore/database"
	"github.com/t-bonatti/license-manager/server"
)

func main() {
	var cfg = config.Get()

	database.StartDB(cfg.DatabaseDSN)
	defer func() {
		if err := database.CloseConn(); err != nil {
			log.WithError(err).Error("failed to close database connections")
		}
	}()

	server := server.New()
	server.Run(fmt.Sprintf(":%s", cfg.Port))
}
