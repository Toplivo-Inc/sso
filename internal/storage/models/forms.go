package models

// UserRegisterForm is sent to /register handler
type UserRegisterForm struct {
	Username string `json:"username" example:"John" binding:"required"`
	Email    string `json:"email" example:"john@gmail.com" binding:"required,email"`
	Password string `json:"password" example:"not12345" binding:"required"`
}

// UserLoginForm is sent to /login handler
//
// Login could be either username or email
type UserLoginForm struct {
	Login    string `json:"login" example:"John" binding:"required"`
	Password string `json:"password" example:"not12345" binding:"required"`
}

// LoginMetadata is sent to /login handler
type LoginMetadata struct {
	UserAgent string
	IP string
}
