package service

import (
	"dubai-auto/internal/config"
	"dubai-auto/internal/model"
	"dubai-auto/internal/repository"
	"dubai-auto/pkg/files"
	"errors"
	"mime/multipart"
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

func (s *ThirdPartyService) GetProfile(ctx *fasthttp.RequestCtx, id int) model.Response {
	return s.repo.GetProfile(ctx, id)
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
