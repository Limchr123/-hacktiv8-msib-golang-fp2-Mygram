package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	GetUserByid(ID int) (User, error)
	DeleteUser(ID int) (User, error)
	UpdatedUser(getUpdatedInput DeletedUser, inputUser UpdatedUser) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) UpdatedUser(getUpdatedInput DeletedUser, inputUser UpdatedUser) (User, error) {

	user, err := s.repository.FindById(getUpdatedInput.ID)
	if err != nil {
		return user, err
	}

	if user.ID != inputUser.User.ID {
		return user, errors.New("not an owner the account")
	}

	user.Email = inputUser.Email
	user.Username = inputUser.Username

	userUpdated, err := s.repository.Update(user)
	if err != nil {
		return userUpdated, err
	}

	return userUpdated, nil

}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}

	user.Username = input.Username
	user.Email = input.Email
	user.Password = input.Password
	user.Age = input.Age
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

func (s *service) Login(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}
	if user.ID == 0 {
		return user, errors.New("User not found that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}
	return user, nil

}

func (s *service) DeleteUser(userID int) (User, error) {
	user, err := s.repository.FindById(userID)
	if err != nil {
		return user, err
	}
	userDel, err := s.repository.Delete(user)

	if err != nil {
		return userDel, err
	}
	return userDel, nil
}

func (s *service) GetUserByid(ID int) (User, error) {
	user, err := s.repository.FindById(ID)

	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("User Not Found With That ID")
	}

	return user, nil

}
