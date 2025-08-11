package main

import (
	"fmt"
	"time"

	"github.com/mentalcaries/go-gator/internal/database"
)

func handlerAgg(s *state, cmd command, currentUser database.User) error {
	
    if len(cmd.args) < 1 {
        return  fmt.Errorf("please specificy time between requests")
    }

    timeBetweenRequests, err := time.ParseDuration(cmd.args[0])
    if err != nil {
        return fmt.Errorf("invalid duration: %v", err)
    }

    ticker := time.NewTicker(timeBetweenRequests)
    for ; ; <-ticker.C {
        fmt.Printf("Collecting feeds every %v ...\n", timeBetweenRequests)
        scrapeFeeds(s)
        fmt.Println()
        fmt.Println("Sit tight. We'll check for new posts in a bit...")
        fmt.Println()

    }
}
