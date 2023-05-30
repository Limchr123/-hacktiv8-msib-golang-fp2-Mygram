package sosialMedia

import (
	"errors"
	"fmt"
)

type Service interface {
	GetSosmed(userID int) ([]SosialMedia, error)
	CreateSosmed(sosmedInput SosmedInput) (SosialMedia, error)
	UpdateSosmed(getSosmedInput GetSosmedInput, sosmedInput SosmedInput) (SosialMedia, error)
	DeletedSosmed(ID int) (SosialMedia, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetSosmed(userID int) ([]SosialMedia, error) {
	sosmed, err := s.repository.FindByUserId(userID)
	fmt.Println(sosmed)
	if err != nil {
		return sosmed, err
	}
	return sosmed, nil
}

func (s *service) CreateSosmed(input SosmedInput) (SosialMedia, error) {
	sosmed := SosialMedia{}

	sosmed.Name = input.Name
	sosmed.SosialMediaUrl = input.SosialMediaUrl
	sosmed.UserId = input.User.ID

	newSosmed, err := s.repository.Create(sosmed)

	if err != nil {
		return newSosmed, err
	}
	return newSosmed, nil
}

func (s *service) UpdateSosmed(getSosmedInput GetSosmedInput, sosmedInput SosmedInput) (SosialMedia, error) {
	sosmed, err := s.repository.FindById(getSosmedInput.ID)

	if err != nil {
		return sosmed, err
	}
	if sosmed.UserId != sosmedInput.User.ID {
		return sosmed, errors.New("not an owner the account")
	}
	sosmed.Name = sosmedInput.Name
	sosmed.SosialMediaUrl = sosmedInput.SosialMediaUrl

	updatedSosmed, err := s.repository.Update(sosmed)
	if err != nil {
		return updatedSosmed, err
	}
	return updatedSosmed, nil
}

func (s *service) DeletedSosmed(ID int) (SosialMedia, error) {
	sosmed, err := s.repository.FindById(ID)

	if err != nil {
		return sosmed, err
	}

	sosmedDel, err := s.repository.Delete(sosmed)

	if err != nil {
		return sosmedDel, err
	}
	return sosmedDel, nil

}
