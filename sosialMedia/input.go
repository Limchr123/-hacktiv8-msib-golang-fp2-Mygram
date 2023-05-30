package sosialMedia

import (
	"mygram/user"
)

type SosmedInput struct {
	Name           string `json:"name" binding:"required"`
	SosialMediaUrl string `json:"sosial_media_url" binding:"required"`
	User           user.User
}

type GetSosmedInput struct {
	ID int `uri:"id" binding:"required"`
}
