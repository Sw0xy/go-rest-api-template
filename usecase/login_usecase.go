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
	"golang.org/x/crypto/bcrypt"
)

type loginUseCase struct {
	userRepository repository.UserRepository
	contextTimeout time.Duration
}

func NewLoginUseCase(userRepository repository.UserRepository, timeout time.Duration) domain.LoginUseCase {
	return &loginUseCase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (lu *loginUseCase) Login(ctx context.Context, request domain.LoginRequest) (accessToken string, refreshToken string, err error) {
	var user *models.User
	user, err = lu.userRepository.GetUserByEmail(ctx, request.Email)
	if err != nil {
		log.Error(err)
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
		log.Error("Invalid password")
		err = domain.ErrInvalidPassword
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
