package campaign

import (
	"errors"
	"fmt"
)

type Service interface {
	CreateImage(input PhotoInput) (Campaign, error)
	GetCampaigns(userID int) ([]Campaign, error)
	UpdateCampaigns(inputID GetPhotoDetailInput, input PhotoInput) (Campaign, error)
	DeletePhoto(ID int) (Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateImage(input PhotoInput) (Campaign, error) {
	photo := Campaign{}

	photo.Title = input.Title
	photo.Caption = input.Caption
	photo.PhotoUrl = input.PhotoUrl
	photo.UserId = input.User.ID

	newPhoto, err := s.repository.CreateImage(photo)
	if err != nil {
		return newPhoto, err
	}
	return newPhoto, nil
}

func (s *service) DeletePhoto(ID int) (Campaign, error) {
	Photo, err := s.repository.FindById(ID)
	if err != nil {
		return Photo, err
	}
	photoDel, err := s.repository.Delete(Photo)

	if err != nil {
		return photoDel, err
	}
	return photoDel, nil
}

// func (s *service) SaveAvatar(input PhotoInput, ID int, fileLocation string) (Photo, error) {
// 	user := Photo{}

// 	user.Title = input.Title
// 	user.Caption = input.Caption

// 	user, err := s.repository.FindById(ID)

// 	if err != nil {
// 		return user, err
// 	}

// 	user.PhotoUrl = fileLocation

// 	newPhoto, err := s.repository.CreateImage(user)
// 	if err != nil {
// 		return newPhoto, err
// 	}
// 	return newPhoto, nil

// }

func (s *service) GetCampaigns(userID int) ([]Campaign, error) {
	fmt.Println(userID)
	if userID != 0 {
		photo, err := s.repository.FindByUserId(userID)
		if err != nil {
			return photo, err
		}
		return photo, nil
	}

	photo, err := s.repository.FindAll()
	if err != nil {
		return photo, err
	}
	return photo, nil
}

func (s *service) UpdateCampaigns(inputID GetPhotoDetailInput, input PhotoInput) (Campaign, error) {
	campaign, err := s.repository.FindById(inputID.ID)
	if err != nil {
		return campaign, err
	}

	if campaign.UserId != input.User.ID {
		return campaign, errors.New("not an owner of the campaign")
	}

	campaign.Title = input.Title
	campaign.Caption = input.Caption
	campaign.PhotoUrl = input.PhotoUrl

	updatedCampaign, err := s.repository.Update(campaign)
	if err != nil {
		return updatedCampaign, err
	}
	return updatedCampaign, nil
}
