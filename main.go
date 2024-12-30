package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type Feed struct {
	Entries []Entry `xml:"entry"`
}

type Entry struct {
	Content string `xml:"content"`
	Link string `xml:"link"`
	Title string `xml:"title"`
	Updated time.Time `xml:"updated"`
	Published time.Time `xml:"published"`
}

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

func main() {
	url := "https://www.reddit.com/r/golang/.rss"
	feedBytes, err := fetchRSSFeed(url)
	if err != nil {
		panic(err)
	}

	feed := parseRSSFeed(feedBytes)
	for _, value := range feed.Entries {
		fmt.Printf("Title: %s\n", value.Title)
		fmt.Printf("Content: %s\n\n", value.Content)
	}
}