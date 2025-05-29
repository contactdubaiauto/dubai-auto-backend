package service

import (
	"context"

	"dubai-auto/internal/model"
	"dubai-auto/internal/repository"
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

func (s *UserService) GetAllUsers(ctx context.Context) ([]*model.User, error) {
	return s.UserRepository.GetAll(ctx)
}

func (s *UserService) GetBrands(ctx context.Context) ([]*model.GetBrandsResponse, error) {
	return s.UserRepository.GetBrands(ctx)
}

func (s *UserService) GetModelsByBrandID(ctx context.Context, brandID int64) ([]model.Model, error) {
	return s.UserRepository.GetModelsByBrandID(ctx, brandID)
}

func (s *UserService) GetBodyTypes(ctx context.Context) ([]model.BodyType, error) {
	return s.UserRepository.GetBodyTypes(ctx)
}

func (s *UserService) GetTransmissions(ctx context.Context) ([]model.Transmission, error) {
	return s.UserRepository.GetTransmissions(ctx)
}

func (s *UserService) GetEngines(ctx context.Context) ([]model.Engine, error) {
	return s.UserRepository.GetEngines(ctx)
}

func (s *UserService) GetDrives(ctx context.Context) ([]model.Drive, error) {
	return s.UserRepository.GetDrives(ctx)
}

func (s *UserService) GetFuelTypes(ctx context.Context) ([]model.FuelType, error) {
	return s.UserRepository.GetFuelTypes(ctx)
}
