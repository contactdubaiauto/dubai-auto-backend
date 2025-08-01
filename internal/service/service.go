package service

import (
	"net/http"

	"dubai-auto/internal/model"
	"dubai-auto/internal/repository"
	"dubai-auto/pkg"

	"github.com/valyala/fasthttp"
)

type UserService struct {
	UserRepository *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo}
}

func (s *UserService) GetMyCars(ctx *fasthttp.RequestCtx, userID *int) *model.Response {
	cars, err := s.UserRepository.GetMyCars(ctx, userID)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	return &model.Response{Data: cars}
}

func (s *UserService) OnSale(ctx *fasthttp.RequestCtx, userID *int) *model.Response {
	cars, err := s.UserRepository.OnSale(ctx, userID)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	return &model.Response{Data: cars}
}

func (s *UserService) Cancel(ctx *fasthttp.RequestCtx, carID *int, dir string) *model.Response {
	err := s.UserRepository.Cancel(ctx, carID)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusInternalServerError}
	}
	pkg.RemoveFolder(dir)

	return &model.Response{Data: model.Success{Message: "succesfully cancelled"}}
}

func (s *UserService) DeleteCar(ctx *fasthttp.RequestCtx, carID *int, dir string) *model.Response {
	err := s.UserRepository.DeleteCar(ctx, carID)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusInternalServerError}
	}
	pkg.RemoveFolder(dir)

	return &model.Response{Data: model.Success{Message: "succesfully deleted"}}
}

func (s *UserService) DontSell(ctx *fasthttp.RequestCtx, carID, userID *int) *model.Response {
	err := s.UserRepository.DontSell(ctx, carID, userID)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	return &model.Response{Data: model.Success{Message: "succesfully updated status"}}
}

func (s *UserService) Sell(ctx *fasthttp.RequestCtx, carID, userID *int) *model.Response {
	err := s.UserRepository.Sell(ctx, carID, userID)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	return &model.Response{Data: model.Success{Message: "succesfully updated status"}}
}

