package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(input RegisterUserInput) (User, error)
	Login(input LoginUserInput) (User, error)
	GetUserById(userId int) (User, error)
	// GetProfile(ID int) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Register(input RegisterUserInput) (User, error) {
	user := User{}
	user.Email = input.Email
	user.Username = input.Username

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	if err != nil {
		return user, err
	}

	user.Password = string(passwordHash)

	newUser, err := s.repository.Save(user)

	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) Login(input LoginUserInput) (User, error) {
	user, err := s.repository.FindByEmail(input.Email)

	if err != nil {
		return user, errors.New("user with that email doesn't exist")
	}

	return user, nil
}

func (s *service) GetUserById(userId int) (User, error) {
	user, err := s.repository.GetById(userId)

	if err != nil {
		return user, errors.New("user with that id doesn't exist")
	}

	return user, nil
}
