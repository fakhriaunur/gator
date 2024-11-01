package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/fakhriaunur/gator/internal/database"
	"github.com/google/uuid"
)

// func handlerAggOld(s *state, cmd command) error {
// 	url := "https://www.wagslane.dev/index.xml"
// 	ctx := context.Background()

// 	rssFeed, err := fetchFeed(ctx, url)
// 	if err != nil {
// 		return err
// 	}

// 	fmt.Printf("Feed: %+v\n", rssFeed)

// 	return nil
// }

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("expecting an argument: <time_duration>")
	}

	timeBetweenReqsStr := cmd.args[0]
	timeBetweenReqs, err := time.ParseDuration(timeBetweenReqsStr)
	if err != nil {
		return fmt.Errorf("couldn't parse the time duration: %w", err)
	}

	fmt.Printf("Collecting feeds every %v\n", timeBetweenReqs)

	ticker := time.NewTicker(timeBetweenReqs)
	defer ticker.Stop()

	for ; ; <-ticker.C {
		ctx := context.Background()
		if err := scrapeFeeds(ctx, s); err != nil {
			return fmt.Errorf("couldn't scrape feeds: %w", err)
		}
	}
}

func scrapeFeeds(ctx context.Context, s *state) error {
	nextFeedToFetch, err := s.db.GetNextFeedToFetch(ctx)
	if err != nil {
		return fmt.Errorf("couldn't fetch the next feed, %w", err)
	}
	return scrapeFeed(ctx, s, nextFeedToFetch)
}

func scrapeFeed(ctx context.Context, s *state, feed database.Feed) error {
	_, err := s.db.MarkFeedFetched(ctx, feed.ID)
	if err != nil {
		return fmt.Errorf("couldn't mark the fetched feed, %w", err)
	}

	feedData, err := fetchFeed(ctx, feed.Url)
	if err != nil {
		return fmt.Errorf("couldn't get the feed by url: %w", err)
	}

	for _, item := range feedData.Channel.Item {
		fmt.Printf("Found post: %s\n", item.Title)
	}
	log.Printf("Feed %s collected, %v posts found\n", feed.Name, len(feedData.Channel.Item))

	return nil
}

func handlerFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("expecting <name> <url>")
	}

	ctx := context.Background()
	name := cmd.args[0]
	url := cmd.args[1]

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

	feedFollow, err := s.db.CreateFeedFollow(ctx, feedFollowParams)
	if err != nil {
		return fmt.Errorf("couldn't create a feed follow: %w", err)
	}

	fmt.Println("Feed followed successfully!")
	printFeedFollow(feedFollow.UserName, feedFollow.FeedName)

	return nil
}

func handlerListAllFeeds(s *state, cmd command) error {
	ctx := context.Background()

	feeds, err := s.db.GetAllFeeds(ctx)
	if err != nil {
		return fmt.Errorf("couldn't get all feeds: %w", err)
	}

	if len(feeds) == 0 {
		fmt.Println("No feeds found")
		return nil
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
	fmt.Printf("* User:\t%v\n", user.Name)
}
