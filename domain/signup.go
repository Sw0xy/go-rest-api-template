package domain

import (
	"context"
)

type SignupRequest struct {
	Username string `form:"username" binding:"required"`
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

type SignupResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type SignupUseCase interface {
	SignUp(ctx context.Context, request SignupRequest) (accessToken string, refreshToken string, err error)
}
