package main

import (
	"github.com/google/uuid"
	"github.com/segunkayode1/blog-aggregator/internal/database"
	"time"
)

type User struct {
	Id        uuid.UUID    `json:"id"`
	CreatedAt time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey string `json:"api_key"`
}

func GetUserFromDatabaseUser(user database.User) User {
	return User{
		Id:  user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name: user.Name,
		ApiKey: user.ApiKey,
	}
}

type Feed struct {
	Id        uuid.UUID    `json:"id"`
	CreatedAt time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string `json:"name"`
	Url       string `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

func GetFeedFromDatabaseFeed(feed database.Feed) Feed {
	return Feed{
		Id:  feed.ID,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		Name: feed.Name,
		Url: feed.Url,
		UserID: feed.UserID,
	}
}