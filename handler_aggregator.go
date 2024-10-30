package main

import (
	"context"
	"fmt"
)

func handlerRss(s *state, cmd command) error {
	url := "https://www.wagslane.dev/index.xml"
	ctx := context.Background()

	rssFeed, err := fetchFeed(ctx, url)
	if err != nil {
		return err
	}

	fmt.Println(rssFeed)

	return nil
}
