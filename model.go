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

func GetUsserFromDatabaseUser(user database.User) User {
	return User{
		Id:  user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name: user.Name,
		ApiKey: user.ApiKey,
	}
}