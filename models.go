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
	LastFetchedAt * time.Time  `json:"last_fetched_at"`
}

func GetFeedFromDatabaseFeed(feed database.Feed) Feed {
	return Feed{
		Id:  feed.ID,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		Name: feed.Name,
		Url: feed.Url,
		UserID: feed.UserID,
		LastFetchedAt: &feed.LastFetchedAt.Time,
	}
}


type UsersFeed struct {
	Id        uuid.UUID    `json:"id"`
	CreatedAt time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
}

func GetUsersFeedsFromDatabaseUsersFeeds(userfeed database.UsersFeed) UsersFeed {
	return UsersFeed{
		Id:  userfeed.ID,
		CreatedAt: userfeed.CreatedAt,
		UpdatedAt: userfeed.UpdatedAt,
		UserID: userfeed.UserID,
		FeedID: userfeed.FeedID,
	}
}

type Post struct {
	ID          uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Title       string
	Url         string
	Description string
	PublishedAt string
	FeedID      uuid.UUID
}

func GetPostsFromDatabasePosts(post database.Post) Post {
	return Post{
		ID:  post.ID,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
		Title:  post.Title,
		Url: post.Url,
		Description: post.Description,
		PublishedAt: post.PublishedAt,
		FeedID: post.FeedID,
	}
}