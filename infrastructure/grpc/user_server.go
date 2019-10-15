package grpc

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"golang-ddd-boilerplate/domain"
	userPb "golang-ddd-boilerplate/protobuf-files/user"
)

type userServer struct {
	svc domain.UserService
}

func NewUserServer(svc domain.UserService) *userServer {
	return &userServer{svc: svc}
}

func (s *userServer) GetUsers(ctx context.Context, req *userPb.GetUsersRequest) (*userPb.GetUsersResponse, error) {
	users, err := s.svc.GetUsers()
	if err != nil {
		fmt.Errorf("error %v", err)
		return nil, status.Error(codes.Internal, fmt.Sprint("internal error"))
	}

	pbUsers := make([]*userPb.User, len(users))

	for i, user := range users {
		pbUsers[i] = &userPb.User{
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		}
	}

	return &userPb.GetUsersResponse{Users: pbUsers}, nil
}
