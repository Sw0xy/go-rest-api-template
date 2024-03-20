package usecase

import (
	"context"
	"time"

	"github.com/Sw0xy/go-rest-api-template/domain"
	"github.com/Sw0xy/go-rest-api-template/models"
	"github.com/Sw0xy/go-rest-api-template/repository"
)

type userUseCase struct {
	userRepository repository.UserRepository
	contextTimeout time.Duration
}

func NewUserUseCase(userRepository repository.UserRepository, timeout time.Duration) *userUseCase {
	return &userUseCase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (uu *userUseCase) CreateUser(c context.Context, user *models.User) (*domain.UserResponse, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	user, err := uu.userRepository.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	newUser := &domain.UserResponse{
		Id:        user.Id.Hex(),
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	return newUser, nil
}

func (uu *userUseCase) GetUserByEmail(c context.Context, user *models.User) (*domain.UserResponse, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	user, err := uu.userRepository.GetUserByEmail(ctx, user.Email)
	if err != nil {
		return nil, err
	}
	usr := &domain.UserResponse{
		Id:        user.Id.Hex(),
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	return usr, nil
}

func (uu *userUseCase) GetUserById(c context.Context, id string) (*domain.UserResponse, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	user, err := uu.userRepository.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}

	usr := &domain.UserResponse{
		Id:        user.Id.Hex(),
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	return usr, nil
}

func (uu *userUseCase) GetUsers(c context.Context) ([]*domain.UserResponse, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	users, err := uu.userRepository.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	var urs []*domain.UserResponse

	for _, user := range users {
		urs = append(urs, &domain.UserResponse{
			Id:        user.Id.Hex(),
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
		})
	}

	return urs, nil
}

func (uu *userUseCase) DeleteUserById(c context.Context, id string) error {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	err := uu.userRepository.DeleteUserById(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (uu *userUseCase) UpdateUser(c context.Context, user *models.User) error {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	return uu.userRepository.UpdateUser(ctx, user)
}
