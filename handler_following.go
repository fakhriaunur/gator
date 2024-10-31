package main

import (
	"context"
	"fmt"
	"time"

	"github.com/fakhriaunur/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("expecting an argument <url>")
	}

	ctx := context.Background()
	url := cmd.args[0]

	currentUserName := s.cfg.CurrentUserName
	currentUser, err := s.db.GetUser(ctx, currentUserName)
	if err != nil {
		return err
	}

	currentFeed, err := s.db.GetFeedByURL(ctx, url)
	if err != nil {
		return err
	}

	feedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    currentUser.ID,
		FeedID:    currentFeed.ID,
	}

	feedFollow, err := s.db.CreateFeedFollow(ctx, feedFollowParams)
	if err != nil {
		return fmt.Errorf("couldn't create a feed follow: %w", err)
	}

	fmt.Printf("FeedFollow: %+v\n", feedFollow)

	return nil
}

func handlerFollowing(s *state, cmd command) error {

	ctx := context.Background()

	currentUserName := s.cfg.CurrentUserName

	feedFollows, err := s.db.GetFeedFollowsForUser(ctx, currentUserName)
	if err != nil {
		return fmt.Errorf("couldn't get feed follows: %w", err)
	}

	if len(feedFollows) == 0 {
		fmt.Println("No feed follows found")
		return nil
	}

	for i, ff := range feedFollows {
		fmt.Printf("Feed Follows #%d: %v\n", i, ff.FeedName)
	}

	return nil
}

func printFeedFollow(userName, feedName string) {
	fmt.Printf("* User:\t%v\n", userName)
	fmt.Printf("* Feed:\t%v\n", feedName)
}
