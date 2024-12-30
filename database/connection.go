package database

import (
	"database/sql"
	"fmt"
	"os"
)

type DBConfig struct {
    Host     string
    Port     string
    User     string
    Password string
    DBName   string
} 

func connect(config DBConfig) (*sql.DB, error) {
	connStr := fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        config.Host, config.Port, config.User, config.Password, config.DBName,
    )

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting database: %v", err)
	}

	return db, nil
}

func InsertData(feed Feed) (int64, error) {
	config := DBConfig{
		Host: "localhost",
		Port: "5432",
		User: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName: os.Getenv("DB_NAME"),
	}

	db, err := connect(config)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO feed(content, link, title, updated, published) VALUES ($1, $2, $3, $4, $5)", feed.Content, feed.Link, feed.Title, feed.Updated, feed.Published)
	if err != nil {
		return 0, fmt.Errorf("error occurred during the insertion of the data: %v", err)
	}

	if err != nil {
		panic(err)
	}

	return 1, nil
}