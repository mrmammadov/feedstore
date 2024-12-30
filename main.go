package main

import (
	"encoding/xml"
	"feedstore/database"
	"fmt"
	"io"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func fetchRSSFeed(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Problem occured reading the response: %s", err)
		return nil, err
	}
	return bytes, nil
}

func parseRSSFeed(feedXML []byte) Feed {
	var feed Feed
	err := xml.Unmarshal(feedXML, &feed)
	if err != nil {
		fmt.Printf("Problem occured parsing the XML:  %s", err)
		os.Exit(1)
	}
	return feed
}

func (f Feed) ToDBEntry() []database.Feed {
	dbEntries := make([]database.Feed, len(f.Entries))
	for i, e := range f.Entries {
		db_feed := database.Feed{
			Content:   e.Content,
			Link:      e.Link,
			Title:     e.Title,
			Updated:   e.Updated,
			Published: e.Published,
		}
		dbEntries[i] = db_feed
	}
	return dbEntries
}

func main() {
	url := "https://www.reddit.com/r/golang/.rss"
	feedBytes, err := fetchRSSFeed(url)
	if err != nil {
		panic(err)
	}

	feed := parseRSSFeed(feedBytes)
	db_feed := feed.ToDBEntry()

	for _, e := range db_feed{
		_, err = database.InsertData(e)
		if err != nil {
			panic(err)
		}
	}

}