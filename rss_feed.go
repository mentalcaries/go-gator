package main

import (
	"context"
	"database/sql"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"time"

	"github.com/mentalcaries/go-gator/internal/database"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

// const feedURL = "https://www.wagslane.dev/index.xml"

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return &RSSFeed{}, err
	}
	req.Header.Set("User-Agent", "gator")

	res, err := client.Do(req)
	if err != nil {
		return &RSSFeed{}, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return &RSSFeed{}, err
	}

	var rssData RSSFeed

	xml.Unmarshal(body, &rssData)

	var formattedRss RSSFeed

	formattedRss.Channel.Title = html.UnescapeString(rssData.Channel.Title)
	formattedRss.Channel.Description = html.UnescapeString(rssData.Channel.Description)
	formattedRss.Channel.Link = rssData.Channel.Link

	for _, item := range rssData.Channel.Item {
		formattedRss.Channel.Item = append(formattedRss.Channel.Item, RSSItem{
			Title:       html.UnescapeString(item.Title),
			Link:        item.Link,
			Description: html.UnescapeString(item.Description),
			PubDate:     html.UnescapeString(item.PubDate),
		})
	}

	return &formattedRss, nil
}


func scrapeFeeds(s *state) error{
    feedToFetch, err := s.db.GetNextFeedToFetch(context.Background())
    if err != nil {
        return fmt.Errorf("could not fetch feed: \n%v", err)
    }

    updatedFeed, err := s.db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
        ID: feedToFetch.ID,
        LastFetchedAt: sql.NullTime{Time: time.Now().UTC(), Valid: true},
        UpdatedAt: time.Now().UTC(),
    })

    if err != nil {
        return fmt.Errorf("could not update last fetched at")
    }

    nextFeed, err := fetchFeed(context.Background(), updatedFeed.Url)

    if err != nil {
        return  fmt.Errorf("could not fetch next feed")
    }

    fmt.Printf("Posts from %v:\n", nextFeed.Channel.Title)
    for _, post := range nextFeed.Channel.Item {
        fmt.Println()
        fmt.Println(post.Title)
        fmt.Println("===========")
    }
    
    return nil
}