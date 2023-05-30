package sosialMedia

import "time"

type SosmedFormatter struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	SosialMediaUrl string    `json:"sosial_media_url"`
	UserID         int       `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
	// UpdatedAt      time.Time     `json:"updated_at"`
	// User           UserFormatter `json:"user"`
}

type UserFormatter struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

func FormatterSosmed(sosmed SosialMedia) SosmedFormatter {
	formatter := SosmedFormatter{
		ID:             sosmed.ID,
		Name:           sosmed.Name,
		SosialMediaUrl: sosmed.SosialMediaUrl,
		UserID:         sosmed.UserId,
		CreatedAt:      sosmed.CreatedAt,
	}
	return formatter
}

type GetSosmedFormatter struct {
	ID             int           `json:"id"`
	Name           string        `json:"name"`
	SosialMediaUrl string        `json:"sosial_media_url"`
	UserID         int           `json:"user_id"`
	CreatedAt      time.Time     `json:"created_at"`
	UpdatedAt      time.Time     `json:"updated_at"`
	User           UserFormatter `json:"user"`
}

func FormatterGet(sosmed SosialMedia) GetSosmedFormatter {
	formatterGet := GetSosmedFormatter{}

	formatterGet.ID = sosmed.ID
	formatterGet.Name = sosmed.Name
	formatterGet.SosialMediaUrl = sosmed.SosialMediaUrl
	formatterGet.UserID = sosmed.UserId
	formatterGet.CreatedAt = sosmed.CreatedAt
	formatterGet.UpdatedAt = sosmed.UpdatedAt

	user := sosmed.User

	sosmedUserFormatter := UserFormatter{}
	sosmedUserFormatter.ID = user.ID
	sosmedUserFormatter.Username = user.Username

	formatterGet.User = sosmedUserFormatter

	return formatterGet
}

func FormatterGetSosmed(sosmed []SosialMedia) []GetSosmedFormatter {
	sosmedGetFormatter := []GetSosmedFormatter{}

	for _, sosmed := range sosmed {
		sosmedFormatter := FormatterGet(sosmed)
		sosmedGetFormatter = append(sosmedGetFormatter, sosmedFormatter)
	}

	return sosmedGetFormatter
}

type UpdatedSosmedFormatter struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	SosialMediaUrl string `json:"sosial_media_url"`
	UserID         int    `json:"user_id"`
	// CreatedAt      time.Time     `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// User           UserFormatter `json:"user"`
}

func FormatterUpdate(sosmed SosialMedia) UpdatedSosmedFormatter {
	formatter := UpdatedSosmedFormatter{
		ID:             sosmed.ID,
		Name:           sosmed.Name,
		SosialMediaUrl: sosmed.SosialMediaUrl,
		UserID:         sosmed.UserId,
		UpdatedAt:      sosmed.UpdatedAt,
	}
	return formatter
}