func (s *UserService) GetBrands(ctx *fasthttp.RequestCtx, text string) *model.Response {
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

func (s *UserService) GetProfile(ctx *fasthttp.RequestCtx, userID int) *model.Response {
	profile, err := s.UserRepository.GetProfile(ctx, userID)

	if err != nil {
		return &model.Response{
			Error:  err,
			Status: http.StatusBadRequest,
		}
	}
	return &model.Response{
		Data: profile,
	}
}

func (s *UserService) UpdateProfile(ctx *fasthttp.RequestCtx, userID int, profile *model.UpdateProfileRequest) *model.Response {
	err := s.UserRepository.UpdateProfile(ctx, userID, profile)

	if err != nil {
		return &model.Response{
			Error:  err,
			Status: http.StatusBadRequest,
		}
	}
	return &model.Response{
		Data: model.Success{Message: "Profile updated successfully"},
	}
}

func (s *UserService) GetFilterBrands(ctx *fasthttp.RequestCtx, text string) *model.Response {
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

func (s *UserService) GetCities(ctx *fasthttp.RequestCtx, text string) *model.Response {
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

func (s *UserService) GetModelsByBrandID(ctx *fasthttp.RequestCtx, brandID int64, text string) *model.Response {
	data, err := s.UserRepository.GetModelsByBrandID(ctx, brandID, text)

	if err != nil {
		return &model.Response{Error: err, Status: 400}
	}
	return &model.Response{Data: data}
}

func (s *UserService) GetFilterModelsByBrandID(ctx *fasthttp.RequestCtx, brandID int64, text string) *model.Response {
	data, err := s.UserRepository.GetFilterModelsByBrandID(ctx, brandID, text)

	if err != nil {
		return &model.Response{Error: err, Status: 400}
	}
	return &model.Response{Data: data}
}

func (s *UserService) GetFilterModelsByBrands(ctx *fasthttp.RequestCtx, brands []int, text string) *model.Response {
	data, err := s.UserRepository.GetFilterModelsByBrands(ctx, brands, text)

	if err != nil {
		return &model.Response{Error: err, Status: 400}
	}
	return &model.Response{Data: data}
}

func (s *UserService) GetGenerationsByModelID(ctx *fasthttp.RequestCtx, modelID int, wheel bool, year, bodyTypeID string) *model.Response {
	data, err := s.UserRepository.GetGenerationsByModelID(ctx, modelID, wheel, year, bodyTypeID)

	if err != nil {
		return &model.Response{Error: err, Status: 400}
	}
	return &model.Response{Data: data}
}

func (s *UserService) GetGenerationsByModels(ctx *fasthttp.RequestCtx, models []int) *model.Response {
	data, err := s.UserRepository.GetGenerationsByModels(ctx, models)

	if err != nil {
		return &model.Response{Error: err, Status: 400}
	}
	return &model.Response{Data: data}
}

func (s *UserService) GetYearsByModelID(ctx *fasthttp.RequestCtx, modelID int64, wheel bool) *model.Response {
	data, err := s.UserRepository.GetYearsByModelID(ctx, modelID, wheel)

	if err != nil {
		return &model.Response{Error: err, Status: 400}
	}
	return &model.Response{Data: data}
}

func (s *UserService) GetYearsByModels(ctx *fasthttp.RequestCtx, models []int, wheel bool) *model.Response {
	data, err := s.UserRepository.GetYearsByModels(ctx, models, wheel)

	if err != nil {
		return &model.Response{Error: err, Status: 400}
	}
	return &model.Response{Data: data}
}

func (s *UserService) GetBodysByModelID(ctx *fasthttp.RequestCtx, modelID int, wheel bool, year string) *model.Response {
	data, err := s.UserRepository.GetBodysByModelID(ctx, modelID, wheel, year)

	if err != nil {
		return &model.Response{Error: err, Status: 400}
	}
	return &model.Response{Data: data}
}

func (s *UserService) GetBodysByModels(ctx *fasthttp.RequestCtx, wheel bool, models, years []int) *model.Response {
	data, err := s.UserRepository.GetBodysByModels(ctx, wheel, models, years)

	if err != nil {
		return &model.Response{Error: err, Status: 400}
	}
	return &model.Response{Data: data}
}

func (s *UserService) GetBodyTypes(ctx *fasthttp.RequestCtx) *model.Response {
	data, err := s.UserRepository.GetBodyTypes(ctx)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	return &model.Response{Data: data}
}

func (s *UserService) GetTransmissions(ctx *fasthttp.RequestCtx) *model.Response {
	data, err := s.UserRepository.GetTransmissions(ctx)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusInternalServerError}
	}
	return &model.Response{Data: data}
}

func (s *UserService) GetEngines(ctx *fasthttp.RequestCtx) *model.Response {
	data, err := s.UserRepository.GetEngines(ctx)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusInternalServerError}
	}
	return &model.Response{Data: data}
}

func (s *UserService) GetDrivetrains(ctx *fasthttp.RequestCtx) *model.Response {
	data, err := s.UserRepository.GetDrivetrains(ctx)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusInternalServerError}
	}
	return &model.Response{Data: data}
}

func (s *UserService) GetFuelTypes(ctx *fasthttp.RequestCtx) *model.Response {
	data, err := s.UserRepository.GetFuelTypes(ctx)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusInternalServerError}
	}
	return &model.Response{Data: data}
}

func (s *UserService) GetColors(ctx *fasthttp.RequestCtx) *model.Response {
	data, err := s.UserRepository.GetColors(ctx)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusInternalServerError}
	}
	return &model.Response{Data: data}
}

func (s *UserService) GetHome(ctx *fasthttp.RequestCtx, userID int) *model.Response {
	data, err := s.UserRepository.GetHome(ctx, userID)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusInternalServerError}
	}
	return &model.Response{Data: data}
}

