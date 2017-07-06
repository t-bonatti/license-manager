package model

import "time"

type License struct {
	ID        string            `json:"id"`
	Version   string            `json:"version"`
	CreatedAt time.Time         `json:"createdAt"`
	Info      map[string]string `json:"info"`
}
