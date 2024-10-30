package main

import (
	"context"
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
		return fmt.Errorf("couldn't find user: %w", err)
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

	userParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      username,
	}

	user, err := s.db.CreateUser(ctx, userParams)
	if err != nil {
		return fmt.Errorf("couldn't create user: %w", err)
	}

	if err := s.cfg.SetUser(user.Name); err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Printf("Username: %s has been registered!\n", user.Name)
	printUser(user)

	return nil
}

func handlerReset(s *state, _ command) error {
	ctx := context.Background()

	if err := s.db.DeleteAllUsers(ctx); err != nil {
		return fmt.Errorf("couldn't delete all users: %w", err)
	}

	fmt.Println("successfully reset the user database")
	return nil
}

func printUser(name database.User) {
	fmt.Printf("ID:\t%v\n", name.ID)
	fmt.Printf("Name:\t%v\n", name.Name)
}
