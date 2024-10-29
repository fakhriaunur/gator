package main

import "github.com/fakhriaunur/gator/internal/config"

type state struct {
	cfg *config.Config
}

func NewState(cfg *config.Config) *state {
	return &state{
		cfg: cfg,
	}
}
