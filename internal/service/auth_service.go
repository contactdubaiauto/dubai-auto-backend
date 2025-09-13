package service

import (
	"fmt"
	"log"
	"net/http"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/valyala/fasthttp"
	"golang.org/x/crypto/bcrypt"

	"dubai-auto/internal/config"
	"dubai-auto/internal/model"
	"dubai-auto/internal/repository"
	"dubai-auto/internal/utils"
	"dubai-auto/pkg/auth"
)

type AuthService struct {
	repo *repository.AuthRepository
}

func NewAuthService(repo *repository.AuthRepository) *AuthService {
	return &AuthService{repo}
}

func (s *AuthService) UserLoginGoogle(ctx *fasthttp.RequestCtx, tokenID string) model.Response {
	var claims model.IDTokenClaims
	provider, err := oidc.NewProvider(ctx, "https://accounts.google.com")

	if err != nil {
		log.Fatalf("failed to get provider: %v", err)
	}

	verifier := provider.Verifier(&oidc.Config{ClientID: config.ENV.ClientID})
	idToken, err := verifier.Verify(ctx, tokenID)

	if err != nil {
		return model.Response{Error: err, Status: http.StatusUnauthorized}
	}

	if err := idToken.Claims(&claims); err != nil {
		return model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	u, err := s.repo.UserLoginGoogle(ctx, claims)

	if err != nil {
		return model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	accessToken, refreshToken := auth.CreateRefreshAccsessToken(u.ID, 1)

	return model.Response{
		Data: model.LoginFiberResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}
}

func (s *AuthService) DeleteAccount(ctx *fasthttp.RequestCtx, userID int) *model.Response {
	err := s.repo.DeleteAccount(ctx, userID)
	if err != nil {
		return &model.Response{Error: err, Status: http.StatusInternalServerError}
	}
	return &model.Response{Data: model.Success{Message: "Account deleted successfully"}}
}

func (s *AuthService) UserEmailConfirmation(ctx *fasthttp.RequestCtx, user *model.UserEmailConfirmationRequest) model.Response {
	u, err := s.repo.UserByEmail(ctx, &user.Email)

	if err != nil {
		return model.Response{
			Error:  err,
			Status: http.StatusNotFound,
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.OTP), []byte(user.OTP))

	if err != nil {
		return model.Response{
			Error:  err,
			Status: http.StatusBadRequest,
		}
	}

	u.ID, err = s.repo.UserEmailGetOrRegister(ctx, u.Username, user.Email, u.OTP)

	if err != nil {
		return model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	accessToken, refreshToken := auth.CreateRefreshAccsessToken(u.ID, 1)

	return model.Response{
		Data: model.LoginFiberResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}
}

func (s *AuthService) UserPhoneConfirmation(ctx *fasthttp.RequestCtx, user *model.UserPhoneConfirmationRequest) model.Response {
	u, err := s.repo.UserByPhone(ctx, &user.Phone)

	if err != nil {
		return model.Response{
			Error:  err,
			Status: http.StatusNotFound,
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.OTP), []byte(user.OTP))

	if err != nil {
		return model.Response{
			Error:  err,
			Status: http.StatusBadRequest,
		}
	}
	u.ID, err = s.repo.UserPhoneGetOrRegister(ctx, u.Username, user.Phone, u.OTP)

	if err != nil {
		return model.Response{Error: err, Status: http.StatusInternalServerError}
	}
	accessToken, refreshToken := auth.CreateRefreshAccsessToken(u.ID, 1)

	return model.Response{
		Data: model.LoginFiberResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}
}

func (s *AuthService) UserLoginEmail(ctx *fasthttp.RequestCtx, user *model.UserLoginEmail) model.Response {
	otp := utils.RandomOTP()
	// todo: send otp to the mail
	otp = 123456
	username := utils.RandomUsername()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(fmt.Sprintf("%d", otp)), bcrypt.DefaultCost)

	if err != nil {
		return model.Response{
			Error:  err,
			Status: http.StatusInternalServerError,
		}
	}

	err = s.repo.TempUserEmailGetOrRegister(ctx, username, user.Email, string(hashedPassword))

	if err != nil {
		return model.Response{
			Error:  err,
			Status: http.StatusInternalServerError,
		}
	}
	return model.Response{
		Data: model.Success{Message: "Successfully created the user"},
	}
}

func (s *AuthService) UserLoginPhone(ctx *fasthttp.RequestCtx, user *model.UserLoginPhone) model.Response {
	otp := utils.RandomOTP()
	// todo: send otp to the mail
	otp = 123456
	username := utils.RandomUsername()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(fmt.Sprintf("%d", otp)), bcrypt.DefaultCost)

	if err != nil {
		return model.Response{
			Error:  err,
			Status: http.StatusInternalServerError,
		}
	}

	err = s.repo.TempUserPhoneGetOrRegister(ctx, username, user.Phone, string(hashedPassword))

	if err != nil {
		return model.Response{
			Error:  err,
			Status: http.StatusInternalServerError,
		}
	}

	return model.Response{
		Data: model.Success{Message: "Successfully created the user."},
	}
}
