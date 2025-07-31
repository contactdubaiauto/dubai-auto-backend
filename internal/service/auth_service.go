package service

import (
	"dubai-auto/internal/model"
	"dubai-auto/internal/repository"
	"dubai-auto/pkg"
	"fmt"
	"net/http"

	"github.com/valyala/fasthttp"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo *repository.AuthRepository
}

func NewAuthService(repo *repository.AuthRepository) *AuthService {
	return &AuthService{repo}
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

	accessToken, refreshToken := pkg.CreateRefreshAccsessToken(u.ID, 1)

	return model.Response{
		Data: model.LoFiberResponse{
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
	accessToken, refreshToken := pkg.CreateRefreshAccsessToken(u.ID, 1)

	return model.Response{
		Data: model.LoFiberResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}
}

func (s *AuthService) UserLoginEmail(ctx *fasthttp.RequestCtx, user *model.UserLoginEmail) model.Response {
	otp := pkg.RandomOTP()
	// todo: send otp to the mail
	otp = 123456
	username := pkg.RandomUsername()
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
	otp := pkg.RandomOTP()
	// todo: send otp to the mail
	otp = 123456
	username := pkg.RandomUsername()
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
