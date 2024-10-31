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

	feedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	_, err = s.db.CreateFeedFollow(ctx, feedFollowParams)
	if err != nil {
		return fmt.Errorf("couldn't create a feed follow: %w", err)
	}

	return nil
}

func handlerListAllFeeds(s *state, cmd command) error {
	ctx := context.Background()

	feeds, err := s.db.GetAllFeeds(ctx)
	if err != nil {
		return fmt.Errorf("couldn't get all feeds: %w", err)
	}

	for i, feed := range feeds {
		user, err := s.db.GetUserByID(ctx, feed.UserID)
		if err != nil {
			return err
		}
		fmt.Printf("Feed #%d:\n", i)
		printFeed(feed, user)
	}

	return nil
}

func printFeed(feed database.Feed, user database.User) {
	fmt.Printf("* ID:\t\t%v\n", feed.ID)
	fmt.Printf("* Name:\t\t%v\n", feed.Name)
	fmt.Printf("* URL:\t\t%v\n", feed.Url)
	fmt.Printf("* Creator:\t%v\n", user.Name)
}
