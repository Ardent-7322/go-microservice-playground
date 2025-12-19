package user

import (
	"context"
	"errors"

	"github.com/go-microservice-playground/user-service/internal/domain"
	"github.com/go-microservice-playground/user-service/internal/repository"
)

var (
	ErrUserNotFound = errors.New("user not found")
	ErrUserInactive = errors.New("user inactive")
)

type Service struct {
	repo repository.UserRepository
}

func NewService(r repository.UserRepository) *Service {
	return &Service{repo: r}
}

func (s *Service) GetUserForOrder(
	ctx context.Context,
	userID uint64,
) (*domain.User, error) {

	user, err := s.repo.FindByID(ctx, userID)
	if err != nil {
		return nil, ErrUserNotFound
	}

	if !user.IsActive {
		return nil, ErrUserInactive
	}

	return user, nil
}
