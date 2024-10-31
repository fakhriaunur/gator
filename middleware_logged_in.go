package main

import (
	"context"
	"fmt"

	"github.com/fakhriaunur/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		ctx := context.Background()
		currentUserName := s.cfg.CurrentUserName

		user, err := s.db.GetUser(ctx, currentUserName)
		if err != nil {
			return fmt.Errorf("couldn't get user: %w", err)
		}

		return handler(s, cmd, user)
	}
}
