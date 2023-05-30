package comment

import (
	"errors"
	"fmt"
)

type Service interface {
	GetComment(userID int, photoId int) ([]Comment, error)
	CreateComment(commentInput CommentInput) (Comment, error)
	UpdateComment(getCommentInput GetCommentInput, commentInput UpdateCommentInput) (Comment, error)
	DeleteComment(ID int) (Comment, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetComment(userID int, photoID int) ([]Comment, error) {
	comment, err := s.repository.FindByUserId(userID, photoID)
	fmt.Println(comment)
	if err != nil {
		return comment, err
	}
	return comment, nil

}

func (s *service) CreateComment(input CommentInput) (Comment, error) {
	comment := Comment{}

	comment.Message = input.Comment
	comment.PhotoId = input.PhotoId
	// comment.PhotoId = input.PhotoID.ID
	comment.UserId = input.User.ID

	newComment, err := s.repository.Create(comment)

	if err != nil {
		return newComment, err
	}
	return newComment, nil
}

func (s *service) UpdateComment(getCommentInput GetCommentInput, commentInput UpdateCommentInput) (Comment, error) {
	comment, err := s.repository.FindById(getCommentInput.ID)

	if err != nil {
		return comment, err
	}
	if comment.UserId != commentInput.User.ID {
		return comment, errors.New("not an owner the account")
	}

	comment.Message = commentInput.Comment

	updatedCampaign, err := s.repository.Update(comment)
	if err != nil {
		return updatedCampaign, err
	}
	return updatedCampaign, nil
}

func (s *service) DeleteComment(ID int) (Comment, error) {
	comment, err := s.repository.FindById(ID)

	if err != nil {
		return comment, err
	}

	commentDel, err := s.repository.Delete(comment)

	if err != nil {
		return commentDel, err
	}
	return commentDel, nil

}
