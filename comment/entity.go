package comment

import (
	"mygram/campaign"
	"mygram/user"
	"time"
)

type Comment struct {
	ID        int
	UserId    int
	PhotoId   int
	Message   string
	CreatedAt time.Time
	UpdatedAt time.Time
	User      user.User         `gorm:"foreignKey:UserId"`
	Campaign  campaign.Campaign `gorm:"foreignKey:PhotoId;preload"`
}
