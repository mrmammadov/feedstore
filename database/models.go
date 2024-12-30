package database

import "time"

type Feed struct {
	Content string
	Link string
	Title string
	Updated time.Time
	Published time.Time
}