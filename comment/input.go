package comment

import (
	"mygram/campaign"
	"mygram/user"
)

type CommentInput struct {
	Comment string `json:"comment" binding:"required"`
	PhotoId int    `json:"photo_id" binding:"required"`
	PhotoID campaign.Campaign
	User    user.User
}

type UpdateCommentInput struct {
	Comment string `json:"comment" binding:"required"`
	PhotoID campaign.Campaign
	User    user.User
}

type GetCommentInput struct {
	ID int `uri:"id" binding:"required"`
}
