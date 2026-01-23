package service

import (
	"errors"

	"github.com/valyala/fasthttp"

	"dubai-auto/internal/model"
	"dubai-auto/internal/repository"
	"dubai-auto/pkg/files"
)

type MotorcycleService struct {
	repository *repository.MotorcycleRepository
}

func NewMotorcycleService(repository *repository.MotorcycleRepository) *MotorcycleService {
	return &MotorcycleService{repository}
}

func (s *MotorcycleService) GetMotorcycleCategories(ctx *fasthttp.RequestCtx, nameColumn string) model.Response {
	data, err := s.repository.GetMotorcycleCategories(ctx, nameColumn)

	if err != nil {
		return model.Response{
			Status: 500,
			Error:  err,
		}
	}
	return model.Response{
		Status: 200,
		Data:   data,
	}
}

// func (s *MotorcycleService) GetMotorcycleParameters(ctx *fasthttp.RequestCtx, categoryID string, nameColumn string) model.Response {
// 	data, err := s.repository.GetMotorcycleParameters(ctx, categoryID, nameColumn)
// 	if err != nil {
// 		return model.Response{
// 			Status: 500,
// 			Error:  err,
// 		}
// 	}
// 	return model.Response{
// 		Status: 200,
// 		Data:   data,
// 	}
// }

func (s *MotorcycleService) GetMotorcycleBrands(ctx *fasthttp.RequestCtx, categoryID string, nameColumn string) model.Response {
	data, err := s.repository.GetMotorcycleBrands(ctx, categoryID, nameColumn)
	if err != nil {
		return model.Response{
			Status: 500,
			Error:  err,
		}
	}
	return model.Response{
		Status: 200,
		Data:   data,
	}
}

func (s *MotorcycleService) GetNumberOfCycles(ctx *fasthttp.RequestCtx, nameColumn string) model.Response {
	data, err := s.repository.GetNumberOfCycles(ctx, nameColumn)
	if err != nil {
		return model.Response{
			Status: 500,
			Error:  err,
		}
	}
	return model.Response{
		Status: 200,
		Data:   data,
	}
}

func (s *MotorcycleService) GetMotorcycleModelsByBrandID(ctx *fasthttp.RequestCtx, categoryID string, brandID string, nameColumn string) model.Response {
	data, err := s.repository.GetMotorcycleModelsByBrandID(ctx, categoryID, brandID, nameColumn)
	if err != nil {
		return model.Response{
			Status: 500,
			Error:  err,
		}
	}
	return model.Response{
		Status: 200,
		Data:   data,
	}
}

func (s *MotorcycleService) CreateMotorcycle(ctx *fasthttp.RequestCtx, motorcycle model.CreateMotorcycleRequest, userID int) model.Response {
	data, err := s.repository.CreateMotorcycle(ctx, motorcycle, userID)
	if err != nil {
		return model.Response{
			Status: 500,
			Error:  err,
		}
	}
	return model.Response{
		Status: 200,
		Data:   data,
	}
}

func (s *MotorcycleService) GetMotorcycles(ctx *fasthttp.RequestCtx, userID int, brands, models, regions, cities,
	generations, transmissions, engines, drivetrains, body_types, fuel_types, ownership_types, colors, dealers []string,
	year_from, year_to, credit, price_from, price_to, tradeIn, owners, crash, odometer, newQ, wheelQ string,
	limit, lastID int, nameColumn string) model.Response {
	var wheel *bool
	var new *bool

	if newQ != "" {
		if newQ == "false" {
			tmp := false
			new = &tmp
		} else {
			tmp := true
			new = &tmp
		}

	}

	if wheelQ != "" {
		if wheelQ == "false" {
			tmp := false
			wheel = &tmp
		} else {
			tmp := true
			wheel = &tmp
		}

	}

	data, err := s.repository.GetMotorcycles(ctx, userID, brands, models, regions, cities,
		generations, transmissions, engines, drivetrains, body_types, fuel_types,
		ownership_types, colors, dealers, year_from, year_to, credit,
		price_from, price_to, tradeIn, owners, crash, odometer, new, wheel, limit, lastID, nameColumn)
	if err != nil {
		return model.Response{
			Status: 500,
			Error:  err,
		}
	}

	return model.Response{
		Status: 200,
		Data:   data,
	}
}

func (s *MotorcycleService) CreateMotorcycleImages(ctx *fasthttp.RequestCtx, motorcycleID int, images []string) model.Response {

	if len(images) == 0 {
		return model.Response{
			Status: 400,
			Error:  errors.New("images are required"),
		}
	}
	err := s.repository.CreateMotorcycleImages(ctx, motorcycleID, images)
	if err != nil {
		return model.Response{
			Status: 500,
			Error:  err,
		}
	}

	return model.Response{
		Data: model.Success{Message: "Motorcycle images created successfully"},
	}
}

