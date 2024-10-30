package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/fakhriaunur/gator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("expecting an argument")
	}

	ctx := context.Background()
	username := cmd.args[0]

	_, err := s.db.GetUser(ctx, username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("user does not exist")
		}
		return fmt.Errorf("unable to check existing user: %w", err)
	}

	if err := s.cfg.SetUser(username); err != nil {
		return err
	}

	fmt.Printf("Username: %s has been set\n", s.cfg.CurrentUserName)

	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("expecting an argument")
	}

	ctx := context.Background()
	username := cmd.args[0]

	existingUser, err := s.db.GetUser(ctx, username)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("unable to check existing user: %w", err)
	}
	if existingUser.Name != "" {
		return errors.New("username is already registered")
	}

	userParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      username,
	}

	user, err := s.db.CreateUser(ctx, userParams)
	if err != nil {
		return err
	}

	if err := s.cfg.SetUser(user.Name); err != nil {
		return err
	}

	fmt.Printf("Username: %s has been registered!\n", user.Name)

	return nil
}
