package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"time"

	
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

const feedURL = "https://www.wagslane.dev/index.xml"

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

func handleFetchFeed(s *state, cmd command) error {
	articles, err := fetchFeed(context.Background(), feedURL)

	if err != nil {
		return fmt.Errorf("could not get articles: %v", err)
	}

	fmt.Println(articles)
	return nil
}


