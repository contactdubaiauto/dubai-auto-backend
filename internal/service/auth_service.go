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

func (s *AuthService) UserLogin(ctx context.Context, user *model.UserLoginRequest) model.Response {
	userByEmail, err := s.repo.UserLogin(ctx, user)

	if err != nil {
		return model.Response{
			Error:  err,
			Status: http.StatusNotFound,
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(userByEmail.Password), []byte(user.Password))

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
