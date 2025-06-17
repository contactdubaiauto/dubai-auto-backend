package model

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserByEmail struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type UserRegister struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Phone    string `json:"phone" binding:"required,min=6,max=15"`
	Username string `json:"username" binding:"required,min=3,max=20"`
}
