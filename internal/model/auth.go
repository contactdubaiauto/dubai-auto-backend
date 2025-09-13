package model

type UserLoginGoogle struct {
	TokenID string `json:"token_id" binding:"required"`
}

type UserEmailConfirmationRequest struct {
	Email string `json:"email" binding:"required,email"`
	OTP   string `json:"otp" binding:"required"`
}

type UserPhoneConfirmationRequest struct {
	Phone string `json:"phone" binding:"required"`
	OTP   string `json:"otp" binding:"required"`
}

type UserLoginEmail struct {
	Email string `json:"email" binding:"required,email"`
}

type UserLoginPhone struct {
	Phone string `json:"phone" binding:"required"`
}

type UserByEmail struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	OTP      string `json:"otp"`
	Type     int    `json:"type"`
}

type UserByPhone struct {
	ID       int    `json:"id"`
	Phone    string `json:"phone"`
	Username string `json:"username"`
	OTP      string `json:"otp"`
}

type LoginFiberResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type UserRegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Phone    string `json:"phone" binding:"required,min=6,max=15"`
	Username string `json:"username" binding:"required,min=3,max=20"`
}
