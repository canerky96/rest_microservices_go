package services

import (
	"bookstore_users-api/domain/users"
	"bookstore_users-api/utils/crypto_utils"
	"bookstore_users-api/utils/errors"
)

var (
	UsersService usersServiceInterface = &usersService{}
)

type usersService struct {
}

type usersServiceInterface interface {
	GetUser(int64) (*users.User, *errors.RestErr)
	CreateUser(users.User) (*users.User, *errors.RestErr)
	UpdateUser(users.User) (*users.User, *errors.RestErr)
	PartialUpdateUser(users.User) (*users.User, *errors.RestErr)
	DeleteUser(int64) *errors.RestErr
	Search(string) (users.Users, *errors.RestErr)
}

func (s *usersService) GetUser(userId int64) (*users.User, *errors.RestErr) {
	user := &users.User{Id: userId}
	if err := user.Get(); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *usersService) CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	user.Status = users.StatusActive
	user.Password = crypto_utils.GetMd5(user.Password)
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *usersService) UpdateUser(user users.User) (*users.User, *errors.RestErr) {
	current := &users.User{Id: user.Id}
	if err := current.Get(); err != nil {
		return nil, err
	}
	current.FirstName = user.Email
	current.LastName = user.LastName
	current.Email = user.Email
	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

func (s *usersService) PartialUpdateUser(user users.User) (*users.User, *errors.RestErr) {
	current := &users.User{Id: user.Id}
	if err := current.Get(); err != nil {
		return nil, err
	}

	if user.FirstName != "" {
		current.FirstName = user.FirstName
	}
	if user.LastName != "" {
		current.LastName = user.LastName
	}
	if user.Email != "" {
		current.Email = user.Email
	}
	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

func (s *usersService) DeleteUser(id int64) *errors.RestErr {
	user := &users.User{Id: id}
	return user.Delete()
}

func (s *usersService) Search(status string) (users.Users, *errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}
