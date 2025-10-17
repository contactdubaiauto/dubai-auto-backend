package service

import (
	"dubai-auto/internal/config"
	"dubai-auto/internal/model"
	"dubai-auto/internal/repository"
	"dubai-auto/pkg/files"
	"errors"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/valyala/fasthttp"
)

type ThirdPartyService struct {
	repo *repository.ThirdPartyRepository
}

func NewThirdPartyService(repo *repository.ThirdPartyRepository) *ThirdPartyService {
	return &ThirdPartyService{repo}
}

func (s *ThirdPartyService) Profile(ctx *fasthttp.RequestCtx, id int, profile model.ThirdPartyProfileReq) model.Response {
	return s.repo.Profile(ctx, id, profile)
}

func (s *ThirdPartyService) FirstLogin(ctx *fasthttp.RequestCtx, id int, profile model.ThirdPartyFirstLoginReq) model.Response {
	return s.repo.FirstLogin(ctx, id, profile)
}

func (s *ThirdPartyService) GetProfile(ctx *fasthttp.RequestCtx, id int) model.Response {
	return s.repo.GetProfile(ctx, id)
}

func (s *ThirdPartyService) GetMyCars(ctx *fasthttp.RequestCtx, userID int) model.Response {
	cars, err := s.repo.GetMyCars(ctx, userID)

	if err != nil {
		return model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	return model.Response{Data: cars}
}

func (s *ThirdPartyService) OnSale(ctx *fasthttp.RequestCtx, userID int) model.Response {
	cars, err := s.repo.OnSale(ctx, userID)

	if err != nil {
		return model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	return model.Response{Data: cars}
}

func (s *ThirdPartyService) GetRegistrationData(ctx *fasthttp.RequestCtx) model.Response {
	return s.repo.GetRegistrationData(ctx)
}

func (s *ThirdPartyService) CreateAvatarImages(ctx *fasthttp.RequestCtx, form *multipart.Form, id int) model.Response {

	if form == nil {
		return model.Response{
			Status: 400,
			Error:  errors.New("didn't upload the file"),
		}
	}

	images := form.File["avatar_image"]

	if len(images) > 1 {
		return model.Response{
			Status: 400,
			Error:  errors.New("must load maximum 1 file"),
		}
	}

	paths, status, err := files.SaveFiles(images, config.ENV.STATIC_PATH+"users/"+strconv.Itoa(id)+"/avatar", config.ENV.DEFAULT_IMAGE_WIDTHS)

	if err != nil {
		return model.Response{
			Status: status,
			Error:  err,
		}
	}

	err = s.repo.CreateAvatarImages(ctx, id, paths)

	if err != nil {
		return model.Response{
			Status: 500,
			Error:  err,
		}
	}

	return model.Response{Data: model.Success{Message: "Avatar images updated successfully"}}
}

func (s *ThirdPartyService) CreateBannerImage(ctx *fasthttp.RequestCtx, form *multipart.Form, id int) model.Response {

	if form == nil {
		return model.Response{
			Status: 400,
			Error:  errors.New("didn't upload the file"),
		}
	}

	images := form.File["banner_image"]

	if len(images) > 1 {
		return model.Response{
			Status: 400,
			Error:  errors.New("must load maximum 1 file"),
		}
	}

	paths, status, err := files.SaveFiles(images, config.ENV.STATIC_PATH+"users/"+strconv.Itoa(id)+"/banner", config.ENV.DEFAULT_IMAGE_WIDTHS)

	if err != nil {
		return model.Response{
			Status: status,
			Error:  err,
		}
	}

	err = s.repo.CreateBannerImage(ctx, id, paths)

	if err != nil {
		return model.Response{
			Status: 500,
			Error:  err,
		}
	}

	return model.Response{Data: model.Success{Message: "Banner image updated successfully"}}
}

func (s *ThirdPartyService) CreateDealerCar(ctx *fasthttp.RequestCtx, car *model.ThirdPartyCreateCarRequest, dealerID int) model.Response {
	id, err := s.repo.CreateDealerCar(ctx, car, dealerID)

	if err != nil {
		return model.Response{
			Status: 400,
			Error:  err,
		}
	}

	return model.Response{
		Data: model.SuccessWithId{Id: id, Message: "Car created successfully"},
	}
}

func (s *ThirdPartyService) UpdateDealerCar(ctx *fasthttp.RequestCtx, car *model.UpdateCarRequest, dealerID int) model.Response {
	err := s.repo.UpdateDealerCar(ctx, car, dealerID)

	if err != nil {
		return model.Response{
			Status: 400,
			Error:  err,
		}
	}

	return model.Response{
		Data: model.Success{Message: "Car updated successfully"},
	}
}

func (s *ThirdPartyService) GetEditDealerCarByID(ctx *fasthttp.RequestCtx, carID, dealerID int) *model.Response {
	car, err := s.repo.GetEditDealerCarByID(ctx, carID, dealerID)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusNotFound}
	}

	return &model.Response{Data: car}
}

