package domain

import (
	"context"
	"time"

	"github.com/Sw0xy/go-rest-api-template/models"
)

type UserResponse struct {
	Id        string    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type UserUseCase interface {
	CreateUser(ctx context.Context, user *models.User) (*UserResponse, error)
	GetUsers(ctx context.Context) ([]*UserResponse, error)
	GetUserById(ctx context.Context, id string) (*UserResponse, error)
	DeleteUserById(ctx context.Context, id string) error
	UpdateUser(ctx context.Context, user *models.User) error
}
