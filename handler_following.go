package main

import (
	"context"
	"fmt"
	"time"

	"github.com/fakhriaunur/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("expecting an argument <url>")
	}

	ctx := context.Background()
	url := cmd.args[0]

	currentFeed, err := s.db.GetFeedByURL(ctx, url)
	if err != nil {
		return err
	}

	feedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    currentFeed.ID,
	}

	feedFollow, err := s.db.CreateFeedFollow(ctx, feedFollowParams)
	if err != nil {
		return fmt.Errorf("couldn't create a feed follow: %w", err)
	}

	fmt.Printf("FeedFollow: %+v\n", feedFollow)

	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {

	ctx := context.Background()

	feedFollows, err := s.db.GetFeedFollowsForUser(ctx, user.Name)
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

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("expecting an argument <url>")
	}

	ctx := context.Background()
	url := cmd.args[0]

	feed, err := s.db.GetFeedByURL(ctx, url)
	if err != nil {
		return fmt.Errorf("couldn't get the feed: %w", err)
	}

	if err := s.db.DeleteFeedFollowByFeedURL(ctx, database.DeleteFeedFollowByFeedURLParams{
		Url:    url,
		UserID: user.ID,
	}); err != nil {
		return fmt.Errorf("couldn't delete the feed follow: %w", err)
	}
	fmt.Printf("%v was successfully unfollowed!", feed.Name)

	return nil
}

func printFeedFollow(userName, feedName string) {
	fmt.Printf("* User:\t%v\n", userName)
	fmt.Printf("* Feed:\t%v\n", feedName)
}
