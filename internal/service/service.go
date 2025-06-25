package service

import (
	"context"
	"net/http"

	"dubai-auto/internal/model"
	"dubai-auto/internal/repository"
)

type UserService struct {
	UserRepository *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo}
}

func (s *UserService) GetProfileCars(ctx *context.Context, userID *int) *model.Response {
	cars, err := s.UserRepository.GetProfileCars(ctx, userID)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	return &model.Response{Data: cars}
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

func (s *UserService) GetBodyTypes(ctx *context.Context) *model.Response {
	data, err := s.UserRepository.GetBodyTypes(ctx)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	return &model.Response{Data: data, Status: http.StatusOK}
}

func (s *UserService) GetTransmissions(ctx *context.Context) *model.Response {
	data, err := s.UserRepository.GetTransmissions(ctx)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusInternalServerError}
	}
	return &model.Response{Data: data, Status: http.StatusOK}
}

func (s *UserService) GetEngines(ctx *context.Context) ([]model.Engine, error) {
	return s.UserRepository.GetEngines(ctx)
}

func (s *UserService) GetDrivetrains(ctx *context.Context) ([]model.Drivetrain, error) {
	return s.UserRepository.GetDrivetrains(ctx)
}

func (s *UserService) GetFuelTypes(ctx *context.Context) *model.Response {
	data, err := s.UserRepository.GetFuelTypes(ctx)
	if err != nil {
		return &model.Response{Error: err, Status: http.StatusInternalServerError}
	}
	return &model.Response{Data: data}
}

func (s *UserService) GetColors(ctx *context.Context) *model.Response {
	data, err := s.UserRepository.GetColors(ctx)
	if err != nil {
		return &model.Response{Error: err, Status: http.StatusInternalServerError}
	}
	return &model.Response{Data: data}
}

func (s *UserService) GetCars(ctx *context.Context) *model.Response {
	cars, err := s.UserRepository.GetCars(ctx)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	return &model.Response{Data: cars}
}

func (s *UserService) GetCarByID(ctx *context.Context, carID int) *model.Response {
	car, err := s.UserRepository.GetCarByID(ctx, carID)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusNotFound}
	}

	return &model.Response{Data: car}
}

func (s *UserService) CreateCar(ctx *context.Context, car *model.CreateCarRequest) *model.Response {

	id, err := s.UserRepository.CreateCar(ctx, car)

	if err != nil {
		return &model.Response{
			Status: 400,
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
