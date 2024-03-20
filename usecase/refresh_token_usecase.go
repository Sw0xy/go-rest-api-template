package usecase

import (
	"context"
	"os"
	"time"

	"github.com/Sw0xy/go-rest-api-template/domain"
	"github.com/Sw0xy/go-rest-api-template/models"
	"github.com/Sw0xy/go-rest-api-template/repository"
	"github.com/Sw0xy/go-rest-api-template/utils"
	log "github.com/sirupsen/logrus"
)

type refreshTokenUseCase struct {
	userRepository repository.UserRepository
	contextTimeout time.Duration
}

func NewRefreshTokenUseCase(userRepository repository.UserRepository, timeout time.Duration) domain.RefreshTokenUseCase {
	return &refreshTokenUseCase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (rtu *refreshTokenUseCase) RefreshToken(ctx context.Context, request domain.RefreshTokenRequest) (accessToken string, refreshToken string, err error) {
	var id string
	id, err = utils.ExtractIDFromToken(request.RefreshToken, os.Getenv("REFRESH_TOKEN_SECRET"))
	if err != nil {
		log.Error(err)
		return
	}

	var user *models.User
	user, err = rtu.userRepository.GetUserById(ctx, id)
	if err != nil {
		log.Error(err)
		return
	}

	accessToken, err = utils.CreateAccessToken(user, os.Getenv("ACCESS_TOKEN_SECRET"), 1)
	if err != nil {
		log.Error(err)
		return
	}

	refreshToken, err = utils.CreateRefreshToken(user, os.Getenv("REFRESH_TOKEN_SECRET"), 1)
	if err != nil {
		log.Error(err)
		return
	}

	return accessToken, refreshToken, nil
}
