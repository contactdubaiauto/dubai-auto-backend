package service

import (
	"context"

	"empty/internal/model"
	"empty/internal/repository"
)

type UserService struct {
	UserRepository *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo}
}

func (s *UserService) CreateUser(ctx context.Context, user *model.User) error {
	return s.UserRepository.Create(ctx, user)
}

func (s *UserService) GetUserByID(ctx context.Context, id int64) (*model.User, error) {
	return s.UserRepository.GetByID(ctx, id)
}

// func (s *UserService) GetUserByID(id int64) (*model.User, error) {
// 	return s.UserInframodel.GetByID(id)
// }

// func (s *UserService) GetAllUsers() ([]*model.User, error) {
// 	return s.UserInframodel.GetAll()
// }

// func (s *UserService) UpdateUser(user *model.User) error {
// 	return s.UserInframodel.Update(user)
// }

// func (s *UserService) DeleteUser(id int64) error {
// 	return s.UserInframodel.Delete(id)
// }
