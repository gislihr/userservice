package service

import (
	"context"

	"github.com/gislihr/userservice"
	"github.com/gislihr/userservice/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	store userservice.Store
}

func New(store userservice.Store) *Service {
	return &Service{store: store}
}

func (s *Service) AddUser(ctx context.Context, request *proto.AddUserRequest) (*proto.UserResponse, error) {
	input := userservice.UserInput{
		Name:           request.UserName,
		Email:          request.Email,
		HashedPassword: "hashed pass",
	}

	if err := input.Valid(); err != nil {
		return nil, status.Error(codes.FailedPrecondition, err.Error())
	}

	user, err := s.store.AddUser(input)

	if err != nil {
		return nil, err
	}

	return &proto.UserResponse{
		User: &proto.User{
			Id:    user.Id,
			Name:  user.Name,
			Email: user.Email,
		},
	}, nil
}
