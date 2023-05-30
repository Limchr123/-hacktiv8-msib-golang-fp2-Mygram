package comment

import (
	"time"
)

type CommentFormatter struct {
	ID        int       `json:"id"`
	Message   string    `json:"message"`
	UserID    int       `json:"user_id"`
	PhotoID   int       `json:"photo_id"`
	CreatedAt time.Time `json:"created_at"`
	// PhotoUrl  string    `json:"photo_url"`
	// UpdatedAt time.Time `json:"updated_at"`
}

func FormatterComment(comment Comment) CommentFormatter {
	formatter := CommentFormatter{
		ID:        comment.ID,
		Message:   comment.Message,
		UserID:    comment.UserId,
		PhotoID:   comment.PhotoId,
		CreatedAt: comment.CreatedAt,
	}
	return formatter
}

type GetCommentFormatter struct {
	ID        int                      `json:"id"`
	Message   string                   `json:"message"`
	UserID    int                      `json:"user_id"`
	PhotoID   int                      `json:"photo_id"`
	CreatedAt time.Time                `json:"created_at"`
	User      CommentUserFormatter     `json:"user"`
	Campaign  CommentCampaignFormatter `json:"photo"`
}

type CommentCampaignFormatter struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
	UserID   int    `json:"user_id"`
}

type CommentUserFormatter struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	UserName string `json:"username"`
}

func FormatterGet(comment Comment) GetCommentFormatter {
	formatterGet := GetCommentFormatter{}
	formatterGet.ID = comment.ID
	formatterGet.Message = comment.Message
	formatterGet.UserID = comment.UserId
	formatterGet.PhotoID = comment.PhotoId
	formatterGet.CreatedAt = comment.CreatedAt

	user := comment.User

	commentUserFormatter := CommentUserFormatter{}
	commentUserFormatter.ID = user.ID
	commentUserFormatter.Email = user.Email
	commentUserFormatter.UserName = user.Username

	formatterGet.User = commentUserFormatter

	campaign := comment.Campaign

	commentCampaignFormatter := CommentCampaignFormatter{}
	commentCampaignFormatter.ID = campaign.ID
	commentCampaignFormatter.Title = campaign.Title
	commentCampaignFormatter.Caption = campaign.Caption
	commentCampaignFormatter.PhotoUrl = campaign.PhotoUrl
	commentCampaignFormatter.UserID = campaign.UserId

	formatterGet.Campaign = commentCampaignFormatter

	return formatterGet
}

func FormatterGetComment(comments []Comment) []GetCommentFormatter {
	commentGetFormatter := []GetCommentFormatter{}

	for _, comment := range comments {
		commentFormatter := FormatterGet(comment)
		commentGetFormatter = append(commentGetFormatter, commentFormatter)
	}

	return commentGetFormatter
}

type UpdatedFormatter struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
}

func FormatterUpdated(comment Comment) UpdatedFormatter {
	formatterUpdated := UpdatedFormatter{
		ID:      comment.ID,
		Message: comment.Message,
	}
	return formatterUpdated
}
