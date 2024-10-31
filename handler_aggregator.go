package main

import (
	"context"
	"fmt"
	"time"

	"github.com/fakhriaunur/gator/internal/database"
	"github.com/google/uuid"
)

func handlerRss(s *state, cmd command) error {
	url := "https://www.wagslane.dev/index.xml"
	ctx := context.Background()

	rssFeed, err := fetchFeed(ctx, url)
	if err != nil {
		return err
	}

	fmt.Printf("Feed: %+v\n", rssFeed)

	return nil
}

func handlerFeed(s *state, cmd command) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("expecting <name> <url>")
	}

	ctx := context.Background()
	name := cmd.args[0]
	url := cmd.args[1]

	userName := s.cfg.CurrentUserName
	user, err := s.db.GetUser(ctx, userName)
	if err != nil {
		return err
	}

	feedParams := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	}

	feed, err := s.db.CreateFeed(ctx, feedParams)
	if err != nil {
		return fmt.Errorf("couldn't create a feed: %w", err)
	}

	fmt.Printf("Feed: %+v\n", feed)

	return nil
}

func handlerListAllFeeds(s *state, cmd command) error {
	ctx := context.Background()

	feeds, err := s.db.GetAllFeeds(ctx)
	if err != nil {
		return fmt.Errorf("couldn't get all feeds: %w", err)
	}

	for i, feed := range feeds {
		fmt.Printf("Feed #%d: %+v\n", i, feed)
	}

	return nil
}
