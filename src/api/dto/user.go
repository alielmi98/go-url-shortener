package dto

type RegisterUserByUsernameRequest struct {
	Username string `json:"username" binding:"required,min=5"`
	Email    string `json:"email" binding:"min=6,email"`
	Password string `json:"password" binding:"required,min=6"`
}
