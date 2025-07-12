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

func (s *UserService) GetMyCars(ctx *context.Context, userID *int) *model.Response {
	cars, err := s.UserRepository.GetMyCars(ctx, userID)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	return &model.Response{Data: cars}
}

func (s *UserService) OnSale(ctx *context.Context, userID *int) *model.Response {
	cars, err := s.UserRepository.OnSale(ctx, userID)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	return &model.Response{Data: cars}
}

func (s *UserService) Cancel(ctx *context.Context, carID *int) *model.Response {
	err := s.UserRepository.Cancel(ctx, carID)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	return &model.Response{Data: model.Success{Message: "succesfully cancelled"}}
}

func (s *UserService) Delete(ctx *context.Context, carID *int) *model.Response {
	err := s.UserRepository.Delete(ctx, carID)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	return &model.Response{Data: model.Success{Message: "succesfully cancelled"}}
}

func (s *UserService) DontSell(ctx *context.Context, carID, userID *int) *model.Response {
	err := s.UserRepository.DontSell(ctx, carID, userID)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	return &model.Response{Data: model.Success{Message: "succesfully cancelled"}}
}

func (s *UserService) Sell(ctx *context.Context, carID, userID *int) *model.Response {
	err := s.UserRepository.Sell(ctx, carID, userID)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	return &model.Response{Data: model.Success{Message: "succesfully cancelled"}}
}

func (s *UserService) GetBrands(ctx *context.Context, text string) *model.Response {
	brands, err := s.UserRepository.GetBrands(ctx, text)

	if err != nil {
		return &model.Response{
			Error:  err,
			Status: http.StatusBadRequest,
		}
	}
	return &model.Response{
		Data: brands,
	}
}

func (s *UserService) GetFilterBrands(ctx *context.Context, text string) *model.Response {
	brands, err := s.UserRepository.GetFilterBrands(ctx, text)

	if err != nil {
		return &model.Response{
			Error:  err,
			Status: http.StatusBadRequest,
		}
	}
	return &model.Response{
		Data: brands,
	}
}

func (s *UserService) GetCities(ctx *context.Context, text string) *model.Response {
	cities, err := s.UserRepository.GetCities(ctx, text)

	if err != nil {
		return &model.Response{
			Error:  err,
			Status: http.StatusBadRequest,
		}
	}
	return &model.Response{
		Data: cities,
	}
}

func (s *UserService) GetModifications(ctx *context.Context, generationID, bodyTypeID, fuelTypeID, drivetrainID, transmissionID int) *model.Response {
	modifications, err := s.UserRepository.GetModifications(ctx, generationID, bodyTypeID, fuelTypeID, drivetrainID, transmissionID)

	if err != nil {
		return &model.Response{
			Error:  err,
			Status: http.StatusBadRequest,
		}
	}
	return &model.Response{
		Data: modifications,
	}
}

func (s *UserService) GetModelsByBrandID(ctx *context.Context, brandID int64, text string) *model.Response {
	data, err := s.UserRepository.GetModelsByBrandID(ctx, brandID, text)

	if err != nil {
		return &model.Response{Error: err, Status: 400}
	}
	return &model.Response{Data: data}
}

func (s *UserService) GetFilterModelsByBrandID(ctx *context.Context, brandID int64, text string) *model.Response {
	data, err := s.UserRepository.GetFilterModelsByBrandID(ctx, brandID, text)

	if err != nil {
		return &model.Response{Error: err, Status: 400}
	}
	return &model.Response{Data: data}
}

func (s *UserService) GetGenerationsByModelID(ctx *context.Context, modelID int, wheel bool, year, bodyTypeID string) *model.Response {
	data, err := s.UserRepository.GetGenerationsByModelID(ctx, modelID, wheel, year, bodyTypeID)

	if err != nil {
		return &model.Response{Error: err, Status: 400}
	}
	return &model.Response{Data: data}
}

func (s *UserService) GetYearsByModelID(ctx *context.Context, modelID int64, wheel bool) *model.Response {
	data, err := s.UserRepository.GetYearsByModelID(ctx, modelID, wheel)

	if err != nil {
		return &model.Response{Error: err, Status: 400}
	}
	return &model.Response{Data: data}
}

func (s *UserService) GetBodysByModelID(ctx *context.Context, modelID int, wheel bool, year string) *model.Response {
	data, err := s.UserRepository.GetBodysByModelID(ctx, modelID, wheel, year)

	if err != nil {
		return &model.Response{Error: err, Status: 400}
	}
	return &model.Response{Data: data}
}

func (s *UserService) GetBodyTypes(ctx *context.Context) *model.Response {
	data, err := s.UserRepository.GetBodyTypes(ctx)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	return &model.Response{Data: data}
}

func (s *UserService) GetTransmissions(ctx *context.Context) *model.Response {
	data, err := s.UserRepository.GetTransmissions(ctx)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusInternalServerError}
	}
	return &model.Response{Data: data}
}

func (s *UserService) GetEngines(ctx *context.Context) *model.Response {
	data, err := s.UserRepository.GetEngines(ctx)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusInternalServerError}
	}
	return &model.Response{Data: data}
}

func (s *UserService) GetDrivetrains(ctx *context.Context) *model.Response {
	data, err := s.UserRepository.GetDrivetrains(ctx)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusInternalServerError}
	}
	return &model.Response{Data: data}
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

func (s *UserService) GetCars(ctx *context.Context, userID int, brands, models, regions, cities,
	generations, transmissions, engines, drivetrains, body_types, fuel_types, ownership_types []string, year_from, year_to, exchange, credit, right_hand_drive, price_from, price_to string) *model.Response {

	cars, err := s.UserRepository.GetCars(ctx, userID, brands, models, regions, cities,
		generations, transmissions, engines, drivetrains, body_types, fuel_types,
		ownership_types, year_from, year_to, exchange, credit,
		right_hand_drive, price_from, price_to)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	return &model.Response{Data: cars}
}

func (s *UserService) GetCarByID(ctx *context.Context, carID, userID int) *model.Response {
	car, err := s.UserRepository.GetCarByID(ctx, carID, userID)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusNotFound}
	}

	return &model.Response{Data: car}
}

func (s *UserService) BuyCar(ctx *context.Context, carID, userID int) *model.Response {
	err := s.UserRepository.BuyCar(ctx, carID, userID)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusNotFound}
	}

	return &model.Response{Data: model.Success{Message: "successfully buy a car"}}
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
