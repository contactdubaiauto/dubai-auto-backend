package service

import (
	"net/http"

	"github.com/valyala/fasthttp"

	"dubai-auto/internal/model"
	"dubai-auto/internal/repository"
)

type AdminService struct {
	AdminRepository *repository.AdminRepository
}

func NewAdminService(repo *repository.AdminRepository) *AdminService {
	return &AdminService{repo}
}

// Cities service methods
func (s *AdminService) GetCities(ctx *fasthttp.RequestCtx) *model.Response {
	cities, err := s.AdminRepository.GetCities(ctx)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	return &model.Response{Data: cities}
}

func (s *AdminService) CreateCity(ctx *fasthttp.RequestCtx, req *model.CreateCityRequest) *model.Response {
	id, err := s.AdminRepository.CreateCity(ctx, req)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	return &model.Response{Data: model.SuccessWithId{Id: id, Message: "City created successfully"}}
}

func (s *AdminService) UpdateCity(ctx *fasthttp.RequestCtx, id int, req *model.UpdateCityRequest) *model.Response {
	err := s.AdminRepository.UpdateCity(ctx, id, req)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	return &model.Response{Data: model.Success{Message: "City updated successfully"}}
}

func (s *AdminService) DeleteCity(ctx *fasthttp.RequestCtx, id int) *model.Response {
	err := s.AdminRepository.DeleteCity(ctx, id)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	return &model.Response{Data: model.Success{Message: "City deleted successfully"}}
}

// Brands service methods
func (s *AdminService) GetBrands(ctx *fasthttp.RequestCtx) *model.Response {
	brands, err := s.AdminRepository.GetBrands(ctx)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	return &model.Response{Data: brands}
}

func (s *AdminService) CreateBrand(ctx *fasthttp.RequestCtx, req *model.CreateBrandRequest) *model.Response {
	id, err := s.AdminRepository.CreateBrand(ctx, req)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	return &model.Response{Data: model.SuccessWithId{Id: id, Message: "Brand created successfully"}}
}

func (s *AdminService) UpdateBrand(ctx *fasthttp.RequestCtx, id int, req *model.UpdateBrandRequest) *model.Response {
	err := s.AdminRepository.UpdateBrand(ctx, id, req)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	return &model.Response{Data: model.Success{Message: "Brand updated successfully"}}
}

func (s *AdminService) DeleteBrand(ctx *fasthttp.RequestCtx, id int) *model.Response {
	err := s.AdminRepository.DeleteBrand(ctx, id)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	return &model.Response{Data: model.Success{Message: "Brand deleted successfully"}}
}

// Models service methods
func (s *AdminService) GetModels(ctx *fasthttp.RequestCtx) *model.Response {
	models, err := s.AdminRepository.GetModels(ctx)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	return &model.Response{Data: models}
}

func (s *AdminService) CreateModel(ctx *fasthttp.RequestCtx, req *model.CreateModelRequest) *model.Response {
	id, err := s.AdminRepository.CreateModel(ctx, req)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	return &model.Response{Data: model.SuccessWithId{Id: id, Message: "Model created successfully"}}
}

func (s *AdminService) UpdateModel(ctx *fasthttp.RequestCtx, id int, req *model.UpdateModelRequest) *model.Response {
	err := s.AdminRepository.UpdateModel(ctx, id, req)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	return &model.Response{Data: model.Success{Message: "Model updated successfully"}}
}

func (s *AdminService) DeleteModel(ctx *fasthttp.RequestCtx, id int) *model.Response {
	err := s.AdminRepository.DeleteModel(ctx, id)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	return &model.Response{Data: model.Success{Message: "Model deleted successfully"}}
}

// Body Types service methods
func (s *AdminService) GetBodyTypes(ctx *fasthttp.RequestCtx) *model.Response {
	bodyTypes, err := s.AdminRepository.GetBodyTypes(ctx)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	return &model.Response{Data: bodyTypes}
}

func (s *AdminService) CreateBodyType(ctx *fasthttp.RequestCtx, req *model.CreateBodyTypeRequest) *model.Response {
	id, err := s.AdminRepository.CreateBodyType(ctx, req)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	return &model.Response{Data: model.SuccessWithId{Id: id, Message: "Body type created successfully"}}
}

func (s *AdminService) UpdateBodyType(ctx *fasthttp.RequestCtx, id int, req *model.UpdateBodyTypeRequest) *model.Response {
	err := s.AdminRepository.UpdateBodyType(ctx, id, req)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	return &model.Response{Data: model.Success{Message: "Body type updated successfully"}}
}

func (s *AdminService) DeleteBodyType(ctx *fasthttp.RequestCtx, id int) *model.Response {
	err := s.AdminRepository.DeleteBodyType(ctx, id)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	return &model.Response{Data: model.Success{Message: "Body type deleted successfully"}}
}
