package model

import (
	"time"

	"github.com/jmoiron/sqlx/types"
)

// License struct
type License struct {
	ID        string         `json:"id"`
	Version   string         `json:"version"`
	CreatedAt time.Time      `json:"createdAt" db:"created_at"`
	Info      types.JSONText `json:"info"`
}
