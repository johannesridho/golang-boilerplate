package grpc

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"golang-boilerplate/domain"
	userPB "golang-boilerplate/protobuf-files/user"
)

type userServer struct {
	svc domain.UserService
}

func NewUserServer(svc domain.UserService) *userServer {
	return &userServer{svc: svc}
}

func (s *userServer) GetUsers(ctx context.Context, req *userPB.GetUsersRequest) (*userPB.GetUsersResponse, error) {
	users, err := s.svc.GetUsers()
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprint("internal error"))
	}

	pbUsers := make([]*userPB.User, len(users))

	for i, user := range users {
		pbUsers[i] = &userPB.User{
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		}
	}

	return &userPB.GetUsersResponse{Users: pbUsers}, nil
}