func (s *UserService) GetCars(ctx *fasthttp.RequestCtx, userID int, brands, models, regions, cities,
	generations, transmissions, engines, drivetrains, body_types, fuel_types, ownership_types, colors []string,
	year_from, year_to, credit, price_from, price_to, tradeIn, owners, crash string, new, wheel *bool) *model.Response {

	cars, err := s.UserRepository.GetCars(ctx, userID, brands, models, regions, cities,
		generations, transmissions, engines, drivetrains, body_types, fuel_types,
		ownership_types, colors, year_from, year_to, credit,
		price_from, price_to, tradeIn, owners, crash, new, wheel)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	return &model.Response{Data: cars}
}

func (s *UserService) GetCarByID(ctx *fasthttp.RequestCtx, carID, userID int) *model.Response {
	car, err := s.UserRepository.GetCarByID(ctx, carID, userID)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusNotFound}
	}

	return &model.Response{Data: car}
}

func (s *UserService) GetEditCarByID(ctx *fasthttp.RequestCtx, carID, userID int) *model.Response {
	car, err := s.UserRepository.GetEditCarByID(ctx, carID, userID)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusNotFound}
	}

	return &model.Response{Data: car}
}

func (s *UserService) BuyCar(ctx *fasthttp.RequestCtx, carID, userID int) *model.Response {
	err := s.UserRepository.BuyCar(ctx, carID, userID)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusNotFound}
	}

	return &model.Response{Data: model.Success{Message: "successfully buy a car"}}
}

func (s *UserService) CreateCar(ctx *fasthttp.RequestCtx, car *model.CreateCarRequest) *model.Response {

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

func (s *UserService) UpdateCar(ctx *fasthttp.RequestCtx, car *model.UpdateCarRequest, userID int) *model.Response {
	err := s.UserRepository.UpdateCar(ctx, car, userID)

	if err != nil {
		return &model.Response{
			Status: 400,
			Error:  err,
		}
	}

	return &model.Response{
		Data: model.Success{Message: "Car updated successfully"},
	}
}

func (s *UserService) CarLike(ctx *fasthttp.RequestCtx, carID, userID *int) *model.Response {
	err := s.UserRepository.CarLike(ctx, carID, userID)

	if err != nil {
		return &model.Response{
			Status: 409,
			Error:  err,
		}
	}

	return &model.Response{
		Data: model.Success{Message: "Like created successfully"},
	}
}

func (s *UserService) RemoveLike(ctx *fasthttp.RequestCtx, carID, userID *int) *model.Response {
	err := s.UserRepository.RemoveLike(ctx, carID, userID)

	if err != nil {
		return &model.Response{
			Status: 409,
			Error:  err,
		}
	}

	return &model.Response{
		Data: model.Success{Message: "Like removed successfully"},
	}
}

func (s *UserService) Likes(ctx *fasthttp.RequestCtx, userID *int) *model.Response {
	data, err := s.UserRepository.Likes(ctx, userID)

	if err != nil {
		return &model.Response{
			Status: 409,
			Error:  err,
		}
	}

	return &model.Response{Data: data}
}

func (s *UserService) CreateCarImages(ctx *fasthttp.RequestCtx, carID int, images []string) *model.Response {
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

func (s *UserService) CreateCarVideos(ctx *fasthttp.RequestCtx, carID int, video string) *model.Response {
	err := s.UserRepository.CreateCarVideos(ctx, carID, video)

	if err != nil {
		return &model.Response{
			Status: 500,
			Error:  err,
		}
	}

	return &model.Response{
		Data: model.Success{Message: "Car videos created successfully"},
	}
}

func (s *UserService) DeleteCarImage(ctx *fasthttp.RequestCtx, carID int, imagePath string) *model.Response {
	err := s.UserRepository.DeleteCarImage(ctx, carID, imagePath)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusInternalServerError}
	}
	return &model.Response{Data: model.Success{Message: "Car image deleted successfully"}}
}

func (s *UserService) DeleteCarVideo(ctx *fasthttp.RequestCtx, carID int, videoPath string) *model.Response {
	err := s.UserRepository.DeleteCarVideo(ctx, carID, videoPath)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusInternalServerError}
	}
	return &model.Response{Data: model.Success{Message: "Car video deleted successfully"}}
}
