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
