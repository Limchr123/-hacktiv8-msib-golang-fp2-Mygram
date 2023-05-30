package user

type RegisterUserInput struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Age      int    `json:"age" binding:"required"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type DeletedUser struct {
	ID int `uri:"id" binding:"required"`
}

type UpdatedUser struct {
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	User     User
}
