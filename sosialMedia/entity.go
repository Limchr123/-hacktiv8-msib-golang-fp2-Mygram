package sosialMedia

import (
	"mygram/user"
	"time"
)

type SosialMedia struct {
	ID             int
	Name           string
	SosialMediaUrl string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	UserId         int
	User           user.User `gorm:"foreignKey:UserId"`
}
