package grpc

import (
	"context"

	"github.com/Ardent-7322/go-microservice-playground/internal/user"
	userpb "github.com/Ardent-7322/go-microservice-playground/user-service/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	userpb.UnimplementedUserServiceServer
	svc *user.Service
}

func NewHandler(svc *user.Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) GetUserForOrder(
	ctx context.Context,
	req *userpb.GetUserForOrderRequest,
) (*userpb.GetUserForOrderResponse, error) {

	u, err := h.svc.GetUserForOrder(ctx, req.UserId)
	if err != nil {
		switch err {
		case user.ErrUserNotFound:
			return nil, status.Error(codes.NotFound, err.Error())
		case user.ErrUserInactive:
			return nil, status.Error(codes.FailedPrecondition, err.Error())
		default:
			return nil, status.Error(codes.Internal, "internal error")
		}
	}

	return &userpb.GetUserForOrderResponse{
		UserId:   u.ID,
		Email:    u.Email,
		IsActive: u.IsActive,
	}, nil
}
