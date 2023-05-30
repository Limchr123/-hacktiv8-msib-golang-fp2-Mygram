package user

import "time"

type RegisterFormatter struct {
	Age      int    `json:"age"`
	Email    string `json:"email"`
	ID       int    `json:"id"`
	Username string `json:"username"`
}

func FormatterRegister(user User) RegisterFormatter {
	formatterRegister := RegisterFormatter{
		Age:      user.Age,
		Email:    user.Email,
		ID:       user.ID,
		Username: user.Username,
	}
	return formatterRegister

}

type UserFormatter struct {
	Token string `json:"token"`
}

func FormatterUser(Token string) UserFormatter {
	formatterLogin := UserFormatter{
		Token: Token,
	}
	return formatterLogin
}

type UpdateUserFormatter struct {
	Age       int       `json:"age"`
	Email     string    `json:"email"`
	ID        int       `json:"id"`
	Username  string    `json:"name"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FormatterUpdateUser(user User) UpdateUserFormatter {
	formatterUpdatedUser := UpdateUserFormatter{
		Age:       user.Age,
		Email:     user.Email,
		ID:        user.ID,
		Username:  user.Username,
		UpdatedAt: user.UpdatedAt,
	}
	return formatterUpdatedUser

}

type DeletedUserFormatter struct {
	Message string `json:"message"`
}

func FormatterDeletedUser(user string) DeletedUserFormatter {
	formatterDeletedUser := DeletedUserFormatter{
		Message: user,
	}
	return formatterDeletedUser
}
