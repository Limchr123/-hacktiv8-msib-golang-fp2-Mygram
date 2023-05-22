package campaign

import (
	"mygram/user"
	"time"
)

type Campaign struct {
	ID        int
	Title     string
	Caption   string
	PhotoUrl  string
	UserId    int
	CreatedAt time.Time
	UpdatedAt time.Time
	User      user.User
}
