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

func (s *AuthService) UserLoginMail(ctx context.Context, user *model.UserLoginMailRequest) model.Response {
	userByEmail, err := s.repo.UserByMail(ctx, &user.Email)

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

func (s *AuthService) UserMailConfirmation(ctx context.Context, user *model.UserMailConfirmationRequest) model.Response {
	otp := pkg.RandomOTP()
	// todo: send otp to the mail
	fmt.Println(otp)
	otp = 123456
	username := pkg.RandomUsername()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(fmt.Sprintf("%d", otp)), bcrypt.DefaultCost)

	if err != nil {
		return model.Response{
			Error:  err,
			Status: http.StatusInternalServerError,
		}
	}

	err = s.repo.UserMailGetOrRegister(ctx, username, user.Email, string(hashedPassword))

	if err != nil {
		return model.Response{
			Error:  err,
			Status: http.StatusInternalServerError,
		}
	}
	return model.Response{
		Data: model.Success{Message: "Successfully sended mail confirmation code."},
	}
}

func (s *AuthService) UserRegister(ctx context.Context, user *model.UserRegisterRequest) model.Response {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return model.Response{
			Error:  err,
			Status: http.StatusInternalServerError,
		}
	}

	user.Password = string(hashedPassword)
	id, err := s.repo.UserRegister(ctx, user)

	if err != nil {
		return model.Response{
			Error:  fmt.Errorf("user already exist: %w", err),
			Status: http.StatusConflict,
		}
	}

	accessToken, refreshToken := pkg.CreateRefreshAccsessToken(*id, 1)
	return model.Response{
		Data: model.LoginResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}
}
