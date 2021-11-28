package service

import (
	"context"
	"errors"

	"github.com/gislihr/userservice"
	"github.com/gislihr/userservice/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	store           userservice.Store
	passwordManager userservice.PasswordManager
}

func New(store userservice.Store, passwordManager userservice.PasswordManager) *Service {
	return &Service{store: store, passwordManager: passwordManager}
}

func (s *Service) AddUser(ctx context.Context, request *proto.AddUserRequest) (*proto.UserResponse, error) {
	hashedPassword, err := s.passwordManager.HashPassword(request.Password)
	if err != nil {
		return nil, status.Error(codes.Internal, "error adding user")
	}
	input := userservice.UserInput{
		UserName:       request.UserName,
		Name:           request.Name,
		Email:          request.Email,
		HashedPassword: hashedPassword,
	}

	if err := input.Valid(); err != nil {
		return nil, status.Error(codes.FailedPrecondition, err.Error())
	}

	user, err := s.store.AddUser(input)

	if err != nil {
		if errors.Is(err, userservice.ErrorInvalidUserName) || errors.Is(err, userservice.ErrorInvalidEmail) {
			return nil, status.Error(codes.FailedPrecondition, err.Error())
		}
		logrus.WithError(err).Error("error adding user")
		return nil, status.Error(codes.Internal, "error adding user")
	}

	return &proto.UserResponse{
		User: &proto.User{
			Id:    user.Id,
			Name:  user.Name,
			Email: user.Email,
		},
	}, nil
}

func (s *Service) Login(ctx context.Context, request *proto.LoginRequest) (*proto.AuthenticationResponse, error) {
	user, err := s.store.GetUserByEmailOrUsername(request.UserNameOrEmail)
	if err != nil {
		return nil, err
	}

	err = s.passwordManager.CompareHashAndPassword(user.HashedPassword, request.Password)
	if err != nil {
		return nil, status.Error(codes.FailedPrecondition, "wrong username/email or password")
	}

	return &proto.AuthenticationResponse{
		JwtToken: "jwt token is supposed to be here",
		User: &proto.User{
			Id:       user.Id,
			Name:     user.Name,
			UserName: user.UserName,
			Email:    user.Email,
		},
	}, nil
}
