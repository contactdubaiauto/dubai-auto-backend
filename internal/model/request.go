package model

type CreateNameRequest struct {
	Name string `json:"name" validate:"required,min=2,max=255"`
}

type CreateBodyTypeRequest struct {
	Name  string `json:"name" validate:"required,min=2,max=50"`
	Image string `json:"image"`
}

type IDTokenClaims struct {
	Sub           string `json:"sub"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
}
