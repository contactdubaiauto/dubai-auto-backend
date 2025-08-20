package service

import (
	"dubai-auto/internal/model"
	"dubai-auto/internal/repository"

	"github.com/valyala/fasthttp"
)

type MotorcycleService struct {
	repository *repository.MotorcycleRepository
}

func NewMotorcycleService(repository *repository.MotorcycleRepository) *MotorcycleService {
	return &MotorcycleService{repository}
}

func (s *MotorcycleService) GetMotorcycleCategories(ctx *fasthttp.RequestCtx) (data []model.GetMotorcycleCategoriesResponse, err error) {
	return s.repository.GetMotorcycleCategories(ctx)
}

func (s *MotorcycleService) GetMotorcycleParameters(ctx *fasthttp.RequestCtx, categoryID string) (data []model.GetMotorcycleParametersResponse, err error) {
	return s.repository.GetMotorcycleParameters(ctx, categoryID)
}

func (s *MotorcycleService) GetMotorcycleBrands(ctx *fasthttp.RequestCtx, categoryID string) (data []model.GetMotorcycleBrandsResponse, err error) {
	return s.repository.GetMotorcycleBrands(ctx, categoryID)
}

func (s *MotorcycleService) GetMotorcycleModelsByBrandID(ctx *fasthttp.RequestCtx, categoryID string, brandID string) (data []model.GetMotorcycleModelsResponse, err error) {
	return s.repository.GetMotorcycleModelsByBrandID(ctx, categoryID, brandID)
}

func (s *MotorcycleService) CreateMotorcycle(ctx *fasthttp.RequestCtx, motorcycle model.CreateMotorcycleRequest, userID int) (data model.SuccessWithId, err error) {
	return s.repository.CreateMotorcycle(ctx, motorcycle, userID)
}

func (s *MotorcycleService) GetMotorcycles(ctx *fasthttp.RequestCtx) (data []model.GetMotorcyclesResponse, err error) {
	return s.repository.GetMotorcycles(ctx)
}

// func (s *MotorcycleService) GetMotorcycleByID(ctx *fasthttp.RequestCtx, id int) (data model.GetMotorcyclesResponse, err error) {
// 	return s.repository.GetMotorcycleByID(ctx, id)
// }

// func (s *MotorcycleService) GetEditMotorcycleByID(ctx *fasthttp.RequestCtx, id int) (data model.GetMotorcyclesResponse, err error) {
// 	return s.repository.GetEditMotorcycleByID(ctx, id)
// }
