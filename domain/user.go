package domain

type User struct {
	Email     string
	FirstName string
	LastName  string
}

type UserService interface {
	GetUsers() ([]User, error)
}
