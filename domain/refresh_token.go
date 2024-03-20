package domain

import (
	"context"
)

type RefreshTokenRequest struct {
	RefreshToken string `form:"refreshToken" binding:"required"`
}

type RefreshTokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type RefreshTokenUseCase interface {
	RefreshToken(ctx context.Context, request RefreshTokenRequest) (accessToken string, refreshToken string, err error)
}
