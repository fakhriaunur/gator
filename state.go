package main

import (
	"github.com/fakhriaunur/gator/internal/config"
	"github.com/fakhriaunur/gator/internal/database"
)

type state struct {
	cfg *config.Config
	db  *database.Queries
}

func NewState(cfg *config.Config, db *database.Queries) *state {
	return &state{
		cfg: cfg,
		db:  db,
	}
}