func (s *MotorcycleService) CreateMotorcycleVideos(ctx *fasthttp.RequestCtx, motorcycleID int, video string) model.Response {
	err := s.repository.CreateMotorcycleVideos(ctx, motorcycleID, video)

	if err != nil {
		return model.Response{
			Status: 500,
			Error:  err,
		}
	}

	return model.Response{
		Data: model.Success{Message: "Motorcycle videos created successfully"},
	}
}

func (s *MotorcycleService) DeleteMotorcycleImage(ctx *fasthttp.RequestCtx, motorcycleID int, imageID int) model.Response {

	err := s.repository.DeleteMotorcycleImage(ctx, motorcycleID, imageID)

	if err != nil {
		return model.Response{
			Status: 500,
			Error:  err,
		}
	}
	// todo: delete image if exists
	return model.Response{
		Data: model.Success{Message: "Motorcycle image deleted successfully"},
	}
}

func (s *MotorcycleService) DeleteMotorcycleVideo(ctx *fasthttp.RequestCtx, motorcycleID int, videoID int) model.Response {

	err := s.repository.DeleteMotorcycleVideo(ctx, motorcycleID, videoID)
	if err != nil {
		return model.Response{
			Status: 500,
			Error:  err,
		}
	}

	// todo: delete video if exists
	return model.Response{
		Data: model.Success{Message: "Motorcycle video deleted successfully"},
	}
}

func (s *MotorcycleService) GetMotorcycleByID(ctx *fasthttp.RequestCtx, motorcycleID, userID int, nameColumn string) model.Response {
	motorcycle, err := s.repository.GetMotorcycleByID(ctx, motorcycleID, userID, nameColumn)
	if err != nil {
		return model.Response{
			Status: 404,
			Error:  err,
		}
	}
	return model.Response{
		Status: 200,
		Data:   motorcycle,
	}
}

func (s *MotorcycleService) GetEditMotorcycleByID(ctx *fasthttp.RequestCtx, motorcycleID, userID int, nameColumn string) model.Response {
	motorcycle, err := s.repository.GetEditMotorcycleByID(ctx, motorcycleID, userID, nameColumn)
	if err != nil {
		return model.Response{
			Status: 404,
			Error:  err,
		}
	}
	return model.Response{
		Status: 200,
		Data:   motorcycle,
	}
}

func (s *MotorcycleService) BuyMotorcycle(ctx *fasthttp.RequestCtx, motorcycleID, userID int) model.Response {
	err := s.repository.BuyMotorcycle(ctx, motorcycleID, userID)
	if err != nil {
		return model.Response{
			Status: 500,
			Error:  err,
		}
	}
	return model.Response{
		Status: 200,
		Data:   model.Success{Message: "Successfully bought the motorcycle"},
	}
}

func (s *MotorcycleService) DontSellMotorcycle(ctx *fasthttp.RequestCtx, motorcycleID, userID int) model.Response {
	err := s.repository.DontSellMotorcycle(ctx, motorcycleID, userID)
	if err != nil {
		return model.Response{
			Status: 500,
			Error:  err,
		}
	}
	return model.Response{
		Status: 200,
		Data:   model.Success{Message: "Successfully updated motorcycle status to not for sale"},
	}
}

func (s *MotorcycleService) SellMotorcycle(ctx *fasthttp.RequestCtx, motorcycleID, userID int) model.Response {
	err := s.repository.SellMotorcycle(ctx, motorcycleID, userID)
	if err != nil {
		return model.Response{
			Status: 500,
			Error:  err,
		}
	}
	return model.Response{
		Status: 200,
		Data:   model.Success{Message: "Successfully updated motorcycle status to for sale"},
	}
}

func (s *MotorcycleService) CancelMotorcycle(ctx *fasthttp.RequestCtx, motorcycleID *int, dir string) model.Response {
	err := s.repository.CancelMotorcycle(ctx, motorcycleID)

	if err != nil {
		return model.Response{
			Status: 500,
			Error:  err,
		}
	}

	if dir != "" {
		files.RemoveFolder(dir)
	}

	return model.Response{
		Status: 200,
		Data:   model.Success{Message: "Successfully cancelled motorcycle"},
	}
}

func (s *MotorcycleService) UpdateMotorcycle(ctx *fasthttp.RequestCtx, motorcycle *model.UpdateMotorcycleRequest, userID int) model.Response {
	err := s.repository.UpdateMotorcycle(ctx, motorcycle, userID)

	if err != nil {
		return model.Response{
			Status: 400,
			Error:  err,
		}
	}

	return model.Response{
		Status: 200,
		Data:   model.Success{Message: "Motorcycle updated successfully"},
	}
}

func (s *MotorcycleService) DeleteMotorcycle(ctx *fasthttp.RequestCtx, motorcycleID int, dir string) model.Response {
	err := s.repository.DeleteMotorcycle(ctx, motorcycleID)

	if err != nil {
		return model.Response{
			Status: 500,
			Error:  err,
		}
	}

	if dir != "" {
		files.RemoveFolder(dir)
	}

	return model.Response{
		Status: 200,
		Data:   model.Success{Message: "Successfully deleted motorcycle"},
	}
}
