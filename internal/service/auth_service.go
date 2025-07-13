package service

import (
	"context"
	"dubai-auto/internal/model"
	"dubai-auto/internal/repository"
	"dubai-auto/pkg"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo *repository.AuthRepository
}

func NewAuthService(repo *repository.AuthRepository) *AuthService {
	return &AuthService{repo}
}

func (s *AuthService) UserEmailConfirmation(ctx context.Context, user *model.UserEmailConfirmationRequest) model.Response {
	userByEmail, err := s.repo.UserByEmail(ctx, &user.Email)

	if err != nil {
		return model.Response{
			Error:  err,
			Status: http.StatusNotFound,
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(userByEmail.OTP), []byte(user.OTP))

	if err != nil {
		return model.Response{
			Error:  err,
			Status: http.StatusBadRequest,
		}
	}

	accessToken, refreshToken := pkg.CreateRefreshAccsessToken(userByEmail.ID, 1)

	return model.Response{
		Data: model.LoginResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}
}

func (s *AuthService) UserPhoneConfirmation(ctx context.Context, user *model.UserPhoneConfirmationRequest) model.Response {
	userByPhone, err := s.repo.UserByPhone(ctx, &user.Phone)

	if err != nil {
		return model.Response{
			Error:  err,
			Status: http.StatusNotFound,
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(userByPhone.OTP), []byte(user.OTP))

	if err != nil {
		return model.Response{
			Error:  err,
			Status: http.StatusBadRequest,
		}
	}

	accessToken, refreshToken := pkg.CreateRefreshAccsessToken(userByPhone.ID, 1)

	return model.Response{
		Data: model.LoginResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}
}

func (s *AuthService) UserLoginEmail(ctx context.Context, user *model.UserLoginEmail) model.Response {
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

	err = s.repo.UserEmailGetOrRegister(ctx, username, user.Email, string(hashedPassword))

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

func (s *AuthService) UserLoginPhone(ctx context.Context, user *model.UserLoginPhone) model.Response {
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

	err = s.repo.UserPhoneGetOrRegister(ctx, username, user.Phone, string(hashedPassword))

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
