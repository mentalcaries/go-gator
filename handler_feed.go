package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mentalcaries/go-gator/internal/database"
)

func handleAddFeed(s *state, cmd command) error {
	if len(cmd.args) < 2 {
		return fmt.Errorf("invalid args")
	}
	name := cmd.args[0]
	url := cmd.args[1]

	currentUser, err := s.db.GetUserByName(context.Background(), s.config.CurrentUserName)
	if err != nil {
		return fmt.Errorf("failed to get current user: %v", err)
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		Name:      name,
		Url:       url,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    currentUser.ID,
	})

	if err != nil {
		return fmt.Errorf("could not save feed: %v", err)
	}

	fmt.Println("Feed created successfully:")
	printFeed(feed)
	fmt.Println()
	fmt.Println("=====================================")

	return nil
}

func handleAllFeeds(s *state, cmd command) error {

	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("could not fetch feeds: %v", err)
	}

	for _, feed := range feeds {
		fmt.Println("*Name:", feed.Name)
		fmt.Println("*Url:", feed.Url)
		fmt.Println("*User:", feed.UserName)
		fmt.Println()
		fmt.Println("=====================================")

	}
	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf("* ID:            %s\n", feed.ID)
	fmt.Printf("* Created:       %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:       %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* URL:           %s\n", feed.Url)
	fmt.Printf("* UserID:        %s\n", feed.UserID)
}