func (s *ThirdPartyService) CreateDealerCarImages(ctx *fasthttp.RequestCtx, carID int, images []string) *model.Response {
	err := s.repo.CreateDealerCarImages(ctx, carID, images)

	if err != nil {
		return &model.Response{
			Status: 500,
			Error:  err,
		}
	}

	return &model.Response{
		Data: model.Success{Message: "Dealer car images created successfully"},
	}
}

func (s *ThirdPartyService) CreateDealerCarVideos(ctx *fasthttp.RequestCtx, carID int, video string) *model.Response {
	err := s.repo.CreateDealerCarVideos(ctx, carID, video)

	if err != nil {
		return &model.Response{
			Status: 500,
			Error:  err,
		}
	}

	return &model.Response{
		Data: model.Success{Message: "Dealer car videos created successfully"},
	}
}

func (s *ThirdPartyService) DealerDontSell(ctx *fasthttp.RequestCtx, carID, dealerID *int) model.Response {
	err := s.repo.DealerDontSell(ctx, carID, dealerID)

	if err != nil {
		return model.Response{
			Status: 500,
			Error:  err,
		}
	}

	return model.Response{Data: model.Success{Message: "successfully updated status"}}
}

func (s *ThirdPartyService) DealerSell(ctx *fasthttp.RequestCtx, carID, dealerID *int) model.Response {
	err := s.repo.DealerSell(ctx, carID, dealerID)

	if err != nil {
		return model.Response{
			Status: 500,
			Error:  err,
		}
	}

	return model.Response{Data: model.Success{Message: "successfully updated status"}}
}

func (s *ThirdPartyService) DeleteDealerCarImage(ctx *fasthttp.RequestCtx, carID int, imagePath string) *model.Response {
	err := s.repo.DeleteDealerCarImage(ctx, carID, imagePath)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	return &model.Response{Data: model.Success{Message: "Dealer car image deleted successfully"}}
}

func (s *ThirdPartyService) DeleteDealerCarVideo(ctx *fasthttp.RequestCtx, carID int, videoPath string) *model.Response {
	err := s.repo.DeleteDealerCarVideo(ctx, carID, videoPath)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	return &model.Response{Data: model.Success{Message: "Dealer car video deleted successfully"}}
}

func (s *ThirdPartyService) DeleteDealerCar(ctx *fasthttp.RequestCtx, id int) model.Response {
	err := s.repo.DeleteDealerCar(ctx, id)

	if err != nil {
		return model.Response{
			Status: 500,
			Error:  err,
		}
	}

	// delete files

	return model.Response{Data: model.Success{Message: "Car deleted successfully"}}
}

func (s *ThirdPartyService) GetLogistDestinations(ctx *fasthttp.RequestCtx) model.Response {
	destinations, err := s.repo.GetLogistDestinations(ctx)

	if err != nil {
		return model.Response{
			Status: 500,
			Error:  err,
		}
	}

	return model.Response{Data: destinations}
}

func (s *ThirdPartyService) CreateLogistDestination(ctx *fasthttp.RequestCtx, req model.CreateLogistDestinationRequest) model.Response {
	id, err := s.repo.CreateLogistDestination(ctx, req)

	if err != nil {
		return model.Response{
			Status: 400,
			Error:  err,
		}
	}

	return model.Response{
		Data: model.SuccessWithId{Id: id, Message: "Destination created successfully"},
	}
}

func (s *ThirdPartyService) DeleteLogistDestination(ctx *fasthttp.RequestCtx, id int) model.Response {
	err := s.repo.DeleteLogistDestination(ctx, id)

	if err != nil {
		return model.Response{
			Status: 500,
			Error:  err,
		}
	}

	return model.Response{Data: model.Success{Message: "Destination deleted successfully"}}
}
