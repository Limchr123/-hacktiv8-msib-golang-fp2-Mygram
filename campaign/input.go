package campaign

import "mygram/user"

type PhotoInput struct {
	Title    string `json:"title" binding:"required"`
	Caption  string `json:"caption" binding:"required"`
	PhotoUrl string `json:"photo_url" binding:"required"`
	User     user.User
}

type GetPhotoDetailInput struct {
	ID int `uri:"id" binding:"required"`
}
