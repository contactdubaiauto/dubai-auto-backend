package service

import (
	"context"
	"dubai-auto/internal/model"
	"dubai-auto/internal/repository"
	"dubai-auto/pkg"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo *repository.AuthRepository
}

func NewAuthService(repo *repository.AuthRepository) *AuthService {
	return &AuthService{repo}
}

func (s *AuthService) UserLogin(ctx context.Context, user *model.UserLogin) model.Response {
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
