package model

type UserLoginGoogle struct {
	TokenID string `json:"token_id" binding:"required"`
}

type UserApplication struct {
	CompanyName       string `json:"company_name" binding:"required"`
	CompanyTypeID     string `json:"company_type_id" binding:"required"`
	ActivityFieldID   string `json:"activity_field_id" binding:"required"`
	LicenseIssueDate  string `json:"licensei_issue_date" binding:"required"`
	LicenseExpiryDate string `json:"licensei_expiry_date" binding:"required"`
	FullName          string `json:"full_name" binding:"required"`
	Email             string `json:"email" binding:"required"`
	Phone             string `json:"phone" binding:"required"`
	Address           string `json:"address" binding:"required"`
	VATNumber         string `json:"vat_number" binding:"required"`
	RoleID            int    `json:"role_id" binding:"required"`
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
	RoleID   int    `json:"role_id"`
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
