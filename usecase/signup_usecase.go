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
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type signupUseCase struct {
	userRepository repository.UserRepository
	contextTimeout time.Duration
}

func NewSignupUseCase(userRepository repository.UserRepository, timeout time.Duration) domain.SignupUseCase {
	return &signupUseCase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (su *signupUseCase) SignUp(ctx context.Context, request domain.SignupRequest) (accessToken string, refreshToken string, err error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		log.Error(err)
		return
	}

	request.Password = string(encryptedPassword)

	user := &models.User{
		Id:        primitive.NewObjectID(),
		Username:  request.Username,
		Password:  request.Password,
		Email:     request.Email,
		CreatedAt: time.Now(),
	}

	user, err = su.userRepository.CreateUser(ctx, user)
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

	return
}
