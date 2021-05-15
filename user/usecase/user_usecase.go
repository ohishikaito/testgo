package usecase

import (
	"app/domain"
	"app/user/infrastructure/repository"
	"context"
)

type UserUsecase interface {
	GetUsers(ctx context.Context) ([]domain.User, error)
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(ur repository.UserRepository) *userUsecase {
	return &userUsecase{
		userRepository: ur,
	}
}

func (u *userUsecase) GetUsers(ctx context.Context) ([]domain.User, error) {
	return u.userRepository.GetUsers(ctx)
}
