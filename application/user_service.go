package application

import "golang-boilerplate/domain"

type userService struct{}

func NewUserService() domain.UserService {
	return &userService{}
}

func (svc *userService) GetUsers() ([]domain.User, error) {
	return []domain.User{
		{Email: "user1@email.com", FirstName: "first name 1", LastName: "last name 1"},
		{Email: "user2@email.com", FirstName: "first name 2", LastName: "last name 2"},
	}, nil
}
