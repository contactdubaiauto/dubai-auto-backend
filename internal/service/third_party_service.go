package service

import (
	"dubai-auto/internal/model"
	"dubai-auto/internal/repository"

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

func (s *ThirdPartyService) GetProfile(ctx *fasthttp.RequestCtx, id int) model.Response {
	return s.repo.GetProfile(ctx, id)
}

func (s *ThirdPartyService) GetRegistrationData(ctx *fasthttp.RequestCtx) model.Response {
	return s.repo.GetRegistrationData(ctx)
}
