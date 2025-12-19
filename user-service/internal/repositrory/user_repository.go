package repositrory

import (
	"context"
	"github.com/Ardent-7322/go-microservice-playground/user-service/intenal/domain"
)

type UserRepository interface {
	FindByID(ctx context.Context, id uint64) (*domain.User, error)
}
