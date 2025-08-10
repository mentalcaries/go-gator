package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mentalcaries/go-gator/internal/database"
)

func handleFollowFeed(s *state, cmd command, currentUser database.User) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("invalid argument")
	}
	url := cmd.args[0]

	existingFeedId, err := s.db.GetFeedIdByUrl(context.Background(), url)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("feed does not exist")
		}
		return err
	}

	newFeedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    currentUser.ID,
		FeedID:    existingFeedId,
	})

	if err != nil {
		return fmt.Errorf("could not follow feed", err)
	}

	fmt.Println("Feed followed:", newFeedFollow.FeedName)
	fmt.Println("User:", newFeedFollow.UserName)

	return nil
}

func handlerListFeedFollows(s *state, cmd command, currentUser database.User) error {

	followedFeeds, err := s.db.GetFeedFollowsForUser(context.Background(), currentUser.ID)
	if err != nil {
		return fmt.Errorf("failed to get feeds: %v", err)
	}

	if len(followedFeeds) == 0 {
		fmt.Println("No followed feeds for this user.")
		return nil
	}

	fmt.Println("Followed Feeds:")
	for _, feed := range followedFeeds {
		fmt.Println("* ", feed.Name)
	}
	return nil
}

func handlerUnfollowFeed(s *state, cmd command, currentUser database.User) error {

    if len(cmd.args) < 1 {
        return fmt.Errorf("invalid arguments")
    }

    url := cmd.args[0]
    _, err := s.db.RemoveFeedFollow(context.Background(), database.RemoveFeedFollowParams{
        UserID: currentUser.ID,
        Url: url,
    })

    if err != nil {
        if errors.Is(err, sql.ErrNoRows){
            fmt.Println("You don't follow this feed")
            return nil
        }
        return  fmt.Errorf("error unfollowing feed: %v", err)
    }
    fmt.Println("You have unfollowed", url)

    return nil
}