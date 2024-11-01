package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/fakhriaunur/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.args) == 1 {
		if specifiedLimit, err := strconv.Atoi(cmd.args[0]); err == nil {
			limit = specifiedLimit
		} else {
			return fmt.Errorf("invalid limit: %w", err)
		}
	} else if len(cmd.args) > 1 {
		return fmt.Errorf("expecting an optional argument: <limit>")
	}

	ctx := context.Background()

	params := database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	}
	posts, err := s.db.GetPostsForUser(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't get posts : %w", err)
	}

	fmt.Printf("Found %d posts for user %s\n", len(posts), user.Name)
	for i, post := range posts {
		fmt.Printf("Post #%d:\n", i+1)
		fmt.Printf("%s from %s\n", post.PublishedAt.Time.Format("Mon Jan 2"), post.FeedName)
		fmt.Printf(" --- %s --- \n", post.Title)
		fmt.Printf("\t%s\n", post.Description.String)
		fmt.Printf("Link: %s\n", post.Url)
		fmt.Println("==================================")
	}

	return nil
}
