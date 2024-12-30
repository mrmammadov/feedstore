package main

import "time"

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
