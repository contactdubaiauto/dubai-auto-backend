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

func (s *UserService) GetBrands(ctx *context.Context, text string) ([]*model.GetBrandsResponse, error) {
	return s.UserRepository.GetBrands(ctx, text)
}

func (s *UserService) GetModelsByBrandID(ctx *context.Context, brandID int64, text string) ([]model.Model, error) {
	return s.UserRepository.GetModelsByBrandID(ctx, brandID, text)
}

func (s *UserService) GetGenerationsByModelID(ctx *context.Context, modelID int64) ([]model.Generation, error) {
	return s.UserRepository.GetGenerationsByModelID(ctx, modelID)
}

func (s *UserService) GetBodyTypes(ctx *context.Context) ([]model.BodyType, error) {
	return s.UserRepository.GetBodyTypes(ctx)
}

func (s *UserService) GetTransmissions(ctx *context.Context) ([]model.Transmission, error) {
	return s.UserRepository.GetTransmissions(ctx)
}

func (s *UserService) GetEngines(ctx *context.Context) ([]model.Engine, error) {
	return s.UserRepository.GetEngines(ctx)
}

func (s *UserService) GetDrivetrains(ctx *context.Context) ([]model.Drivetrain, error) {
	return s.UserRepository.GetDrivetrains(ctx)
}

func (s *UserService) GetFuelTypes(ctx *context.Context) ([]model.FuelType, error) {
	return s.UserRepository.GetFuelTypes(ctx)
}

func (s *UserService) GetCars(ctx *context.Context) ([]model.GetCarsResponse, error) {
	return s.UserRepository.GetCars(ctx)
}

func (s *UserService) CreateCar(ctx *context.Context, car *model.CreateCarRequest) *model.Response {

	id, err := s.UserRepository.CreateCar(ctx, car)

	if err != nil {
		return &model.Response{
			Status: 500,
			Error:  err,
		}
	}

	return &model.Response{
		Data: model.SuccessWithId{Id: id, Message: "Car created successfully"},
	}
}

func (s *UserService) CreateCarImages(ctx *context.Context, carID int, images []string) *model.Response {
	err := s.UserRepository.CreateCarImages(ctx, carID, images)

	if err != nil {
		return &model.Response{
			Status: 500,
			Error:  err,
		}
	}

	return &model.Response{
		Data: model.Success{Message: "Car images created successfully"},
	}
}
